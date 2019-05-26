package command

import (
	"github.com/wlbr/shorty/base"
)

// Delete is a Command that adds a ShortURL to a store.
type Delete struct{}

// Handle is implenting the functionality of the command following the Command design pattern. See above.
func (d *Delete) Handle(app *base.Shorty, surl *base.ShortURL, params ...string) error {
	err := app.Store.DeleteShortURL(surl)
	return err
}

// Keyword is the identifier of the command used as a user input token.
func (d *Delete) Keyword() string {
	return "d"
}
