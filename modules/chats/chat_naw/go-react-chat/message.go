package main

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

type Message struct {
	Id        uuid.Uuid
	Type      string
	Name      string
	Timestamp time.Time
	Text      string
}
