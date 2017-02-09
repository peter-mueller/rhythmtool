package pages

type HomePage struct {
	Base RhythmPage
}

const (
	createButton = "body /deep/ home-page #create"
)

func (homePage *HomePage) CreateRhythm() (CreatePage, error) {
	homePage.Base.Record("Click the circled plus button in the bottom right corner.")
	homePage.Base.Screenshot()
	err := homePage.Base.Chrome.Find(createButton).Click()
	if err != nil {
		return CreatePage{}, err
	}
	return CreatePage{homePage.Base}, nil
}
