package browser

import "os/exec"

// Browser moves/open a directory which specificed in path
type Browser interface {
	Go(string) error
}

// New returns a new browser
func New() Browser {
	return newUnixBrowser()
}

type unixBrowser struct {
}

func newUnixBrowser() *unixBrowser {
	return &unixBrowser{}
}

func (u *unixBrowser) Go(path string) error {
	cmd := exec.Command("gnome-terminal", "--tab")
	cmd.Dir = path
	return cmd.Start()
}
