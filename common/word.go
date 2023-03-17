package common

import (
	"fmt"
	"path/filepath"

	//   "/Users/pengdongjiang/Documents/code/sale/go-service/dict.txt"
	"github.com/importcjj/sensitive"
)

var (
	Default *Word
)

type Word struct {
	filter *sensitive.Filter
}

func New(path string) (*Word, error) {
	s := sensitive.New()
	files, _ := filepath.Glob(fmt.Sprintf("%s/*.txt", path))
	for _, file := range files {
		if err := s.LoadWordDict(file); err != nil {
			return nil, err
		}
	}
	s.UpdateNoisePattern("的")
	return &Word{
		filter: s,
	}, nil
}

func (w *Word) Validate(text string) (bool, string) {
	return w.filter.Validate(text)
}

//移除
func (w *Word) Filter(text string) string {
	return w.filter.Filter(text)
}

//添加
func (w *Word) AddWord(text string) {
	w.filter.AddWord(text)
}

//替换敏感词
func (w *Word) Replace(text string) string {
	return w.filter.Replace(text, '*')
}

//查找全部敏感词
func (w *Word) FindAll(text string) []string {
	return w.filter.FindAll(text)
}
