package storage

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