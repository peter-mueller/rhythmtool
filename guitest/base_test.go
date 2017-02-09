package guitest_test

import (
	"testing"
	"time"
)

func TestUserLogin(t *testing.T) {
	t.Parallel()
	homePage := newPage("user-login")
	defer homePage.Base.Chrome.CloseWindow()

	if err := homePage.Base.Login("admin", "admin"); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)

	if !homePage.Base.IsLoggedIn() {
		t.Fatal("Should be logged in!")
	}

	homePage.Base.Book.Record("Now you are on the home page.")
	homePage.Base.Chrome.Screenshot(homePage.Base.Book.RegisterImage())

	homePage.Base.Book.ClearToFile()
}

func TestUserLoginFail(t *testing.T) {
	t.Parallel()
	homePage := newPage("user-login-fail")
	defer homePage.Base.Chrome.CloseWindow()
	if err := homePage.Base.Login("admin", "wrongpassword"); err != nil {
		t.Fatal(err)
	}
	if homePage.Base.IsLoggedIn() {
		t.Fatal("Should not be logged in!")
	}

	time.Sleep(time.Second)
	homePage.Base.Book.Record("A toast appears, if the credentials were wrong.")
	homePage.Base.Chrome.Screenshot(homePage.Base.Book.RegisterImage())

	homePage.Base.Book.ClearToFile()

}
