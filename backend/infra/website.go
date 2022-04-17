// repositoryに記述してあったinterfaceに合わせて、中身のメソッドを実装
package infra

import (
	"github.com/unemekenta/tiger_app/backend/domain/repository"

	"github.com/unemekenta/tiger_app/backend/domain/model"

	"gorm.io/gorm"
)

// WebsiteRepository website repositoryの構造体
type WebsiteRepository struct {
	Conn *gorm.DB
}

// NewWebsiteRepository website repositoryのコンストラクタ
func NewWebsiteRepository(conn *gorm.DB) repository.WebsiteRepository {
	return &WebsiteRepository{Conn: conn}
}

// Create websiteの保存
func (wr *WebsiteRepository) Create(website *model.Website) (*model.Website, error) {
	if err := wr.Conn.Create(&website).Error; err != nil {
		return nil, err
	}

	return website, nil
}

// FindByID websiteをIDで取得
func (wr *WebsiteRepository) FindByID(id int) (*model.Website, error) {
	website := &model.Website{ID: id}

	if err := wr.Conn.First(&website).Error; err != nil {
		return nil, err
	}

	return website, nil
}

// FindAll websiteを全て取得
func (wr *WebsiteRepository) FindAll() (*[]model.Website, error) {
	websites := &[]model.Website{}

	if err := wr.Conn.Order("name").Find(&websites).Error; err != nil {
		return nil, err
	}

	return websites, nil
}

// FindByCategory カテゴリからwebsiteを取得
func (wr *WebsiteRepository) FindByCategory(id int) (*[]model.Website, error) {
	websites := &[]model.Website{}

	if err := wr.Conn.Order("websites.name").Joins("left join categories_websites on categories_websites.website_id = websites.id").Where("categories_websites.category_id = ?", id).Find(&websites).Error; err != nil {
		return nil, err
	}

	return websites, nil
}

// SearchByName 検索結果からwebsiteを取得
func (wr *WebsiteRepository) SearchByName(query string) (*[]model.Website, error) {
	websites := &[]model.Website{}

	if err := wr.Conn.Order("name").Where("name ILIKE ?", "%"+query+"%").Find(&websites).Error; err != nil {
		return nil, err
	}

	return websites, nil
}

// Update websiteの更新
func (wr *WebsiteRepository) Update(website *model.Website) (*model.Website, error) {
	if err := wr.Conn.Model(&website).Updates(&website).Error; err != nil {
		return nil, err
	}

	return website, nil
}

// Delete websiteの削除
func (wr *WebsiteRepository) Delete(website *model.Website) error {
	if err := wr.Conn.Delete(&website).Error; err != nil {
		return err
	}

	return nil
}
