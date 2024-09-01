package user_model

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"time"
)

const DefaultSaltLength = 16

func init() {
	Pepper := os.Getenv("PEPPER")
	if Pepper == "" {
		log.Fatal("PEPPER is not set in the environment variables")
	}
}

// UserPassword represents a user's hashed password with salt and pepper for security.
type UserPassword struct {
	ID        uint           `json:"id" gorm:"primary_key;autoIncrement"`
	Hash      string         `json:"-" gorm:"size:255"` // The hashed password should not be exposed in JSON.
	Salt      string         `json:"-" gorm:"size:255"` // Salt used for hashing the password, hidden from JSON.
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// BeforeSave hashes the password with salt and pepper before saving to the database.
func (p *UserPassword) BeforeSave(*gorm.DB) error {
	Pepper := os.Getenv("PEPPER")
	if len(p.Hash) == 0 {
		return nil
	}

	salt, err := generateSalt(DefaultSaltLength)
	if err != nil {
		return err
	}
	p.Salt = salt

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Hash+salt+Pepper), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.Hash = string(hashedPassword)

	return nil
}

// generateSalt creates a new salt for password hashing.
func generateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}
