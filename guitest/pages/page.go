package pages

import "github.com/sclevine/agouti"
import "errors"
import "github.com/peter-mueller/rhythmtool/guitest/cookbook"

const (
	usernameInput = "body /deep/ login-check paper-input[label=username] input"
	passwordInput = "body /deep/ login-check paper-input[label=password] input"
	submitButton  = "body /deep/ login-check paper-button"
	loginForm     = "body /deep/ login-check  #login"
)

type RhythmPage struct {
	Chrome *agouti.Page
	Book   cookbook.CookBook
}

func (page *RhythmPage) Login(username, password string) error {

	if page.IsLoggedIn() {
		return errors.New("Already logged in!")
	}

	err := page.setUsernameInput(username)
	if err != nil {
		return err
	}

	err = page.setPasswordInput(password)
	if err != nil {
		return err
	}

	page.Book.Record("Log in with your username and password and click the login button.")
	page.Chrome.Screenshot(page.Book.RegisterImage())

	err = page.Chrome.Find(submitButton).Click()
	if err != nil {
		return err
	}
	return nil
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
