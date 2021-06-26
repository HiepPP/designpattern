package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type PersonLoaderMock struct {
	mock.Mock
}

// Giả lập hàm load, không quan tâm kết quả làm này, mục đích là test func LoadPersonName
func (p *PersonLoaderMock) Load(ID int) (*Person, error) {
	// Call kết quả từ hàm Return trong mock
	output := p.Called()

	person := output.Get(0)
	err := output.Error(1)

	if person != nil {
		return person.(*Person), nil
	}

	return nil, err
}

func TestLoadPersonNameMock(t *testing.T) {
	fakeID := 1

	scenarios := []struct {
		desc          string
		configureMock func(mock *PersonLoaderMock)
		expectedName  string
		expectErr     bool
	}{
		{
			desc: "Happy path",
			configureMock: func(loaderMock *PersonLoaderMock){
				loaderMock.On("Load", mock.Anything).Return(&Person{Name: "Hieppp"}, nil).Once()
			},
			expectedName: "Hieppp",
			expectErr: false,
		},
		{
			desc: "input error",
			configureMock: func(loaderMock *PersonLoaderMock){
				loaderMock.On("Load", mock.Anything).Return(nil, ErrNotFound).Once()
			},
			expectedName: "",
			expectErr: true,
		},
		{
			desc: "system error path",
			configureMock: func(loaderMock *PersonLoaderMock) {
				loaderMock.On("Load", mock.Anything).Return(nil, errors.New("something failed")).Once()
			},
			expectedName: "",
			expectErr: true,
		},
	}

	for _, item:= range scenarios{
		mockLoader := &PersonLoaderMock{}
		item.configureMock(mockLoader)

		resultName, resultError := LoadPersonName(mockLoader, fakeID)
		assert.Equal(t, item.expectedName, resultName, item.desc)
		assert.Equal(t, item.expectErr, resultError != nil, item.desc)
		assert.True(t, mockLoader.AssertExpectations(t), item.desc)
	}
}
