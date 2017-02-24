package pages

import (
	"errors"
	"time"

	"github.com/sclevine/agouti"
)

import "github.com/peter-mueller/rhythmtool/guitest/cookbook"

const (
	usernameInput = "body /deep/ login-check paper-input[label=username] input"
	passwordInput = "body /deep/ login-check paper-input[label=password] input"
	submitButton  = "body /deep/ login-check paper-button"
	loginForm     = "body /deep/ login-check  #login"
	logoutButton  = "body /deep/ rhythm-shelf-toolbar  #logout"
)

type RhythmPage struct {
	Chrome *agouti.Page
	Book   cookbook.CookBook
}

func (page *RhythmPage) Screenshot() {
	time.Sleep(time.Second / 2)
	page.Chrome.Screenshot(page.Book.RegisterImage())
}
func (page *RhythmPage) Record(step string) {
	page.Book.Record(step)
}

var ErrButtonNotClickable = errors.New("The element is disabled and cannot be clicked!")

func (page *RhythmPage) Click(selector string) error {
	disabled, err := page.Chrome.Find(selector).Attribute("disabled")
	if err != nil {
		return err
	}
	if disabled == "true" {
		return ErrButtonNotClickable
	}
	return page.Chrome.Find(selector).Click()
}

var ErrAlreadyLoggedIn = errors.New("Cannot log in. Already logged in.")
var ErrNotLoggedIn = errors.New("Cannot log out. Not logged in.")

func (page *RhythmPage) Logout() error {
	page.Record("Log out of the application by clicking the LOGOUT button in the toolbar.")
	if !page.IsLoggedIn() {
		return ErrNotLoggedIn
	}
	return page.Click(logoutButton)
}

func (page *RhythmPage) Login(username, password string) error {

	if page.IsLoggedIn() {
		return ErrAlreadyLoggedIn
	}

	err := page.setUsernameInput(username)
	if err != nil {
		return err
	}

	err = page.setPasswordInput(password)
	if err != nil {
		return err
	}

	page.Record("Log in with your username and password and click the login button.")
	page.Screenshot()

	return page.Click(submitButton)
}

func (page *RhythmPage) setPasswordInput(password string) error {
	err := page.Chrome.Find(passwordInput).Clear()
	if err != nil {
		return err
	}
	err = page.Chrome.Find(passwordInput).SendKeys(password)
	if err != nil {
		return err
	}
	return nil
}

func (page *RhythmPage) setUsernameInput(username string) error {
	err := page.Chrome.Find(usernameInput).Clear()
	if err != nil {
		return err
	}
	err = page.Chrome.Find(usernameInput).SendKeys(username)
	if err != nil {
		return err
	}
	return nil
}

func (page *RhythmPage) IsLoggedIn() bool {
	visible, err := page.Chrome.Find(loginForm).Visible()
	if err != nil {
		return false
	}
	return !visible
}
