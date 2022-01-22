// repositoryに記述してあったinterfaceに合わせて、中身のメソッドを実装
package infra

import (
	"backend/domain/model"
	"backend/domain/repository"

	"gorm.io/gorm"
)

// WebsiteContentRepository websiteContent repositoryの構造体
type WebsiteContentRepository struct {
	Conn *gorm.DB
}

// NewWebsiteContentRepository websiteContent repositoryのコンストラクタ
func NewWebsiteContentRepository(conn *gorm.DB) repository.WebsiteContentRepository {
	return &WebsiteContentRepository{Conn: conn}
}

// Create websiteContentの保存
func (wr *WebsiteContentRepository) Create(websiteContent *model.WebsiteContent) (*model.WebsiteContent, error) {
	if err := wr.Conn.Create(&websiteContent).Error; err != nil {
		return nil, err
	}

	return websiteContent, nil
}

// FindByID websiteContentをIDで取得
func (wr *WebsiteContentRepository) FindByID(id int) (*model.WebsiteContent, error) {
	websiteContent := &model.WebsiteContent{ID: id}

	if err := wr.Conn.First(&websiteContent).Error; err != nil {
		return nil, err
	}

	return websiteContent, nil
}

// FindByWebsiteID websiteContentをWebsiteIDで取得
func (wr *WebsiteContentRepository) FindByWebsiteID(websiteID int) (*model.WebsiteContent, error) {
	websiteContent := &model.WebsiteContent{}

	if err := wr.Conn.Where("website_id = ?", websiteID).First(&websiteContent).Error; err != nil {
		return nil, err
	}

	return websiteContent, nil
}

// Update websiteContentの更新
func (wr *WebsiteContentRepository) Update(websiteContent *model.WebsiteContent) (*model.WebsiteContent, error) {
	if err := wr.Conn.Model(&websiteContent).Updates(&websiteContent).Error; err != nil {
		return nil, err
	}

	return websiteContent, nil
}

// Delete websiteContentの削除
func (wr *WebsiteContentRepository) Delete(websiteContent *model.WebsiteContent) error {
	if err := wr.Conn.Delete(&websiteContent).Error; err != nil {
		return err
	}

	return nil
}
