package builder

type user struct {
	name  string
	phone string
	role  string
}

// Nếu user implement interface này, mọi property sẽ được che giấu, không thể gọi trực tiếp mà phải thông qua function get
type UserInformation interface {
	getUserName() string
	getUserPhone() string
	getUserRole() string
}

type UserBuilder interface {
	CreateName(name string) UserBuilder
	CreatePhone(phone string) UserBuilder
	CreateRole(role string) UserBuilder
	BuildUser() UserInformation // không return về trực tiếp struct user mà thông qua interface để che giấu property
}

func (u *user) getUserName() string {
	return u.name
}
func (u *user) getUserPhone() string {
	return u.phone
}
func (u *user) getUserRole() string {
	return u.role
}

func (u *user) CreateName(name string) UserBuilder { //return về interface UserBuilder có thể tiếp tục call func thuộc về interface này
	u.name = name
	return u
}

func (u *user) CreatePhone(phone string) UserBuilder {
	u.phone = phone
	return u
}

func (u *user) CreateRole(role string) UserBuilder {
	u.role = role
	return u
}

func (u *user) BuildUser() UserInformation {
	return u
}

func createUser() *user {
	return new(user)
}

type adminne struct {
	name string
}

func (a adminne) test() adminne {
	a.name = "123"
	return a
}

func (a adminne) test1() adminne{
	a.name = "asdasd"
	return a
}

//func newUser(name string, phone string) *user {
//	return &user{name: name, phone: phone}
//}
