// DBとのやりとりを定義
package repository

import (
	"backend/domain/model"
)

// WebsiteContentRepository websiteContent repositoryのinterface
type WebsiteContentRepository interface {
	Create(websiteContent *model.WebsiteContent) (*model.WebsiteContent, error)
	FindByID(id int) (*model.WebsiteContent, error)
	FindByWebsiteID(websiteID int) (*model.WebsiteContent, error)
	Update(websiteContent *model.WebsiteContent) (*model.WebsiteContent, error)
	Delete(websiteContent *model.WebsiteContent) error
}
