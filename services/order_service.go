package services

import (
	"github.com/Abdurahmanit/marketplace/backend/models"

	"github.com/Abdurahmanit/marketplace/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlaceOrder создает заказ на основе корзины пользователя
func PlaceOrder(userID string) error {
	// Получаем корзину пользователя
	cart, err := GetCart(userID)
	if err != nil {
		return err
	}

	// Создаем заказ
	order := models.Order{
		UserID: cart.UserID,
		Games:  cart.Games,
		Total:  calculateTotal(cart.Games),
	}

	// Вставляем заказ в коллекцию `orders`
	_, err = config.GetDBCollection("orders").InsertOne(config.GetContext(), order)
	if err != nil {
		return err
	}

	// Очищаем корзину пользователя
	_, err = config.GetDBCollection("carts").DeleteOne(config.GetContext(), bson.M{"user_id": cart.UserID})
	return err
}

// GetUserOrders возвращает список заказов пользователя
func GetUserOrders(userID string) ([]models.Order, error) {
	var orders []models.Order

	// Преобразуем ID пользователя в ObjectID
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	// Получаем все заказы пользователя
	cursor, err := config.GetDBCollection("orders").Find(config.GetContext(), bson.M{"user_id": userObjID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(config.GetContext())

	// Декодируем результаты в структуру Order
	for cursor.Next(config.GetContext()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// calculateTotal вычисляет общую стоимость заказа
func calculateTotal(games []primitive.ObjectID) float64 {
	// Здесь должна быть логика расчета стоимости на основе цен игр
	// Для примера возвращаем фиксированное значение
	return 99.99
}
