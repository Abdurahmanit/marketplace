package services

import (
	"marketplace/backend/models"
	"marketplace/backend/utils"

	"github.com/Abdurahmanit/marketplace/config"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Insert the user into the database
	_, err = config.GetDBCollection("users").InsertOne(config.GetContext(), user)
	return err
}

func LoginUser(email, password string) (string, error) {
	var user models.User
	err := config.GetDBCollection("users").FindOne(config.GetContext(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", err
	}

	// Verify the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		return "", err
	}

	return token, nil
}
