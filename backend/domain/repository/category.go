// DBとのやりとりを定義
package repository

import (
	"github.com/unemekenta/tiger_app/backend/domain/model"
)

// CategoryRepository category repositoryのinterface
type CategoryRepository interface {
	Create(category *model.Category) (*model.Category, error)
	FindByID(id int) (*model.Category, error)
	FindAllParentCategories() (*[]model.Category, error)
	FindCategoriesByAncestor(id int) (*[]model.Category, error)
	FindCategoryByWebsite(id int) (*model.Category, error)
	Update(category *model.Category) (*model.Category, error)
	Delete(category *model.Category) error
}
