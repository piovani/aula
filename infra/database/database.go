package database

import "go.mongodb.org/mongo-driver/mongo"

type Database struct {
	Conn *mongo.Client
}

func NewDatabase(conn *mongo.Client) *Database {
	return &Database{
		Conn: conn,
	}
}
