package view

import "fmt"

type textPresenter struct {
}

func newTextPresenter() *textPresenter {
	return &textPresenter{}
}

func (p *textPresenter) Present(list map[string]string) {
	if len(list) == 0 {
		fmt.Println("There is no project yet.")
		return
	}

	for k, v := range list {
		fmt.Printf("Project (%s) is at (%s)\n", k, v)
	}
}
