package base

import "fmt"

// ShortURL is the central datatype of the application. It is used to
// handle the redirect URLs and to pass it to the stores
type ShortURL struct {
	LocalPart string
	LongURL   string
	UserName  string
}

func (s ShortURL) String() string {
	return fmt.Sprintf("[%s] %s --> %s", s.UserName, s.LocalPart, s.LongURL)
}
