package pages

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
	if err := createPage.Base.Click(selectText); err != nil {
		return CreateTextPage{}, err
	}
	err := createPage.submit()
	return CreateTextPage{createPage.Base}, err
}

func (createPage *CreatePage) SelectBjorklund() (CreateBjorklundPage, error) {
	createPage.Base.Record("Select the Bjorklund algorithm option from the radio button group and click the continue buttton.")
	if err := createPage.Base.Click(selectBjorklund); err != nil {
		return CreateBjorklundPage{}, err
	}
	err := createPage.submit()
	return CreateBjorklundPage{createPage.Base}, err
}

func (createPage *CreatePage) SelectRandom() (CreateRandomPage, error) {
	createPage.Base.Record("Select the random option from the radio button group and click the continue button.")
	if err := createPage.Base.Click(selectRandom); err != nil {
		return CreateRandomPage{}, err
	}
	return CreateRandomPage{createPage.Base}, createPage.submit()
}

func (createPage *CreatePage) submit() error {
	createPage.Base.Screenshot()
	return createPage.Base.Click(continueButton)
}

type CreateBjorklundPage struct {
	Base RhythmPage
}
