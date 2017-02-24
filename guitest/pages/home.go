package pages

type HomePage struct {
	Base RhythmPage
}

const (
	createButton = "body /deep/ home-page.iron-selected #create"
)

func (homePage *HomePage) CreateRhythm() (CreatePage, error) {
	homePage.Base.Record("Click the circled plus button in the bottom right corner.")
	homePage.Base.Screenshot()

	return CreatePage{homePage.Base}, homePage.Base.Click(createButton)
}
