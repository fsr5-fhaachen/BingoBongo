package model

import "fmt"

type Task struct {
	Description string   `json:"task"`
	Type        FileType `json:"type"`
	ID          int
}

type Tasks []Task

func (t Tasks) TypeByID(id int) (FileType, error) {
	if id >= len(t) {
		return FileTypeUnknown, fmt.Errorf("invalid taskID")
	}
	return t[id].Type, nil
}
