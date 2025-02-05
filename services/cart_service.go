package services

import (
	"github.com/Abdurahmanit/marketplace/backend/models"

	"github.com/Abdurahmanit/marketplace/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddGameToCart добавляет игру в корзину пользователя
func AddGameToCart(userID, gameID string) error {
	// Преобразуем ID пользователя и игры в ObjectID
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	gameObjID, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return err
	}

	// Обновляем корзину пользователя
	_, err = config.GetDBCollection("carts").UpdateOne(
		config.GetContext(),
		bson.M{"user_id": userObjID},
		bson.M{"$addToSet": bson.M{"games": gameObjID}},
		options.Update().SetUpsert(true),
	)
	return err
}

// RemoveGameFromCart удаляет игру из корзины пользователя
func RemoveGameFromCart(userID, gameID string) error {
	// Преобразуем ID пользователя и игры в ObjectID
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	gameObjID, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return err
	}

	// Удаляем игру из корзины пользователя
	_, err = config.GetDBCollection("carts").UpdateOne(
		config.GetContext(),
		bson.M{"user_id": userObjID},
		bson.M{"$pull": bson.M{"games": gameObjID}},
	)
	return err
}

// GetCart возвращает корзину пользователя
func GetCart(userID string) (*models.Cart, error) {
	var cart models.Cart

	// Преобразуем ID пользователя в ObjectID
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	// Ищем корзину пользователя
	err = config.GetDBCollection("carts").FindOne(config.GetContext(), bson.M{"user_id": userObjID}).Decode(&cart)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}
