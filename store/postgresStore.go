package store

import (
	"database/sql"
	"fmt"

	//need to import the postgress driver
	_ "github.com/lib/pq"

	"github.com/wlbr/shorty/base"
	"github.com/wlbr/shorty/gotils"
)

// PostgressStore is storing the Shorturls into a Postgress database. Surprise!
type PostgressStore struct {
	db *sql.DB
}

// NewPostgressStore is a constructor for PostgressStore.
func NewPostgressStore() *PostgressStore {
	p := new(PostgressStore)
	return p
}

// Connect connects the store to its implementation. In this case a postgress database.
func (p *PostgressStore) Connect(config *gotils.Config) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.DataBase.User, config.DataBase.Password, config.DataBase.Database)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		gotils.LogFatal("Error openeing database connection: %s", err)
	} else {
		err = db.Ping()
		if err != nil {
			gotils.LogFatal("Error verifying database connection: %s", err)
		}
	}

	p.db = db
}

func (p *PostgressStore) Disconnect() {
	p.db.Close()
}

func (p *PostgressStore) AddShortURL(surl *base.ShortURL) error {
	return p.Add(surl.LocalPart, surl.LongURL, surl.UserName)
}

func (p *PostgressStore) Add(localpart string, longurl string, username string) error {
	tx, err := p.db.Begin()
	_, err = tx.Exec(
		"INSERT INTO shortnames (localpart, longurl, username) VALUES ($1, $2, $3)",
		localpart, longurl, username)
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return err
}

// List retrieves all
// Take care: *FullTableScan*
func (p *PostgressStore) List() ([]*base.ShortURL, error) {
	var results []*base.ShortURL
	rows, err := p.db.Query("SELECT localpart, longurl, username FROM shortnames")
	//defer rows.Close()
	if err != nil {
		gotils.LogError("Error reading list of all shorturls: %v", err)
	} else {
		for rows.Next() {
			s := &base.ShortURL{}
			if e := rows.Scan(&s.LocalPart, &s.LongURL, &s.UserName); e != nil {
				gotils.LogError("Error scanning for shorturl %v", e)
			} else {
				results = append(results, s)
			}
		}
	}
	return results, err
}

func (p *PostgressStore) ListByUsername(username string) ([]*base.ShortURL, error) {
	var results []*base.ShortURL
	rows, err := p.db.Query("SELECT localpart, longurl, username FROM shortnames WHERE username=$1", username)
	//defer rows.Close()
	if err != nil {
		gotils.LogError("Error reading list of all shorturls: %v", err)
	} else {
		for rows.Next() {
			s := &base.ShortURL{}
			if e := rows.Scan(&s.LocalPart, &s.LongURL, &s.UserName); e != nil {
				gotils.LogError("Error scanning for shorturl %v", e)
			} else {
				results = append(results, s)
			}
		}
	}
	return results, err
}

func (p *PostgressStore) UpdateShortURL(newurl *base.ShortURL) error {
	return p.Update(newurl.LocalPart, newurl.LongURL)
}

func (p *PostgressStore) Update(localPart string, longURL string) error {
	tx, err := p.db.Begin()
	_, err = tx.Exec(
		"UPDATE shortnames set longurl=$2 WHERE localpart=$1",
		localPart, longURL)
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return nil
}

func (p *PostgressStore) DeleteShortURL(url *base.ShortURL) error {
	return p.Delete(url.LocalPart)
}

func (p *PostgressStore) Delete(localpart string) error {
	tx, err := p.db.Begin()
	_, err = tx.Exec(
		"DELETE FROM shortnames WHERE localpart=$1",
		localpart)
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return err
}

func (p *PostgressStore) Length() (int64, error) {
	var len int64
	r := p.db.QueryRow("SELECT COUNT(localpart) FROM shortnames")
	err := r.Scan(&len)

	return len, err

}

func (p *PostgressStore) LengthByUser(username string) (int64, error) {
	var len int64
	r := p.db.QueryRow("SELECT COUNT(localpart) FROM shortnames WHERE username=$1", username)
	err := r.Scan(&len)

	return len, err
}
