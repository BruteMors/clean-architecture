package book

type CreateBookDTO struct {
	Name       string `json:"name,omitempty"`
	Year       int    `json:"year,omitempty"`
	AuthorUUID string `json:"author_uuid"`
}

type UpdateBookDTO struct {
	UUID       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	Year       int    `json:"year,omitempty"`
	AuthorUUID string `json:"author_uuid,omitempty"`
	Busy       bool   `json:"busy,omitempty"`
	Owner      string `json:"owner,omitempty"`
}
