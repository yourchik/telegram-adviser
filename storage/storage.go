package storage

import (
	"crypto/sha1"
	"fmt"
	"io"
	"telegram-adviser/lib/customError"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(username string) (*Page, error)
	Remove(p *Page) error
	IsExist(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
	//Created time.Time
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", customError.Wrap("cannot calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", customError.Wrap("cannot calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
