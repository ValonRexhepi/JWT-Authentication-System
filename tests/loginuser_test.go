package tests

import (
	"fmt"
	"testing"

	"github.com/ValonRexhepi/JWT-Authentication-System/controllers"
	"github.com/ValonRexhepi/JWT-Authentication-System/models"
)

// TestLoginUserSuccess test the successfull login of an user.
func TestLoginUserSuccess(t *testing.T) {
	controllers.Connect()
	controllers.DB.Exec("DROP TABLE IF EXISTS users")
	controllers.Migrate()

	userToAdd := models.User{
		Name:     "Valon",
		Username: "vrex",
		Email:    "valon-rexhepi@outlook.com",
		Password: " j*Ly^ (=14xh/]}?c",
	}

	controllers.AddUser(&userToAdd)

	loginUserSuccessTests := models.User{
		Name:     "",
		Username: "",
		Email:    "valon-rexhepi@outlook.com",
		Password: " j*Ly^ (=14xh/]}?c",
	}

	err := controllers.LoginUser(&loginUserSuccessTests)

	if err != nil {
		t.Errorf("Expected to successfully login user %v, got %v",
			loginUserSuccessTests, err)
	}
	controllers.DB.Exec("DROP TABLE IF EXISTS users")
}

// TestLoginUserFail test the failed login of an user.
func TestLoginUserFail(t *testing.T) {
	controllers.Connect()
	controllers.DB.Exec("DROP TABLE IF EXISTS users")
	controllers.Migrate()

	userToAdd := models.User{
		Name:     "Valon",
		Username: "vrex",
		Email:    "valon-rexhepi@outlook.com",
		Password: " j*Ly^ (=14xh/]}?c",
	}

	controllers.AddUser(&userToAdd)

	loginUserFail := []models.User{
		{
			Name:     "",
			Username: "",
			Email:    "",
			Password: "",
		},
		{
			Name:     "",
			Username: "",
			Email:    "",
			Password: "j*Ly^(=14xh/]}?c",
		},
		{
			Name:     "",
			Username: "",
			Email:    "valon-rexhepi@outlook.com",
			Password: "",
		},
		{
			Name:     "",
			Username: "",
			Email:    "valon-rexhepi@outlook.com",
			Password: "  Completly  	WrongPassword123",
		},
		{
			Name:     "",
			Username: "",
			Email:    "valon-rexhepi@outlook.com",
			Password: "j*Ly^(=14xh/]}?c",
		},
		{
			Name:     "test",
			Username: "",
			Email:    "mail@doesntexist.com",
			Password: " j*Ly^ (=14xh/]}?c",
		},
	}

	for _, tt := range loginUserFail {
		{
			testname := fmt.Sprintf("TEST:%s, %s", tt.Username, tt.Email)
			t.Run(testname, func(t *testing.T) {
				err := controllers.LoginUser(&tt)
				if err == nil {
					t.Errorf("Expected to fail create user %v, got %v",
						tt, err)
				}
			})
		}
	}
	controllers.DB.Exec("DROP TABLE IF EXISTS users")
}
