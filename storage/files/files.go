package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"telegram-adviser/lib/customError"
	"telegram-adviser/storage"
	"time"
)

const defaultPerm = 0755

type Storage struct {
	basePath string
}

var ErrNoSavedPages = errors.New("no saved pages")

func NewStorage(basePath string) Storage {
	return Storage{
		basePath: basePath,
	}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = customError.Wrap("cannot save page", err) }()

	filePath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}

	fileName, err := fileName(page)

	if err != nil {
		return err
	}

	filePath = filepath.Join(filePath, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil

}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = customError.Wrap("cannot pick random page", err) }()

	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, ErrNoSavedPages
	}

	rand.New(rand.NewSource(time.Now().Unix()))

	file := files[rand.Intn(len(files))]

	return s.decodePage(filepath.Join(path, file.Name()))
}

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)

	if err != nil {
		return err
	}

	filePath := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(filePath); err != nil {
		return customError.Wrap(fmt.Sprintf("cannot remove file %s", filePath), err)
	}

	return nil
}

func (s Storage) IsExist(p *storage.Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, customError.Wrap("cannot check file", err)
	}
	filePath := filepath.Join(s.basePath, p.UserName, fileName)

	switch _, err = os.Stat(filePath); {
	case errors.Is(err, os.ErrNotExist):
		return false, err
	case err != nil:
		msg := fmt.Sprintf("cannot check file %s", filePath)

		return false, customError.Wrap(msg, err)
	}

	return false, err
}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, customError.Wrap("cannot decode page", err)
	}
	defer func() { _ := f.Close() }()

	var p storage.Page
	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, customError.Wrap("cannot decode page", err)
	}
	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
