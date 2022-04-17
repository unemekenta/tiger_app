// DBとのやりとりを定義
package repository

import (
	"github.com/unemekenta/tiger_app/backend/domain/model"
)

// UserRepository user repositoryのinterface
type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByID(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) error
	CreateSession(uuid string, value string) error
	LoginCheck(uuid string) (string, error)
}
