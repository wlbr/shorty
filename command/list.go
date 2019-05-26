package command

import (
	"fmt"

	"github.com/wlbr/shorty/base"
)

// List is a Command that adds a ShortURL to a store.
type List struct{}

// Handle is implenting the functionality of the command following the Command design pattern. See above.
func (l *List) Handle(app *base.Shorty, surl base.ShortURL, params ...string) error {
	fmt.Println(app.Store.List())
	return nil
}

// Keyword is the identifier of the command used as a user input token.
func (l *List) Keyword() string {
	return "l"
}
