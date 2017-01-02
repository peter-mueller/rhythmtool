package server_test

import (
	"testing"
	"time"

	"github.com/sclevine/agouti"
)

func TestUserLogin(t *testing.T) {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		t.Fatal("Failed to start Selenium:", err)
	}
	defer driver.Stop()
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		t.Fatal("Failed to open page:", err)
	}

	//s.StartServer("8081")
	if err := page.Navigate("http://localhost:8080"); err != nil {
		t.Fatal("Failed to navigate:", err)
	}

	time.Sleep(time.Second * 1)
	err = page.Find("body /deep/ login-check paper-input[label=username] input").SendKeys("admin")
	err = page.Find("body /deep/ login-check paper-input[label=password] input").SendKeys("wrong")
	page.Screenshot("./screenshots/login.png")
	err = page.Find("body /deep/ login-check paper-button").Click()
	time.Sleep(1 * time.Second)
	page.Screenshot("./screenshots/login-failed.png")
	time.Sleep(2 * time.Second)

	err = page.Find("body /deep/ login-check paper-input[label=password] input").Clear()
	err = page.Find("body /deep/ login-check paper-input[label=password] input").SendKeys("admin")

	err = page.Find("body /deep/ login-check paper-button").Click()

	page.Screenshot("./screenshots/home.png")

	if err != nil {
		t.Fatal("Failed to get login prompt text:", err)
	}

	time.Sleep(time.Second * 2)

}
