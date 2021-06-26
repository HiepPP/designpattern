package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadPersonName(t *testing.T) {
	fakeId := 1
	scenarios := []struct {
		desc         string
		loaderStub   *PersonLoaderStub
		expectedName string
		expectErr    bool
	}{
		{
			desc: "Happy path",
			loaderStub: &PersonLoaderStub{
				Person: &Person{Name: "Hieppp"},
			},
			expectedName: "Hieppp",
			expectErr:    false,
		},
		{
			desc: "Input error",
			loaderStub: &PersonLoaderStub{
				Error: ErrNotFound,
			},
			expectedName: "",
			expectErr:    true,
		},
		{
			desc: "system error path",
			loaderStub: &PersonLoaderStub{
				Error: errors.New("something failed"),
			},
			expectedName: "",
			expectErr:    true,
		},
	}
	for _, item := range scenarios {
		resultName, resultError := LoadPersonName(item.loaderStub, fakeId)

		assert.Equal(t, item.expectedName, resultName, item.desc)
		assert.Equal(t, item.expectErr, resultError != nil, item.desc)
	}
}
