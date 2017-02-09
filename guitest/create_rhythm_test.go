package guitest_test

import "testing"

func TestCreateRhythmByText(t *testing.T) {
	t.Parallel()
	homePage := newPage("create-rhyhtm-by-text")
	defer homePage.Base.Chrome.CloseWindow()

	if err := homePage.Base.Login("admin", "admin"); err != nil {
		t.Fatal(err)
	}
	homePage.CreateRhythm()
	homePage.Base.Book.ClearToFile()
}
