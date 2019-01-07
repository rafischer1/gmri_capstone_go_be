package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/rafischer1/gmri_capstone_go_be/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
