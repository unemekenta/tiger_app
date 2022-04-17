// ユースケースに沿った処理の流れを実装。
// データの取得や保存などでDBにアクセスするときもdomain層のrepositoryを介してアクセスすることによって、
// infrastructure層ではなくdomain層のみに依存させる。
package usecase

import (
	"github.com/unemekenta/tiger_app/backend/domain/model"
	"github.com/unemekenta/tiger_app/backend/domain/repository"

	"time"
)

// TermUsecase term usecaseのinterface
type TermUsecase interface {
	Create(name string, content string, updatedAt time.Time) (*model.Term, error)
	FindByID(id int) (*model.Term, error)
	FindAll() (*[]model.Term, error)
	Update(id int, name string, content string, updatedAt time.Time) (*model.Term, error)
	Delete(id int) error
}

type termUsecase struct {
	termRepo repository.TermRepository
}

// NewTermUsecase term usecaseのコンストラクタ
func NewTermUsecase(termRepo repository.TermRepository) TermUsecase {
	return &termUsecase{termRepo: termRepo}
}

// Create termを保存するときのユースケース
func (tu *termUsecase) Create(name string, content string, updatedAt time.Time) (*model.Term, error) {
	term, err := model.NewTerm(name, content, updatedAt)
	if err != nil {
		return nil, err
	}

	createdTerm, err := tu.termRepo.Create(term)
	if err != nil {
		return nil, err
	}

	return createdTerm, nil
}

// FindByID termをIDで取得するときのユースケース
func (tu *termUsecase) FindByID(id int) (*model.Term, error) {
	foundTerm, err := tu.termRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundTerm, nil
}

// FindAll termを全て取得するときのユースケース
func (tu *termUsecase) FindAll() (*[]model.Term, error) {
	foundTerms, err := tu.termRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return foundTerms, nil
}

// Update termを更新するときのユースケース
func (tu *termUsecase) Update(id int, name string, content string, updatedAt time.Time) (*model.Term, error) {
	targetTerm, err := tu.termRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetTerm.Set(name, content, updatedAt)
	if err != nil {
		return nil, err
	}

	updatedTerm, err := tu.termRepo.Update(targetTerm)
	if err != nil {
		return nil, err
	}

	return updatedTerm, nil
}

// Delete termを削除するときのユースケース
func (tu *termUsecase) Delete(id int) error {
	term, err := tu.termRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = tu.termRepo.Delete(term)
	if err != nil {
		return err
	}

	return nil
}
