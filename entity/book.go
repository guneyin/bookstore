package entity

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Book struct {
	gorm.Model      `json:"-"`
	Title           string `gorm:"uniqueIndex"`
	Author          string `gorm:"index"`
	PublicationYear uint
	Genre           string `gorm:"index"`
	Description     string
	CoverImage      string
}

type BookList []Book

type BookSearchParams struct {
	Title  WildcardString
	Author WildcardString
	Genre  WildcardString
}

type WildcardString string

func (ws WildcardString) ToWildcard() string {
	s := strings.TrimSpace(string(ws))
	s = strings.ToLower(s)

	return fmt.Sprintf("%%%s%%", s)
}
