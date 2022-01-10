// DBとのやりとりを定義
package repository

import (
	"backend/domain/model"
)

// UserRepository user repositoryのinterface
type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByID(id int) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) error
}
