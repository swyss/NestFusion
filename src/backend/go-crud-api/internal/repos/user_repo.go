package repos

import (
	"database/sql"
	"go-crud-api/internal/models"
)

// UserRepository handles the CRUD operations for the User model in the database.
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new instance of UserRepository with the provided database connection.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetAllUsers retrieves all users from the database.
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	// Execute the query to get all users.
	rows, err := r.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	// Ensure rows are closed after processing to avoid resource leaks.
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			// Handle potential error when closing rows.
		}
	}(rows)

	var users []models.User
	// Iterate over each row in the result set.
	for rows.Next() {
		var u models.User
		// Scan the row into a User struct.
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		// Add the user to the slice.
		users = append(users, u)
	}
	// Return the slice of users and check for any error during iteration.
	return users, rows.Err()
}

// GetUserByID retrieves a single user by their ID.
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var u models.User
	// Execute the query to get a user by ID.
	err := r.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	// Return the user if found.
	return &u, nil
}

// CreateUser inserts a new user into the database.
func (r *UserRepository) CreateUser(u *models.User) error {
	// Execute the insert query and retrieve the generated ID.
	return r.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
}

// UpdateUser updates an existing user in the database.
func (r *UserRepository) UpdateUser(u *models.User) error {
	// Execute the update query for the user.
	_, err := r.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, u.ID)
	return err
}

// DeleteUser deletes a user from the database by their ID.
func (r *UserRepository) DeleteUser(id int) error {
	// Execute the delete query for the user.
	_, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
