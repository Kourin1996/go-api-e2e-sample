package books

import (
	"errors"
	"time"
)

var (
	ErrEmptyName   = errors.New("name must not be empty")
	ErrEmptyAuthor = errors.New("author must not be empty")
)

type Book struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"` // レコードが追加されたタイムスタンプ
}

func NewBook(
	id uint,
	dto *CreateBookDTO,
) (*Book, error) {
	if len(dto.Name) == 0 {
		return nil, ErrEmptyName
	}

	if len(dto.Author) == 0 {
		return nil, ErrEmptyAuthor
	}

	return &Book{
		ID:        id,
		Name:      dto.Name,
		Author:    dto.Author,
		CreatedAt: time.Now(),
	}, nil
}
