package database

import (
	"github.com/piovani/aula/entites"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Conn *mongo.Client

	StudentRepository entites.StudentRepository
}

func NewDatabase(conn *mongo.Client, sr entites.StudentRepository) *Database {
	return &Database{
		Conn: conn,

		StudentRepository: sr,
	}
}
