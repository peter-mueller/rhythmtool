package pages

type HomePage struct {
	Base RhythmPage
}

const (
	createButton = "body /deep/ login-check paper-button"
)

func (homePage *HomePage) CreateRhythm() (CreatePage, error) {
	err := homePage.Base.Chrome.Find(createButton).Click()
	if err != nil {
		return CreatePage{}, err
	}

	homePage.Base.Book.Record("Click the circled plus button in the bottom right corner.")

	return CreatePage{homePage.Base}, nil
}
