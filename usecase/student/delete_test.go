package student

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/piovani/aula/entites"
	"github.com/stretchr/testify/assert"
)

func (s *StudentUsecaseSuite) TestDeleteShoudReturnSuccess() {
	id := uuid.New()

	s.StudentRepository.EXPECT().FindByID(id).Return(entites.Student{ID: id}, nil).Times(1)
	s.StudentRepository.EXPECT().Delete(id).Return(nil).Times(1)

	err := s.StudentUsecase.Delete(id)

	assert.Nil(s.T(), err)
}

func (s *StudentUsecaseSuite) TestDeleteShoudReturnErrFindByID() {
	id := uuid.New()
	errExpect := fmt.Errorf("error expected")

	s.StudentRepository.EXPECT().FindByID(id).Return(entites.Student{ID: id}, errExpect).Times(1)

	err := s.StudentUsecase.Delete(id)

	assert.Error(s.T(), err, errExpect.Error())
}

func (s *StudentUsecaseSuite) TestDeleteShoudReturnErrDelete() {
	id := uuid.New()
	errExpect := fmt.Errorf("error expected")

	s.StudentRepository.EXPECT().FindByID(id).Return(entites.Student{ID: id}, nil).Times(1)
	s.StudentRepository.EXPECT().Delete(id).Return(errExpect).Times(1)

	err := s.StudentUsecase.Delete(id)

	assert.Error(s.T(), err, errExpect.Error())
}

func (s *StudentUsecaseSuite) TestDeleteShoudReturnErrIDEmpty() {
	id := uuid.New()

	s.StudentRepository.EXPECT().FindByID(id).Return(entites.Student{}, nil).Times(1)

	err := s.StudentUsecase.Delete(id)

	assert.Error(s.T(), err, "Não foi possivel encontrar o estudante")
}