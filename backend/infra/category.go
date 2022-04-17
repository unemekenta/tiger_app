// repositoryに記述してあったinterfaceに合わせて、中身のメソッドを実装
package infra

import (
	"github.com/unemekenta/tiger_app/backend/domain/model"
	"github.com/unemekenta/tiger_app/backend/domain/repository"

	"gorm.io/gorm"
)

// CategoryRepository category repositoryの構造体
type CategoryRepository struct {
	Conn *gorm.DB
}

// NewCategoryRepository category repositoryのコンストラクタ
func NewCategoryRepository(conn *gorm.DB) repository.CategoryRepository {
	return &CategoryRepository{Conn: conn}
}

// Create categoryの保存
func (cr *CategoryRepository) Create(category *model.Category) (*model.Category, error) {
	if err := cr.Conn.Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// FindByID categoryをIDで取得
func (cr *CategoryRepository) FindByID(id int) (*model.Category, error) {
	category := &model.Category{ID: id}

	if err := cr.Conn.First(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// FindAllParentCategories 末端categoryを全て取得
func (cr *CategoryRepository) FindAllParentCategories() (*[]model.Category, error) {
	categories := &[]model.Category{}

	if err := cr.Conn.Order("name").Where("ancestor_id IS NULL").Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// FindCategoriesByAncestor 親カテゴリから子カテゴリを全て取得
func (cr *CategoryRepository) FindCategoriesByAncestor(id int) (*[]model.Category, error) {
	categories := &[]model.Category{}

	if err := cr.Conn.Order("name").Where("ancestor_id = ?", id).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// FindCategoriesByWebsite websiteから子カテゴリを取得
func (cr *CategoryRepository) FindCategoryByWebsite(id int) (*model.Category, error) {
	category := &model.Category{}

	if err := cr.Conn.Order("name").Joins("left join categories_websites on categories_websites.category_id = categories.id").Where("categories_websites.website_id = ?", id).Where("website_id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// Update categoryの更新
func (cr *CategoryRepository) Update(category *model.Category) (*model.Category, error) {
	if err := cr.Conn.Model(&category).Updates(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// Delete categoryの削除
func (cr *CategoryRepository) Delete(category *model.Category) error {
	if err := cr.Conn.Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
