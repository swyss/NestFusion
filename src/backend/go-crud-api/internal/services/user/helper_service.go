package user_service

import (
	usermodel "go-crud-api/internal/models/user"
	repository "go-crud-api/internal/repos/user"
)

// AuthServiceInterface defines the methods for the authentication service.
type AuthServiceInterface interface {
	Authenticate(input usermodel.AuthInput) error
}

// authService implements the AuthServiceInterface
type authService struct {
	userRepo repository.AuthRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepo repository.AuthRepository) AuthServiceInterface {
	return &authService{userRepo}
}

func (s *authService) Authenticate(input usermodel.AuthInput) error {
	// Implement authentication logic here
	return nil
}

// RoleServiceInterface defines the methods for the role assignment service.
type RoleServiceInterface interface {
	AssignRole(input usermodel.RoleInput) error
}

// roleService implements the RoleServiceInterface
type roleService struct {
	userRepo repository.RoleRepository
}

// NewRoleService creates a new instance of RoleService
func NewRoleService(userRepo repository.RoleRepository) RoleServiceInterface {
	return &roleService{userRepo}
}

func (s *roleService) AssignRole(input usermodel.RoleInput) error {
	// Implement role assignment logic here
	return nil
}

// InfoServiceInterface defines the methods for the user info service.
type InfoServiceInterface interface {
	SetUserInfo(input usermodel.InfoInput) error
}

// infoService implements the InfoServiceInterface
type infoService struct {
	userRepo repository.InfoRepository
}

// NewInfoService creates a new instance of InfoService
func NewInfoService(userRepo repository.InfoRepository) InfoServiceInterface {
	return &infoService{userRepo}
}

func (s *infoService) SetUserInfo(input usermodel.InfoInput) error {
	// Implement user info setting logic here
	return nil
}
