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
	createTextPage, err := createPage.SelectText()
	if err != nil {
		t.Fatal(err)
	}
	err = createTextPage.SetText("texttomakerhythm")
	if err != nil {
		t.Fatal(err)
	}
	backHomePage, err := createTextPage.Save()
	if err != nil {
		t.Fatal(err)
	}

	backHomePage.Base.Record("You are left back at the home page. The new saved rhythm is now in the list.")
	backHomePage.Base.Screenshot()

	backHomePage.Base.Book.ClearToFile()
}

func TestCreateRhythmByRandom(t *testing.T) {
	t.Parallel()

	homePage := newPage("create-rhythm-by-random")
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

	if err := createRandomPage.SetName("RandomRhythm1"); err != nil {
		t.Fatal(err)
	}
	if err := createRandomPage.SetLength(8); err != nil {
		t.Fatal(err)
	}
	if err := createRandomPage.Reroll(); err != nil {
		t.Fatal(err)
	}

	backHomePage, err := createRandomPage.Save()
	if err != nil {
		t.Fatal(err)
	}
	backHomePage.Base.Book.ClearToFile()
}
