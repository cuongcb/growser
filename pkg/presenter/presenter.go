package presenter

// Presenter displays project list in a choosen format
type Presenter interface {
	Present(list map[string]string)
}

// New returns a new Presenter
func New() Presenter {
	return newTextPresenter()
}
