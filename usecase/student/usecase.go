package student

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/piovani/aula/entites"
	"github.com/piovani/aula/infra/cache"
	"github.com/piovani/aula/infra/database"
)

const (
	StudentPrefixKeyCache = "student_"
)

type StudentUsecase struct {
	Database *database.Database
	Cache    *cache.Cache
}

func NewStudentUsecase(db *database.Database, c *cache.Cache) *StudentUsecase {
	return &StudentUsecase{
		Database: db,
		Cache:    c,
	}
}

func (su *StudentUsecase) findStudent(id uuid.UUID) (*entites.Student, error) {
	student, err := su.searchCache(id)
	if err != nil {
		student, err = su.searchDB(id)
		if err != nil {
			return &entites.Student{}, err
		}
		su.saveCache(student)
	}

	return student, nil
}

func (su *StudentUsecase) saveStudent(student *entites.Student) (*entites.Student, error) {
	if err := su.deleteCache(student.ID); err != nil {
		return &entites.Student{}, err
	}

	err := su.updateDB(student)
	if err != nil {
		return &entites.Student{}, err
	}

	su.saveCache(student)

	return student, nil
}

func (su *StudentUsecase) deleteStudent(id uuid.UUID) error {
	if err := su.deleteCache(id); err != nil {
		return err
	}

	return su.Database.StudentRepository.Delete(id)
}

// DATABASE
func (su *StudentUsecase) searchDB(id uuid.UUID) (*entites.Student, error) {
	student, err := su.Database.StudentRepository.FindByID(id)
	return &student, err
}

func (su *StudentUsecase) updateDB(student *entites.Student) error {
	return su.Database.StudentRepository.Update(student)
}

// CASHE
func (su *StudentUsecase) searchCache(id uuid.UUID) (*entites.Student, error) {
	var student entites.Student
	key := StudentPrefixKeyCache + id.String()

	studentJson, err := su.Cache.Get(key)
	if err != nil {
		return &entites.Student{}, err
	}

	err = json.Unmarshal([]byte(studentJson), &student)

	return &student, err
}

func (su *StudentUsecase) saveCache(student *entites.Student) error {
	key := StudentPrefixKeyCache + student.ID.String()
	studentJson, err := json.Marshal(student)
	if err != nil {
		return err
	}

	return su.Cache.Set(key, string(studentJson), time.Second*30)
}

func (su *StudentUsecase) deleteCache(id uuid.UUID) error {
	return su.Cache.Delete(StudentPrefixKeyCache + id.String())
}
