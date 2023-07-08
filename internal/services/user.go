package services

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"golang-easy-project/internal/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func GetUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (us *UserService) CreateUser(c *gin.Context) {
	// Parse request body to get user details
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Generate a salt for password hashing
	salt, err := generateSalt()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate salt"})
		return
	}

	// Hash the user's password with the generated salt
	hashedPassword, err := hashPassword(user.Password, salt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Set the hashed password and other user details
	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.Role = models.RoleUser

	// Save the user to the database
	_, err = us.db.Exec("INSERT INTO users (email, password, created_at, role) VALUES (?, ?, ?, ?)", user.Email, user.Password, user.CreatedAt, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusCreated, user)
}

func (us *UserService) GetUserList(c *gin.Context) {
	// Retrieve the user list from the database
	rows, err := us.db.Query("SELECT id, email, created_at, role FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user list"})
		return
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println("Failed to close rows:", err)
		}
	}()

	var userList []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Email, &user.CreatedAt, &user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user list"})
			return
		}
		userList = append(userList, user)
	}

	// Respond with the user list
	c.JSON(http.StatusOK, gin.H{"users": userList})
}

// Helper function to generate a salt for password hashing
func generateSalt() (string, error) {
	saltBytes := make([]byte, 16)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(saltBytes), nil
}

// Helper function to hash the password with the provided salt
func hashPassword(password, salt string) (string, error) {
	passwordBytes := []byte(password)
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", err
	}

	// Generate the hashed password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(append(passwordBytes, saltBytes...), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
