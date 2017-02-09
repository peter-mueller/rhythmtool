package guitest_test

import "testing"

func TestCreateRhythmByText(t *testing.T) {
	t.Parallel()

	homePage := newPage("create-rhythm-by-text")
	defer homePage.Base.Chrome.CloseWindow()

	if err := homePage.Base.Login("admin", "admin"); err != nil {
		t.Fatal(err)
	}
	createPage, err := homePage.CreateRhythm()
	if err != nil {
		t.Fatal(err)
	}
	createRandomPage, err := createPage.SelectRandom()
	if err != nil {
		t.Fatal(err)
	}
	createRandomPage.Base.Book.ClearToFile()
}
