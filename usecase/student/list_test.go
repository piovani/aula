package student

import (
	"fmt"

	"github.com/piovani/aula/entites"
	"github.com/stretchr/testify/assert"
)

func (s *StudentUsecaseSuite) TestListShoudReturnSuccess() {
	s.StudentRepository.EXPECT().List().Return([]entites.Student{entites.Student{}}, nil)

	list, err := s.StudentUsecase.List()

	assert.NotNil(s.T(), list)
	assert.EqualValues(s.T(), 1, len(list))
	assert.Nil(s.T(), err)
}

func (s *StudentUsecaseSuite) TestListShoudReturnErr() {
	errExpect := fmt.Errorf("error expected")

	s.StudentRepository.EXPECT().List().Return([]entites.Student{entites.Student{}}, errExpect)

	list, err := s.StudentUsecase.List()

	assert.NotNil(s.T(), list)
	assert.EqualValues(s.T(), 1, len(list))
	assert.Error(s.T(), err, errExpect.Error())
}
