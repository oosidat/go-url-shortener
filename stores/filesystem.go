package stores

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type Filesystem struct {
	Root string
	sync.RWMutex
}

func (s *Filesystem) Init(root string) error {
	s.Root = root
	return os.MkdirAll(s.Root, 0744)
}

func (s *Filesystem) Code() (string, error) {
	s.Lock()
	files, err := ioutil.ReadDir(s.Root)
	s.Unlock()

	if err != nil {
		return "", err
	}

	return strconv.FormatUint(uint64(len(files)+1), 36), err
}

func (s *Filesystem) Save(url string) (string, error) {
	code, codeGenerationErr := s.Code()
	if codeGenerationErr != nil {
		return "", codeGenerationErr
	}

	s.Lock()
	err := ioutil.WriteFile(filepath.Join(s.Root, code), []byte(url), 0744)
	if err != nil {
		return "", err
	}
	s.Unlock()

	return code, err
}

func (s *Filesystem) Load(code string) (string, error) {
	s.Lock()
	urlBytes, err := ioutil.ReadFile(filepath.Join(s.Root, code))
	s.Unlock()

	return string(urlBytes), err
}
