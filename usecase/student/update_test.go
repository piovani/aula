package student

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/piovani/aula/entites"
	"github.com/stretchr/testify/assert"
)

func (s *StudentUsecaseSuite) TestUpdateShoudReturnSuccess() {
	id := uuid.New()
	fullName := "test"
	age := 10

	s.StudentRepository.EXPECT().FindByID(id).Return(entites.Student{ID: id}, nil).Times(1)
	s.StudentRepository.EXPECT().Update(gomock.Any()).Return(nil).Times(1)

	student, err := s.StudentUsecase.Update(id, fullName, age)

	assert.NotNil(s.T(), student)
	assert.EqualValues(s.T(), id, student.ID)
	assert.EqualValues(s.T(), fullName, student.FullName)
	assert.EqualValues(s.T(), age, student.Age)
	assert.Nil(s.T(), err)
}

func (s *StudentUsecaseSuite) TestUpdateShoudReturnErrFindByID() {
	id := uuid.New()
	fullName := "test"
	age := 10
	errExpect := fmt.Errorf("error expected")

	s.StudentRepository.EXPECT().FindByID(id).Return(entites.Student{ID: id}, errExpect).Times(1)

	student, err := s.StudentUsecase.Update(id, fullName, age)

	assert.NotNil(s.T(), student)
	assert.Error(s.T(), err, errExpect.Error())
}

func (s *StudentUsecaseSuite) TestUpdateShoudReturnErrIDEmpty() {
	id := uuid.New()
	fullName := "test"
	age := 10

	s.StudentRepository.EXPECT().FindByID(id).Return(entites.Student{}, nil).Times(1)

	student, err := s.StudentUsecase.Update(id, fullName, age)

	assert.NotNil(s.T(), student)
	assert.Error(s.T(), err, "Não foi possivel encontrar o estudante")
}

func (s *StudentUsecaseSuite) TestUpdateShoudReturnErrUpdate() {
	id := uuid.New()
	fullName := "test"
	age := 10
	errExpect := fmt.Errorf("error expected")

	s.StudentRepository.EXPECT().FindByID(id).Return(entites.Student{ID: id}, nil).Times(1)
	s.StudentRepository.EXPECT().Update(gomock.Any()).Return(errExpect).Times(1)

	student, err := s.StudentUsecase.Update(id, fullName, age)

	assert.NotNil(s.T(), student)
	assert.EqualValues(s.T(), id, student.ID)
	assert.EqualValues(s.T(), fullName, student.FullName)
	assert.EqualValues(s.T(), age, student.Age)
	assert.Error(s.T(), err, errExpect.Error())
}
