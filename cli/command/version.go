package command

import (
	"fmt"

	myapp "github.com/handlename/my-golang-template"
)

type Version struct{}

func (v *Version) Run(_ *Context) error {
	fmt.Printf("myapp v%s", myapp.Version)
	return nil
}
