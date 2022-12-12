package student

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/piovani/aula/infra/database"
	mock "github.com/piovani/aula/infra/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentUsecaseSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	StudentRepository *mock.MockStudentRepository
	StudentUsecase    *StudentUsecase
}

func TestStudentUsecaseStart(t *testing.T) {
	suite.Run(t, new(StudentUsecaseSuite))
}

func (s *StudentUsecaseSuite) StudentUsecaseSuiteDown() {
	defer s.ctrl.Finish()
}

func (s *StudentUsecaseSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	conn, _ := mongo.NewClient()
	s.StudentRepository = mock.NewMockStudentRepository(s.ctrl)
	db := database.NewDatabase(conn, s.StudentRepository)

	s.StudentUsecase = NewStudentUsecase(db)
}
