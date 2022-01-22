// ユースケースに沿った処理の流れを実装。
// データの取得や保存などでDBにアクセスするときもdomain層のrepositoryを介してアクセスすることによって、
// infrastructure層ではなくdomain層のみに依存させる。
package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"time"
)

// WebsiteUsecase website usecaseのinterface
type WebsiteUsecase interface {
	Create(name string, url string, companyName string, image string, updatedAt time.Time) (*model.Website, error)
	FindByID(id int) (*model.Website, error)
	FindAll() (*[]model.Website, error)
	FindByCategory(id int) (*[]model.Website, error)
	SearchByName(query string) (*[]model.Website, error)
	Update(id int, name string, url string, companyName string, image string, updatedAt time.Time) (*model.Website, error)
	Delete(id int) error
}

type websiteUsecase struct {
	websiteRepo repository.WebsiteRepository
}

// NewWebsiteUsecase website usecaseのコンストラクタ
func NewWebsiteUsecase(websiteRepo repository.WebsiteRepository) WebsiteUsecase {
	return &websiteUsecase{websiteRepo: websiteRepo}
}

// Create websiteを保存するときのユースケース
func (wu *websiteUsecase) Create(name string, url string, companyName string, image string, updatedAt time.Time) (*model.Website, error) {
	website, err := model.NewWebsite(name, url, companyName, image, updatedAt)
	if err != nil {
		return nil, err
	}

	createdWebsite, err := wu.websiteRepo.Create(website)
	if err != nil {
		return nil, err
	}

	return createdWebsite, nil
}

// FindByID websiteをIDで取得するときのユースケース
func (wu *websiteUsecase) FindByID(id int) (*model.Website, error) {
	foundWebsite, err := wu.websiteRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundWebsite, nil
}

// FindAll websiteを全て取得するときのユースケース
func (wu *websiteUsecase) FindAll() (*[]model.Website, error) {
	foundWebsites, err := wu.websiteRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return foundWebsites, nil
}

// FindByCategory websiteをカテゴリから取得するときのユースケース
func (wu *websiteUsecase) FindByCategory(id int) (*[]model.Website, error) {
	foundWebsites, err := wu.websiteRepo.FindByCategory(id)
	if err != nil {
		return nil, err
	}

	return foundWebsites, nil
}

// SearchByName 検索結果を取得するときのユースケース
func (wu *websiteUsecase) SearchByName(query string) (*[]model.Website, error) {
	foundWebsites, err := wu.websiteRepo.SearchByName(query)
	if err != nil {
		return nil, err
	}

	return foundWebsites, nil
}

// Update websiteを更新するときのユースケース
func (wu *websiteUsecase) Update(id int, name string, url string, companyName string, image string, updatedAt time.Time) (*model.Website, error) {
	targetWebsite, err := wu.websiteRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetWebsite.Set(name, url, companyName, image, updatedAt)
	if err != nil {
		return nil, err
	}

	updatedWebsite, err := wu.websiteRepo.Update(targetWebsite)
	if err != nil {
		return nil, err
	}

	return updatedWebsite, nil
}

// Delete websiteを削除するときのユースケース
func (wu *websiteUsecase) Delete(id int) error {
	website, err := wu.websiteRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = wu.websiteRepo.Delete(website)
	if err != nil {
		return err
	}

	return nil
}
