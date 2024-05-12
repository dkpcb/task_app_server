package model

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func FindUser(u *User) User {
	var user User
	db.Where(u).First(&user)
	return user
}

func AddUser(id uint, name string, password string) (*User, error) {

	user := User{
		ID:       id,
		Name:     name,
		Password: password,
	}

	err := db.Create(&user).Error

	return &user, err
}
