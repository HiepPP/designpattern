package test

import "errors"

// Còn gọi là test double
// Là fake implementations của dependency (là interface)

var ErrNotFound = errors.New("person not found")

type PersonLoader interface {
	Load(ID int) (*Person, error)
}

type Person struct {
	Name string
}

func LoadPersonName(loader PersonLoader, ID int) (string, error) {
	// test không quan tâm kết quả trả về
	person, err := loader.Load(ID)
	if err != nil {
		return "", err
	}

	return person.Name, nil
}

// Giả lập kết quả trả về, implement interface loader, đóng vai trò là một loader khi truyền vào LoadPersonName
type PersonLoaderStub struct {
	Person *Person
	Error  error
}

// Giả lập hàm gọi kết quả trả về
func (p *PersonLoaderStub) Load(ID int) (*Person, error) {
	return p.Person, p.Error
}
