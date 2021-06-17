package creational

import (
	"fmt"
	"testing"
)

func TestBuilderNe(t *testing.T) {
	admin := createUser().CreateName("Administator").CreatePhone("0123213").CreateRole("Admin").BuildUser()
	if admin.getUserRole() == "" {
		t.Error("wrong")
	}
	name := admin.getUserName()
	fmt.Println(name)
	fmt.Printf("%v\n", admin)

	asd := adminne{}.test1().test()
	fmt.Println(asd)
}
