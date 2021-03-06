// ユースケースに沿った処理の流れを実装。
// データの取得や保存などでDBにアクセスするときもdomain層のrepositoryを介してアクセスすることによって、
// infrastructure層ではなくdomain層のみに依存させる。
package usecase

import (
	"github.com/unemekenta/tiger_app/backend/domain/model"
	"github.com/unemekenta/tiger_app/backend/domain/repository"

	"time"
)

// MoneyAccountUsecase moneyAccount usecaseのinterface
type MoneyAccountUsecase interface {
	Create(userID int, moneyAccountLabelID int, amount int, title string, contents string, year int, month int, subscriptionsFlg bool, updatedAt time.Time) (*model.MoneyAccount, error)
	CreateSubscription(startYear int, startMonth int, endYear int, endMonth int, userID int, moneyAccountLabelID int, amount int, title string, contents string, year int, month int, subscriptionsFlg bool, updatedAt time.Time) (*model.SubscriptionWithMoneyAccount, error)
	FindByID(id int) (*model.MoneyAccount, error)
	FindByUser(id int, year int, month int) (*[]model.MoneyAccount, error)
	FindSubscriptionWithMoneyAccountByUser(id int) (*[]model.SubscriptionWithMoneyAccount, error)
	Update(id int, moneyAccountLabelID int, amount int, title string, contents string, year int, month int, updatedAt time.Time) (*model.MoneyAccount, error)
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
func (mu *moneyAccountUsecase) Create(userID int, moneyAccountLabelID int, amount int, title string, contents string, year int, month int, subscriptionsFlg bool, updatedAt time.Time) (*model.MoneyAccount, error) {
	moneyAccount, err := model.NewMoneyAccount(userID, moneyAccountLabelID, amount, title, contents, year, month, subscriptionsFlg, updatedAt)
	if err != nil {
		return nil, err
	}

	createdMoneyAccount, err := mu.moneyAccountRepo.Create(moneyAccount)
	if err != nil {
		return nil, err
	}

	return createdMoneyAccount, nil
}

// CreateSubscription Subscriptionを保存するときのユースケース
func (mu *moneyAccountUsecase) CreateSubscription(startYear int, startMonth int, endYear int, endMonth int, userID int, moneyAccountLabelID int, amount int, title string, contents string, year int, month int, subscriptionsFlg bool, updatedAt time.Time) (*model.SubscriptionWithMoneyAccount, error) {
	subscriptionWithMoneyAccount, err := model.NewSubscriptionWithMoneyAccount(startYear, startMonth, endYear, endMonth, userID, moneyAccountLabelID, amount, title, contents, year, month, subscriptionsFlg, updatedAt)
	if err != nil {
		return nil, err
	}

	createdSubscriptionWithMoneyAccount, err := mu.moneyAccountRepo.CreateSubscription(subscriptionWithMoneyAccount)
	if err != nil {
		return nil, err
	}

	return createdSubscriptionWithMoneyAccount, nil
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
func (mu *moneyAccountUsecase) FindByUser(id int, year int, month int) (*[]model.MoneyAccount, error) {
	foundMoneyAccounts, err := mu.moneyAccountRepo.FindByUser(id, year, month)
	if err != nil {
		return nil, err
	}

	return foundMoneyAccounts, nil
}

// FindSubscriptionWithMoneyAccountByUser moneyAccountのうちSubscriptionをUserで取得するときのユースケース
func (mu *moneyAccountUsecase) FindSubscriptionWithMoneyAccountByUser(id int) (*[]model.SubscriptionWithMoneyAccount, error) {
	foundSubscriptionWithMoneyAccount, err := mu.moneyAccountRepo.FindSubscriptionWithMoneyAccountByUser(id)
	if err != nil {
		return nil, err
	}

	return foundSubscriptionWithMoneyAccount, nil
}

// Update moneyAccountを更新するときのユースケース
func (mu *moneyAccountUsecase) Update(id int, moneyAccountLabelID int, amount int, title string, contents string, year int, month int, updatedAt time.Time) (*model.MoneyAccount, error) {
	targetMoneyAccount, err := mu.moneyAccountRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetMoneyAccount.Set(moneyAccountLabelID, amount, title, contents, year, month, updatedAt)
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
