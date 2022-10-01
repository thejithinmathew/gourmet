package models

type Book struct {
	Generation int    `json:"generation" bson:"generation" default:"1" binding:"required"`
	Name       string `json:"name" bson:"name" binding:"required"`
	Author     string `json:"author" bson:"author" binding:"required"`
	ISBN       string `json:"isbn" bson:"isbn" binding:"required"`
}

type BookReq struct {
	Name   string `json:"name" binding:"required"`
	ISBN   string `json:"isbn" binding:"required"`
	Author string `json:"author" binding:"required"`
}
