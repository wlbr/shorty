package command

import (
	"github.com/wlbr/shorty/base"
)

// Add is a Command that adds a ShortURL to a store.
type Add struct{}

// Handle is implenting the functionality of the command following the Command design pattern. See above.
func (a *Add) Handle(app *base.Shorty, surl *base.ShortURL, params ...string) error {
	app.Store.AddShortURL(surl)
	return nil
}

// Keyword is the identifier of the command used as a user input token.
func (a *Add) Keyword() string {
	return "a"
}
