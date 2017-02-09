package guitest_test

import (
	"os"
	"sync"
	"testing"

	"github.com/peter-mueller/rhythmtool/guitest/cookbook"
	"github.com/peter-mueller/rhythmtool/guitest/pages"
	"github.com/sclevine/agouti"
)

func TestMain(m *testing.M) {
	os.Exit(setup(m))
}

var driver *agouti.WebDriver

//go:generate go test
//go:generate asciidoctor doc.adoc
func setup(m *testing.M) int {
	driver = agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		panic("Failed to start Selenium!")
	}
	defer driver.Stop()

	clearFolders()

	return m.Run()
}

func clearFolders() {
	os.Mkdir("recipes", 0777)
	os.Mkdir("screenshots", 0777)
}

var driverMutex sync.Mutex

func newPage(name string) pages.HomePage {
	driverMutex.Lock()
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		panic("Failed to open page!")
	}
	if err := page.Navigate("http://localhost:8081/app/"); err != nil {
		panic("Failed to navigate!")
	}
	book := cookbook.New(name)
	driverMutex.Unlock()
	return pages.HomePage{Base: pages.RhythmPage{Chrome: page, Book: book}}
}
