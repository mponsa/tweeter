package service

import (
	"fmt"
	"github.com/mponsa/tweeter/src/domain"
	"github.com/mponsa/tweeter/src/errors"
)

type UserManager struct {
	registeredUsers map[string]*domain.User
	loggedUsers map[string]*domain.User
}

var instance *UserManager

func NewUserManager() *UserManager{
	if instance != nil{
		return instance
	}
	instance = new(UserManager)
	instance.registeredUsers = make(map[string]*domain.User)
	instance.loggedUsers = make(map[string]*domain.User)

	return instance
}

func (userManager *UserManager) RegisterUser(name,username,email,password string) *domain.User{
	user := domain.NewUser(name,username,email,password)
	userManager.registeredUsers[user.Username] = user

	return user
}

func (userManager *UserManager) LogIn(username,password string) error{
	user, err := userManager.FindUser(username)
	if err != nil {
		return err
	}
	err = isValidPassword(user,password)
	if err != nil {
		return err
	}
	userManager.loggedUsers[user.Username] = user

	return nil
}

func (userManager *UserManager) FindUser(username string) (*domain.User, error) {
	if user, found := userManager.registeredUsers[username]; found{
		return user, nil
	}

	return nil, fmt.Errorf(errors.ERROR_USER_NOT_REGISTERED)
}

func isValidPassword(user *domain.User, password string) error{
	if user.Password != password {return fmt.Errorf(errors.ERROR_PASSWORD_NOT_VALID)}

	return nil
}

func (userManager *UserManager) IsLoggedIn(username string) bool{
	_, loggedIn := userManager.loggedUsers[username]

	return loggedIn
}

func (userManager *UserManager) LogOut(username string){
	delete(userManager.loggedUsers, username)
}
