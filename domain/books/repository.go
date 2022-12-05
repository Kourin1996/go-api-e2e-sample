package books

type IBooksRepository interface {
	CreateBook(*CreateBookDTO) (*Book, error)
	GetBook(*GetBookDTO) (*Book, error)
}
