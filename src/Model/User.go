package Model

import (
	"github.com/amirhossein2831/DummyGo/src/App"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func GetAllUser() ([]User, error) {
	var users []User
	res := App.App.DB.Find(&users)
	if res.Error != nil {
		return []User{}, res.Error
	}
	return users, nil
}

func GetUser(userID string) (User, error) {
	user := User{}
	res := App.App.DB.First(&user, userID)
	if res.Error != nil {
		return User{}, res.Error
	}
	return user, nil
}
