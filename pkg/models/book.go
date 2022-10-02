package models

type Book struct {
	Name   string `json:"name" binding:"required"`
	ISBN   string `json:"isbn" binding:"required"`
	Author string `json:"author" binding:"required"`
	ID     string `json:"id" bson:"id" binding:"required"`
}

type UpdateBookReq struct {
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type Books struct {
	Books []Book `json:"books" binding:"required"`
}
