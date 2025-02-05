package services

import (
	"github.com/Abdurahmanit/marketplace/backend/models"

	"github.com/Abdurahmanit/marketplace/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllGames возвращает список всех игр
func GetAllGames() ([]models.Game, error) {
	var games []models.Game

	// Получаем все игры из коллекции `games`
	cursor, err := config.GetDBCollection("games").Find(config.GetContext(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(config.GetContext())

	// Декодируем результаты в структуру Game
	for cursor.Next(config.GetContext()) {
		var game models.Game
		if err := cursor.Decode(&game); err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}

// GetGameByID возвращает игру по её ID
func GetGameByID(gameID string) (*models.Game, error) {
	var game models.Game

	// Преобразуем строку в ObjectID
	objID, err := primitive.ObjectIDFromHex(gameID)
	if err != nil {
		return nil, err
	}

	// Ищем игру по ID
	err = config.GetDBCollection("games").FindOne(config.GetContext(), bson.M{"_id": objID}).Decode(&game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}
