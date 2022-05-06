package tests

import (
	"fmt"
	"testing"

	"github.com/ValonRexhepi/JWT-Authentication-System/controllers"
	"github.com/ValonRexhepi/JWT-Authentication-System/models"
)

// TestAddUserSuccess test the successfull addition of a new user.
func TestAddUserSuccess(t *testing.T) {
	controllers.Connect()
	controllers.DB.Exec("DROP TABLE IF EXISTS users")
	controllers.Migrate()

	addUserSuccessTests := []models.User{
		{
			Name:     "Valon",
			Username: "vrex",
			Email:    "valon-rexhepi@outlook.com",
			Password: " mXAG1BR]OwQqcD4",
		},
		{
			Name:     "Harry Potter",
			Username: "harrpott",
			Email:    "harry-potter@outlook.com",
			Password: "[z/Y;Ak)1x()gh&x",
		},
		{
			Name:     "Hermione Granger",
			Username: "germgrang",
			Email:    "hermione-granger@outlook.com",
			Password: "j*Ly^(=14xh/]}?c",
		},
	}

	for _, tt := range addUserSuccessTests {
		{
			testname := fmt.Sprintf("TEST:%s, %s", tt.Username, tt.Email)
			t.Run(testname, func(t *testing.T) {
				id, err := controllers.AddUser(&tt)
				if id == -1 || err != nil {
					t.Errorf("Expected to successfully create user %v, got %v",
						tt, err)
				}
			})
		}
	}
	controllers.DB.Exec("DROP TABLE IF EXISTS users")
}

// TestAddUserFail test the failed addition of a new user.
func TestAddUserFail(t *testing.T) {
	controllers.Connect()
	controllers.DB.Exec("DROP TABLE IF EXISTS users")
	controllers.Migrate()

	userToAdd := models.User{
		Name:     "Valon",
		Username: "vrex",
		Email:    "valon-rexhepi@outlook.com",
		Password: "MyIncrediblePassword",
	}

	controllers.AddUser(&userToAdd)

	addUserFail := []models.User{
		{
			Name:     "test",
			Username: "",
			Email:    "",
			Password: "",
		},
		{
			Name:     "",
			Username: "test",
			Email:    "",
			Password: "",
		},
		{
			Name:     "",
			Username: "",
			Email:    "test",
			Password: "",
		},
		{
			Name:     "",
			Username: "",
			Email:    "",
			Password: "test",
		},
		{
			Name:     "test",
			Username: "test",
			Email:    "",
			Password: "",
		},
		{
			Name:     "test",
			Username: "",
			Email:    "test",
			Password: "",
		},
		{
			Name:     "test",
			Username: "",
			Email:    "",
			Password: "test",
		},
		{
			Name:     "",
			Username: "test",
			Email:    "test",
			Password: "",
		},
		{
			Name:     "",
			Username: "test",
			Email:    "",
			Password: "test",
		},
		{
			Name:     "",
			Username: "",
			Email:    "test",
			Password: "test",
		},
		{
			Name:     "test",
			Username: "test",
			Email:    "test",
			Password: "",
		},
		{
			Name:     "test",
			Username: "test",
			Email:    "",
			Password: "test",
		},
		{
			Name:     "test",
			Username: "",
			Email:    "test",
			Password: "test",
		},
		{
			Name:     "",
			Username: "test",
			Email:    "test",
			Password: "test",
		},
		{
			Name:     "Harry Potter",
			Username: "vrex",
			Email:    "harry-potter@outlook.com",
			Password: "j*Ly^(=14xh/]}?c",
		},
		{
			Name:     "Harry Potter",
			Username: "val rex",
			Email:    "valon-rexhepi@outlook.com",
			Password: "j*Ly^(=14xh/]}?c",
		},
		{
			Name:     "First User",
			Username: "firstuser",
			Email:    "first-user@outlook.com",
			Password: `badentropypassword`,
		},
	}

	for _, tt := range addUserFail {
		{
			testname := fmt.Sprintf("TEST:%s, %s", tt.Username, tt.Email)
			t.Run(testname, func(t *testing.T) {
				id, err := controllers.AddUser(&tt)
				if id != -1 || err == nil {
					t.Errorf("Expected to fail create user %v, got %v",
						tt, err)
				}
			})
		}
	}
	controllers.DB.Exec("DROP TABLE IF EXISTS users")
}
