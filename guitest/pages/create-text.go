package pages

type CreateTextPage struct {
	Base RhythmPage
}

const (
	textInput      = "body /deep/ create-page.iron-selected rhythm-text-input paper-input input"
	saveTextButton = "body /deep/ create-page.iron-selected rhythm-text-input #controls paper-button#save"
)

func (page *CreateTextPage) SetText(text string) error {
	page.Base.Record("Input the text to use for the rhythm into the text field.")
	defer page.Base.Screenshot()

	return page.Base.Chrome.Find(textInput).SendKeys(text)
}

func (page *CreateTextPage) Save() (HomePage, error) {
	page.Base.Record("Save the rhythm by clicking the SAVE button.")
	return HomePage{page.Base}, page.Base.Click(saveTextButton)
}
