package models

import "mime/multipart"

type File struct {
	File multipart.File `json:"file,omitempty"`
}

type Url struct {
	Url string `json:"url,omitempty"`
}

type Image struct {
	AnimalId string `db:"animal_id" json:"animal_id"`
	URL      string `db:"url" json:"url"`
}
