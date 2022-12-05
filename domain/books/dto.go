package books

// Bookを新規作成するためのリクエストのパラメータを保存するDTO
type CreateBookDTO struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

// Bookを1件取得するためのリクエストのパラメータを保存するDTO
type GetBookDTO struct {
	ID uint `param:"id"`
}
