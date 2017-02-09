package pages

import "time"

type CreatePage struct {
	Base RhythmPage
}

const (
	selectGroup     = "body /deep/ create-page paper-radio-group"
	selectText      = selectGroup + " paper-radio-button[name=text]"
	selectBjorklund = selectGroup + " paper-radio-button[name=bjorklund]"
	selectRandom    = selectGroup + " paper-radio-button[name=random]"
	continueButton  = "body /deep/ create-page #continue"
)

func (createPage *CreatePage) SelectText() (CreateTextPage, error) {
	createPage.Base.Record("Select the text option from the radio button group and click the continue button.")
	if err := createPage.Base.Chrome.Find(selectText).Click(); err != nil {
		return CreateTextPage{}, err
	}
	err := createPage.submit()
	return CreateTextPage{createPage.Base}, err
}

func (createPage *CreatePage) SelectBjorklund() (CreateBjorklundPage, error) {
	createPage.Base.Record("Select the Bjorklund algorithm option from the radio button group and click the continue buttton.")
	if err := createPage.Base.Chrome.Find(selectBjorklund).Click(); err != nil {
		return CreateBjorklundPage{}, err
	}
	err := createPage.submit()
	return CreateBjorklundPage{createPage.Base}, err
}

func (createPage *CreatePage) SelectRandom() (CreateRandomPage, error) {
	createPage.Base.Record("Select the random option from the radio button group and click the continue button.")
	if err := createPage.Base.Chrome.Find(selectRandom).Click(); err != nil {
		return CreateRandomPage{}, err
	}
	err := createPage.submit()
	return CreateRandomPage{createPage.Base}, err
}

func (createPage *CreatePage) submit() error {
	time.Sleep(time.Second)
	createPage.Base.Screenshot()
	err := createPage.Base.Chrome.Find(continueButton).Click()
	return err
}

type CreateTextPage struct {
	Base RhythmPage
}

type CreateRandomPage struct {
	Base RhythmPage
}

type CreateBjorklundPage struct {
	Base RhythmPage
}
