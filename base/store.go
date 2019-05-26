package base

import "github.com/wlbr/shorty/gotils"

// Store is an abstracted storage interface. Implementations could use a db, an in memory structure ...
type Store interface {
	Add(localpart string, longURL string, username string) error
	AddShortURL(*ShortURL) error
	List() ([]*ShortURL, error)
	ListByUsername(username string) ([]*ShortURL, error)
	Delete(username string) error
	DeleteShortURL(*ShortURL) error
	Length() (int64, error)
	LengthByUser(username string) (int64, error)
	Update(localpart string, longURL string) error
	UpdateShortURL(*ShortURL) error

	Connect(*gotils.Config)
	Disconnect()
}
