// DBとのやりとりを定義
package repository

import (
	"backend/domain/model"
)

// WebsiteRepository website repositoryのinterface
type WebsiteRepository interface {
	Create(website *model.Website) (*model.Website, error)
	FindByID(id int) (*model.Website, error)
	FindAll() (*[]model.Website, error)
	FindByCategory(id int) (*[]model.Website, error)
	SearchByName(query string) (*[]model.Website, error)
	Update(website *model.Website) (*model.Website, error)
	Delete(website *model.Website) error
}
