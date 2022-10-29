package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/piovani/aula/entites"
	mongo_infra "github.com/piovani/aula/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
	mongo_drive "go.mongodb.org/mongo-driver/mongo"
)

const (
	StudentCollection = "students"
)

type StudentRepository struct {
	Ctx        context.Context
	Collection *mongo_drive.Collection
}

func NewStudentRepository(ctx context.Context, client *mongo_drive.Client) *StudentRepository {
	return &StudentRepository{
		Ctx:        ctx,
		Collection: mongo_infra.GetCollection(ctx, client, StudentCollection),
	}
}

func (sr StudentRepository) Create(student *entites.Student) error {
	_, err := sr.Collection.InsertOne(sr.Ctx, student)

	return err
}

func (sr StudentRepository) List() (students []entites.Student, err error) {
	cur, err := sr.Collection.Find(sr.Ctx, bson.M{})
	if err != nil {
		return students, err
	}

	for cur.Next(sr.Ctx) {
		var student entites.Student

		if err = cur.Decode(&student); err != nil {
			return students, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (sr StudentRepository) FindByID(id uuid.UUID) (student entites.Student, err error) {
	err = sr.Collection.FindOne(sr.Ctx, bson.M{"_id": id}).Decode(&student)

	return student, err
}

func (sr StudentRepository) Update(student *entites.Student) error {
	_, err := sr.Collection.UpdateOne(sr.Ctx, bson.M{"_id": student.ID}, bson.M{"$set": student})

	return err
}

func (sr StudentRepository) Delete(id uuid.UUID) error {
	_, err := sr.Collection.DeleteOne(sr.Ctx, bson.M{"_id": id})

	return err
}
