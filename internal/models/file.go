package models

type File struct {
	Filename string `bson:"filename" json:"filename"`
	Id       string `bson:"_id" json:"_id"`
}
