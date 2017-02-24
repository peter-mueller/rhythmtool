package pages

import "strconv"

type CreateRandomPage struct {
	Base RhythmPage
}

const (
	nameInput        = "body /deep/ create-page.iron-selected rhythm-random-input paper-input#nameInput input"
	lengthInput      = "body /deep/ create-page.iron-selected rhythm-random-input paper-slider paper-input input"
	rerollButton     = "body /deep/ create-page.iron-selected rhythm-random-input paper-button#random"
	saveRandomButton = "body /deep/ create-page.iron-selected rhythm-random-input #controls paper-button#save"
)

func (page *CreateRandomPage) SetName(name string) error {
	page.Base.Record("Give the new rhythm a name by typing it into the input field.")
	defer page.Base.Screenshot()
	return page.Base.Chrome.Find(nameInput).SendKeys(name)
}

func (page *CreateRandomPage) SetLength(length int) error {
	page.Base.Record("Specify the length of the rhythm by adjusting the slider or typing it into the input field.")
	defer page.Base.Screenshot()
	if err := page.Base.Chrome.Find(lengthInput).Clear(); err != nil {
		return err
	}
	return page.Base.Chrome.Find(lengthInput).SendKeys(strconv.Itoa(length))
}

func (page *CreateRandomPage) Reroll() error {
	page.Base.Record("Make a new random rhythm by clicking the REROLL button.")
	defer page.Base.Screenshot()
	return page.Base.Click(rerollButton)
}

func (page *CreateRandomPage) Save() (HomePage, error) {
	page.Base.Record("Save the random rhythm by clicking the SAVE button.")
	defer page.Base.Screenshot()
	return HomePage{page.Base}, page.Base.Click(saveRandomButton)
}
