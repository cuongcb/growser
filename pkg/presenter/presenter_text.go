package presenter

import "fmt"

type textPresenter struct {
}

func newTextPresenter() *textPresenter {
	return &textPresenter{}
}

func (p *textPresenter) Present(list map[string]string) {
	for k, v := range list {
		fmt.Printf("Project (%s) is at <%s>\n", k, v)
	}
}
