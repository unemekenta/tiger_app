// ユースケースに沿った処理の流れを実装。
// データの取得や保存などでDBにアクセスするときもdomain層のrepositoryを介してアクセスすることによって、
// infrastructure層ではなくdomain層のみに依存させる。
package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"time"
)

// CategoryUsecase category usecaseのinterface
type CategoryUsecase interface {
	Create(ancestorId int, name string, updatedAt time.Time) (*model.Category, error)
	FindByID(id int) (*model.Category, error)
	FindAllParentCategories() (*[]model.Category, error)
	FindCategoriesByAncestor(id int) (*[]model.Category, error)
	Update(id int, ancestorId int, name string, updatedAt time.Time) (*model.Category, error)
	Delete(id int) error
}

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

// NewCategoryUsecase category usecaseのコンストラクタ
func NewCategoryUsecase(categoryRepo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepo: categoryRepo}
}

// Create categoryを保存するときのユースケース
func (cu *categoryUsecase) Create(ancestorId int, name string, updatedAt time.Time) (*model.Category, error) {
	category, err := model.NewCategory(ancestorId, name, updatedAt)
	if err != nil {
		return nil, err
	}

	createdCategory, err := cu.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}

// FindByID categoryをIDで取得するときのユースケース
func (cu *categoryUsecase) FindByID(id int) (*model.Category, error) {
	foundCategory, err := cu.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundCategory, nil
}

// FindAllParentCategories 親categoryを全て取得するときのユースケース
func (cu *categoryUsecase) FindAllParentCategories() (*[]model.Category, error) {
	foundCategories, err := cu.categoryRepo.FindAllParentCategories()
	if err != nil {
		return nil, err
	}

	return foundCategories, nil
}

// FindCategoriesByAncestor 親カテゴリから子カテゴリを取得するときのユースケース
func (cu *categoryUsecase) FindCategoriesByAncestor(id int) (*[]model.Category, error) {
	foundCategories, err := cu.categoryRepo.FindCategoriesByAncestor(id)
	if err != nil {
		return nil, err
	}

	return foundCategories, nil
}

// Update categoryを更新するときのユースケース
func (cu *categoryUsecase) Update(id int, ancestorId int, name string, updatedAt time.Time) (*model.Category, error) {
	targetCategory, err := cu.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetCategory.Set(ancestorId, name, updatedAt)
	if err != nil {
		return nil, err
	}

	updatedCategory, err := cu.categoryRepo.Update(targetCategory)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

// Delete categoryを削除するときのユースケース
func (cu *categoryUsecase) Delete(id int) error {
	category, err := cu.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = cu.categoryRepo.Delete(category)
	if err != nil {
		return err
	}

	return nil
}
