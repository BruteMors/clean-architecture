package author

type CreateAuthorDTO struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type UpdateBookDTO struct {
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
