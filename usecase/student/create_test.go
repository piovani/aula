package student

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func (s *StudentUsecaseSuite) TestCreateShoudReturnSuccess() {
	fullName := "test"
	age := 10

	s.StudentRepository.EXPECT().Create(gomock.Any()).Return(nil).Times(1)

	student, err := s.StudentUsecase.Create(fullName, age)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), fullName, student.FullName)
	assert.Equal(s.T(), age, student.Age)
}

func (s *StudentUsecaseSuite) TestCreateShoudReturnErr() {
	fullName := "test"
	age := 10
	errExpect := fmt.Errorf("error expect")

	s.StudentRepository.EXPECT().Create(gomock.Any()).Return(errExpect).Times(1)

	student, err := s.StudentUsecase.Create(fullName, age)

	assert.Error(s.T(), err, errExpect.Error())
	assert.Equal(s.T(), fullName, student.FullName)
	assert.Equal(s.T(), age, student.Age)
}
