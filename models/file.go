package models

type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int    `json:"size"`
}
