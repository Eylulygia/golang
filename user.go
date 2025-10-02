package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Email    string             `json:"email"`
    Password string             `json:"password"`
    Books    []string           `json:"books"`
}
