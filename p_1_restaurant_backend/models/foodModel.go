package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Food struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      *string            `json:"name" bson:"name" validate:"required,min=2,max=100"`
	Price     *float64           `json:"price" bson:"price" validate:"required"`
	FoodImage *string            `json:"food_image" bson:"food_image" validate:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	FoodID    string             `json:"food_id" bson:"food_id"`
	MenuID    *string            `json:"menu_id" bson:"menu_id" validate:"required"`
}
