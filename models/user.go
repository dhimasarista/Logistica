package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) FindAll() []User {
	var users = []User{
		{
			Username: "dhimasarista",
			Password: "01052002",
		},
		{
			Username: "alexanto",
			Password: "2219770213",
		},
	}
	return users
}
