package view

// Presenter displays project list in a choosen format
type Presenter interface {
	Present(list map[string]string)
}

// NewPresenter returns a new Presenter
func NewPresenter() Presenter {
	return newTextPresenter()
}
