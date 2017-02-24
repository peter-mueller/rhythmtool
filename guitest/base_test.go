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

	homePage.Base.Record("Now you are on the home page.")
	homePage.Base.Screenshot()
	homePage.Base.Book.ClearToFile()
}

func TestUserLogout(t *testing.T) {
	t.Parallel()
	homePage := newPage("user-logout")
	defer homePage.Base.Chrome.CloseWindow()

	if err := homePage.Base.Login("admin", "admin"); err != nil {
		t.Fatal(err)
	}
	homePage.Base.Record("Now you are on the home page.")
	homePage.Base.Screenshot()
	time.Sleep(time.Second)
	if err := homePage.Base.Logout(); err != nil {
		t.Fatal(err)
	}

	homePage.Base.Record("You are now logged out of the application")
	homePage.Base.Screenshot()
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
	homePage.Base.Record("A toast appears, if the credentials were wrong.")
	homePage.Base.Screenshot()
	homePage.Base.Book.ClearToFile()
}
