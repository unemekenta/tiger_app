// DBとのやりとりを定義
package repository

import (
	"backend/domain/model"
)

// CategoryRepository category repositoryのinterface
type CategoryRepository interface {
	Create(category *model.Category) (*model.Category, error)
	FindByID(id int) (*model.Category, error)
	FindAllChildCategories() (*[]model.Category, error)
	FindCategoriesByAncestor(id int) (*[]model.Category, error)
	Update(category *model.Category) (*model.Category, error)
	Delete(category *model.Category) error
}
