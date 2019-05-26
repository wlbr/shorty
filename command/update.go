package command

import (
	"github.com/wlbr/shorty/base"
)

// Update is a Command that adds a ShortURL to a store.
type Update struct{}

// Handle is implenting the functionality of the command following the Command design pattern. See above.
func (u *Update) Handle(app *base.Shorty, surl *base.ShortURL, params ...string) error {
	err := app.Store.UpdateShortURL(surl)
	return err
}

// Keyword is the identifier of the command used as a user input token.
func (u *Update) Keyword() string {
	return "u"
}
