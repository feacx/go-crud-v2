package user

import "github.com/feacx/study/pkg/database"

func Get(id string) (userModel User) {
	database.DB.First(&userModel, id)
	return
}
