package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name     string             `bson:"name" json:"name" validate:"required,min=3,max=50"`
    Email    string             `bson:"email" json:"email" validate:"required,email"`
    Status   string             `bson:"status" json:"status" validate:"required,oneof=active inactive blocked"`
    Role     string             `bson:"role" json:"role" validate:"required,oneof=admin user"`
    Password string             `bson:"password,omitempty" json:"password" validate:"required,min=6"`
}


