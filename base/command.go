package base

// A Command is an instruction used in the API, following the command design pattern.
type Command interface {
	//Help() string
	Handle(app *Shorty, surl ShortURL, params ...string) error
	Keyword() string
}
