package SOLID

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSaver struct {
	mock.Mock
}

func (m *mockSaver) SaveNe(data []byte) error {
	output := m.Called(data)
	return output.Error(0)
}

func TestSaverPass(t *testing.T) {
	in := &Person{
		Name:  "HiepPP",
		Phone: "123123123",
	}

	mockNFS := new(mockSaver)
	mockNFS.On("SaveNe", mock.Anything).Return(nil).Once()

	resultError := SavePerson(in, mockNFS)

	assert.NoError(t, resultError)
	assert.True(t, mockNFS.AssertExpectations(t))
}

func TestSaveFail(t *testing.T) {
	in := &Person{
		Name:  "HiepPL",
		Phone: "123213123",
	}

	mockNFS := new(mockSaver)
	mockNFS.On("SaveNe", mock.Anything).Return(errors.New("save failed")).Once()

	resultError := SavePerson(in, mockNFS)
	assert.Error(t, resultError)
	assert.True(t, mockNFS.AssertExpectations(t))

}
