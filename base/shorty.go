package base

import (
	"strings"
)

// Shorty is the main app object, the root of it all.
type Shorty struct {
	Commands        []Command
	commandHandlers map[string][]Command
	Store           Store
}

// RegisterCommand is add a new Command to the list of commands following the Listener/Command design pattern.
func (t *Shorty) RegisterCommand(s Command) {
	t.Commands = append(t.Commands, s)
	if nil == t.commandHandlers {
		t.commandHandlers = make(map[string][]Command)
	}
	t.commandHandlers[s.Keyword()] = append(t.commandHandlers[s.Keyword()], s)
}

// DispatchCommand is commencing an event, in other word it starts all Commands listening to the input keyword.
func (t *Shorty) DispatchCommand(input string) (e error) {
	kword := strings.Fields(input)[0]
	//text := strings.TrimLeft(strings.TrimPrefix(input, kword), " ")

	handlers := t.commandHandlers[kword]
	for _, h := range handlers {
		e = h.Handle(t, ShortURL{})
	}
	return e
}
