package repository

import (
	"fmt"

	"github.com/Kourin1996/go-api-e2e-sample/domain/books"
)

// Repository層の実装
// 今回は簡易的にDBに繋げずインメモリで実装
type BooksRepository struct {
	books []*books.Book
}

var _ books.IBooksRepository = (*BooksRepository)(nil)

func NewRepository() *BooksRepository {
	return &BooksRepository{
		books: make([]*books.Book, 0),
	}
}

func (r *BooksRepository) CreateBook(dto *books.CreateBookDTO) (*books.Book, error) {
	// 新しいレコードのIDを求める
	// (今回はインメモリに保存するため、IDは現配列長から求めているが、実際はDBなどが採番する)
	id := uint(len(r.books))

	// レコードをDTOから作成
	book, err := books.NewBook(id, dto)
	if err != nil {
		return nil, err
	}

	// インメモリに保存する
	r.books = append(r.books, book)

	fmt.Printf("Created a book %+v\n", book)

	// レコードを返す
	return book, nil
}

// IDに基づきBookレコードを返す
func (r *BooksRepository) GetBook(dto *books.GetBookDTO) (*books.Book, error) {
	// もしIDが未使用の場合、単にnullを返す
	if dto.ID >= uint(len(r.books)) {
		return nil, nil
	}

	fmt.Printf("Return a book %+v\n", r.books[dto.ID])

	// IDに対応するBookを返す
	return r.books[dto.ID], nil
}
