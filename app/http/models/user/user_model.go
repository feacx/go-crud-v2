package user

import "github.com/feacx/study/app/http/models"

type User struct {
	models.BaseModel
	Name string `json:"name,omitempty"`
}
