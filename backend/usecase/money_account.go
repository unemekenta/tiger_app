// ユースケースに沿った処理の流れを実装。
// データの取得や保存などでDBにアクセスするときもdomain層のrepositoryを介してアクセスすることによって、
// infrastructure層ではなくdomain層のみに依存させる。
package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"time"
)

// MoneyAccountUsecase moneyAccount usecaseのinterface
type MoneyAccountUsecase interface {
	Create(userID int, moneyAccountLabelID int, amount int, title string, contents string, updatedAt time.Time) (*model.MoneyAccount, error)
	FindByID(id int) (*model.MoneyAccount, error)
	FindByUser(id int) (*[]model.MoneyAccount, error)
	Update(id int, moneyAccountLabelID int, amount int, title string, contents string, updatedAt time.Time) (*model.MoneyAccount, error)
	Delete(id int) error
}

type moneyAccountUsecase struct {
	moneyAccountRepo repository.MoneyAccountRepository
}

// NewMoneyAccountUsecase moneyAccount usecaseのコンストラクタ
func NewMoneyAccountUsecase(moneyAccountRepo repository.MoneyAccountRepository) MoneyAccountUsecase {
	return &moneyAccountUsecase{moneyAccountRepo: moneyAccountRepo}
}

// Create moneyAccountを保存するときのユースケース
func (mu *moneyAccountUsecase) Create(userID int, moneyAccountLabelID int, amount int, title string, contents string, updatedAt time.Time) (*model.MoneyAccount, error) {
	moneyAccount, err := model.NewMoneyAccount(userID, moneyAccountLabelID, amount, title, contents, updatedAt)
	if err != nil {
		return nil, err
	}

	createdMoneyAccount, err := mu.moneyAccountRepo.Create(moneyAccount)
	if err != nil {
		return nil, err
	}

	return createdMoneyAccount, nil
}

// FindByID moneyAccountをIDで取得するときのユースケース
func (mu *moneyAccountUsecase) FindByID(id int) (*model.MoneyAccount, error) {
	foundMoneyAccount, err := mu.moneyAccountRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundMoneyAccount, nil
}

// FindByUser moneyAccountをUserで取得するときのユースケース
func (mu *moneyAccountUsecase) FindByUser(id int) (*[]model.MoneyAccount, error) {
	foundMoneyAccounts, err := mu.moneyAccountRepo.FindByUser(id)
	if err != nil {
		return nil, err
	}

	return foundMoneyAccounts, nil
}

// Update moneyAccountを更新するときのユースケース
func (mu *moneyAccountUsecase) Update(id int, moneyAccountLabelID int, amount int, title string, contents string, updatedAt time.Time) (*model.MoneyAccount, error) {
	targetMoneyAccount, err := mu.moneyAccountRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetMoneyAccount.Set(moneyAccountLabelID, amount, title, contents, updatedAt)
	if err != nil {
		return nil, err
	}

	updatedMoneyAccount, err := mu.moneyAccountRepo.Update(targetMoneyAccount)
	if err != nil {
		return nil, err
	}

	return updatedMoneyAccount, nil
}

// Delete moneyAccountを削除するときのユースケース
func (mu *moneyAccountUsecase) Delete(id int) error {
	moneyAccount, err := mu.moneyAccountRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = mu.moneyAccountRepo.Delete(moneyAccount)
	if err != nil {
		return err
	}

	return nil
}
