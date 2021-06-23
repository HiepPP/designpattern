package SOLID

import (
	"encoding/json"
	"errors"
)

type Saver interface {
	SaveNe(data []byte) error
}

type Person struct {
	Name  string
	Phone string
}

func (p *Person) validate() error {
	if p.Name == "" {
		return errors.New("Name missing")
	}
	if p.Phone == "" {
		return errors.New("Phone missing")
	}
	return nil
}

func (p *Person) encode() ([]byte, error) {
	return json.Marshal(p)
}

func SavePerson(person *Person, saver Saver) error {
	err := person.validate()
	if err != nil {
		return err
	}
	bytes, err := person.encode()
	if err != nil {
		return err
	}

	return saver.SaveNe(bytes)
}
