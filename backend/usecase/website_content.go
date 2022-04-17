// ユースケースに沿った処理の流れを実装。
// データの取得や保存などでDBにアクセスするときもdomain層のrepositoryを介してアクセスすることによって、
// infrastructure層ではなくdomain層のみに依存させる。
package usecase

import (
	"github.com/unemekenta/tiger_app/backend/domain/model"
	"github.com/unemekenta/tiger_app/backend/domain/repository"

	"time"
)

// WebsiteContentUsecase websiteContent usecaseのinterface
type WebsiteContentUsecase interface {
	Create(websiteId int, title string, contents string, updatedAt time.Time) (*model.WebsiteContent, error)
	FindByID(id int) (*model.WebsiteContent, error)
	FindByWebsiteID(websiteID int) (*model.WebsiteContent, error)
	Update(id int, title string, contents string, updatedAt time.Time) (*model.WebsiteContent, error)
	Delete(id int) error
}

type websiteContentUsecase struct {
	websiteContentRepo repository.WebsiteContentRepository
}

// NewWebsiteContentUsecase websiteContent usecaseのコンストラクタ
func NewWebsiteContentUsecase(websiteContentRepo repository.WebsiteContentRepository) WebsiteContentUsecase {
	return &websiteContentUsecase{websiteContentRepo: websiteContentRepo}
}

// Create websiteContentを保存するときのユースケース
func (wcu *websiteContentUsecase) Create(websiteId int, title string, contents string, updatedAt time.Time) (*model.WebsiteContent, error) {
	websiteContent, err := model.NewWebsiteContent(websiteId, title, contents, updatedAt)
	if err != nil {
		return nil, err
	}

	createdWebsiteContent, err := wcu.websiteContentRepo.Create(websiteContent)
	if err != nil {
		return nil, err
	}

	return createdWebsiteContent, nil
}

// FindByID websiteContentをIDで取得するときのユースケース
func (wcu *websiteContentUsecase) FindByID(id int) (*model.WebsiteContent, error) {
	foundWebsiteContent, err := wcu.websiteContentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundWebsiteContent, nil
}

// FindByWebsiteID websiteContentをWebsiteIDで取得するときのユースケース
func (wcu *websiteContentUsecase) FindByWebsiteID(websiteId int) (*model.WebsiteContent, error) {
	foundWebsiteContent, err := wcu.websiteContentRepo.FindByWebsiteID(websiteId)
	if err != nil {
		return nil, err
	}

	return foundWebsiteContent, nil
}

// Update websiteContentを更新するときのユースケース
func (wcu *websiteContentUsecase) Update(id int, title string, contents string, updatedAt time.Time) (*model.WebsiteContent, error) {
	targetWebsiteContent, err := wcu.websiteContentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetWebsiteContent.Set(title, contents, updatedAt)
	if err != nil {
		return nil, err
	}

	updatedWebsiteContent, err := wcu.websiteContentRepo.Update(targetWebsiteContent)
	if err != nil {
		return nil, err
	}

	return updatedWebsiteContent, nil
}

// Delete websiteContentを削除するときのユースケース
func (wcu *websiteContentUsecase) Delete(id int) error {
	websiteContent, err := wcu.websiteContentRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = wcu.websiteContentRepo.Delete(websiteContent)
	if err != nil {
		return err
	}

	return nil
}
