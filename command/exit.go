package command

import (
	"fmt"
	"os"

	"github.com/wlbr/shorty/base"
)

// Exit Add is a Command that adds a ShortURL to a store.
type Exit struct{}

// Handle is implenting the functionality of the command following the Command design pattern. See above.
func (e *Exit) Handle(app *base.Shorty, surl base.ShortURL, params ...string) error {

	fmt.Println("bye bye")
	os.Exit(0)
	return nil
}

// Keyword is the identifier of the command used as a user input token.
func (e *Exit) Keyword() string {
	return "x"
}
