package service_test

import (
	"github.com/mponsa/tweeter/src/domain"
	"github.com/mponsa/tweeter/src/service"
	"testing"
)

func TestUserIsRegistered(t *testing.T){

	//Initialization
	var user *domain.User
	name := "Manuel Ponsa"
	username := "mponsa"
	email := "manuponsa@gmail.com"
	password := "blabla23"
	userManager := service.NewUserManager()

	//Operation
	user = userManager.RegisterUser(name,username,email,password)

	//Validation
	if user.Name != name{
		t.Errorf("Expected name is %s , but received %s ", name, user.Name )
	}

	if user.Username != username{
		t.Errorf("Expected username is %s , but received %s ", username, user.Username )
	}

	if user.Email != email{
		t.Errorf("Expected Email is %s , but received %s ", email, user.Email )
	}

	if user.Password != password{
		t.Errorf("Expected Password is %s , but received %s ", password, user.Password )
	}
}

func TestUserLogsIn(t *testing.T){
	//Initialization
	name := "manuel"
	username := "mponsa"
	email := "manuponsa@gmail.com"
	password := "blabla23"
	userManager := service.NewUserManager()

	//Operacion
	_ = userManager.RegisterUser(name, username, email, password)
	_ = userManager.LogIn(username,password)

	//Validacion
	if(!userManager.IsLoggedIn(username)){
		t.Errorf("User should be loggedin, but received")
	}
}

func TestUserCantLoginWithWrongPassword(t *testing.T){
	//Initialization
	name := "manuel"
	username := "mponsa"
	email := "manuponsa@gmail.com"
	password := "blabla23"
	userManager := service.NewUserManager()
	wrong_password := "baba23"

	//Operation
	_ = userManager.RegisterUser(name, username, email, password)
	err := userManager.LogIn(username,wrong_password)

	//Validation
	if err!= nil && err.Error() != "not a valid password"{
		t.Errorf("Expected error was not a valid password but got %s",err.Error())
	}
}

func TestUserCantLogInIfIsNotRegistered(t *testing.T){
	//Initialization
	username := "mponsa"
	password := "blabla23"
	userManager := service.NewUserManager()
	//Operation
	err := userManager.LogIn(username,password)

	//Validation
	if err!= nil && err.Error() != "user not registered"{
		t.Errorf("Expected error was not a valid password but got %s",err.Error())
	}
}

func TestUserCanLogout(t *testing.T){
	//Initialization
	name := "manuel"
	username := "mponsa"
	email := "manuponsa@gmail.com"
	password := "blabla23"
	userManager := service.NewUserManager()
	_ = userManager.RegisterUser(name, username, email, password)
	_ = userManager.LogIn(username,password)

	//Operation
	userManager.LogOut(username)

	//Validation
	if userManager.IsLoggedIn(username) == true{
		t.Errorf("User shouldn't be logged in")
	}
}