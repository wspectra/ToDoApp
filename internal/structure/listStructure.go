package structure

type List struct {
	Id          int    `json:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description" `
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type AllListResponse struct {
	Status  string
	Message []List
}

type ListResponse struct {
	Status  string
	Message List
}
