package service

import (
	"Go-UserManagement/helper"
	"Go-UserManagement/model/domain"
	"Go-UserManagement/model/web"
	"Go-UserManagement/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
}

func (Service *UserServiceImpl) Register(ctx context.Context, r web.UserRegisterRequest) (web.UserResponse, error) {
	// Always begin the Transaction
	// Create helper function to avoid DRY code
	// Validate Struct
	err := Service.Validate.Struct(r)
	if err != nil {
		return web.UserResponse{}, fmt.Errorf("validation failed: %w", err)
	}
	// Validate input
	if r.Email == "" || r.Password == "" || r.Name == "" {
		return web.UserResponse{}, errors.New("all fields are required")
	}

	tx, err := Service.DB.Begin()
	if err != nil {
		return  web.UserResponse{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer helper.CommitOrRollback(tx)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return  web.UserResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := domain.User{
		Email: r.Email,
		Password: string(hashedPassword),
		Name: r.Name,
	}

	// Save to database
	user, err = Service.UserRepository.Register(ctx, tx, user)
	if err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to register user: %w", err)
	}

	token, err := helper.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	response := helper.ToUserResponse(user)
	response.Token = token

	return helper.ToUserResponse(user), nil
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error) {
	fmt.Println("=== LOGIN DEBUG START ===")
	fmt.Printf("Input Email: '%s'\n", request.Email)
	fmt.Printf("Input Password: '%s'\n", request.Password)
	// Validate input
	err := service.Validate.Struct(request)
	if err != nil {
		fmt.Println("❌ Validation failed:", err)
		return web.UserResponse{}, fmt.Errorf("validation failed: %w", err)
	}

	// Find user by email
	user, err := service.UserRepository.FindByEmail(ctx, service.DB, request.Email)
	if err != nil {
		fmt.Println("❌ FindByEmail error:", err)
		return web.UserResponse{}, errors.New("invalid email or password")
	}

	fmt.Printf("✅ User found in DB\n")
	fmt.Printf("DB Email: '%s'\n", user.Email)
	fmt.Printf("DB Password Hash: '%s'\n", user.Password)
	fmt.Printf("Hash length: %d\n", len(user.Password))

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		fmt.Println("❌ bcrypt.CompareHashAndPassword failed:", err)
		return web.UserResponse{}, errors.New("invalid email or password")
	}

	token, err := helper.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	response := helper.ToUserResponse(user)
	response.Token = token

	fmt.Println("✅ Password matches! Login successful")
	fmt.Println("=== LOGIN DEBUG END ===")
	return helper.ToUserResponse(user), nil
}