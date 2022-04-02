// repositoryに記述してあったinterfaceに合わせて、中身のメソッドを実装
package infra

import (
	"github.com/unemekenta/tiger_app/backend/domain/model"
	"github.com/unemekenta/tiger_app/backend/domain/repository"

	"gorm.io/gorm"
)

// MoneyAccountRepository moneyAccount repositoryの構造体
type MoneyAccountRepository struct {
	Conn *gorm.DB
}

// NewMoneyAccountRepository moneyAccount repositoryのコンストラクタ
func NewMoneyAccountRepository(conn *gorm.DB) repository.MoneyAccountRepository {
	return &MoneyAccountRepository{Conn: conn}
}

// Create moneyAccountの保存
func (mr *MoneyAccountRepository) Create(moneyAccount *model.MoneyAccount) (*model.MoneyAccount, error) {
	if err := mr.Conn.Create(&moneyAccount).Error; err != nil {
		return nil, err
	}

	return moneyAccount, nil
}

// FindByID moneyAccountをIDで取得
func (mr *MoneyAccountRepository) FindByID(id int) (*model.MoneyAccount, error) {
	moneyAccount := &model.MoneyAccount{ID: id}

	if err := mr.Conn.First(&moneyAccount).Error; err != nil {
		return nil, err
	}

	return moneyAccount, nil
}

// FindByUser moneyAccountをUserで取得
func (mr *MoneyAccountRepository) FindByUser(id int, year int, month int) (*[]model.MoneyAccount, error) {
	moneyAccounts := &[]model.MoneyAccount{}

	if err := mr.Conn.Debug().Order("id").
		Joins("left join subscriptions on subscriptions.money_account_id = money_accounts.id").
		Where("user_id = ?", id).
		Where(mr.Conn.
			Where("subscriptions_flg = ? AND year = ? AND month = ?", false, year, month).
			Or("subscriptions_flg = ? AND(((start_year <= ? AND start_month <= ?) OR (start_year IS NULL AND start_month IS NULL) ) AND ( (? <= end_year AND ? <= end_month) OR (end_year IS NULL AND end_month IS NULL)))", true, year, month, year, month),
		).
		Find(&moneyAccounts).Error; err != nil {
		return nil, err
	}

	return moneyAccounts, nil
}

// FindSubscriptionsByUser moneyAccountのうちSubscriptionsをUserで取得
func (mr *MoneyAccountRepository) FindSubscriptionsByUser(id int) (*[]model.Subscription, error) {
	subscriptions := &[]model.Subscription{}

	if err := mr.Conn.Debug().
		Table("subscriptions").
		Order("subscriptions.id").
		Select("subscriptions.*, money_accounts.*").
		Joins("left join money_accounts on subscriptions.money_account_id = money_accounts.id").
		Where("money_accounts.user_id = ?", id).
		Where("subscriptions_flg = ?", true).
		Find(&subscriptions).Error; err != nil {
		return nil, err
	}

	return subscriptions, nil
}

// Update moneyAccountの更新
func (mr *MoneyAccountRepository) Update(moneyAccount *model.MoneyAccount) (*model.MoneyAccount, error) {
	if err := mr.Conn.Model(&moneyAccount).Updates(&moneyAccount).Error; err != nil {
		return nil, err
	}

	return moneyAccount, nil
}

// Delete moneyAccountの削除
func (mr *MoneyAccountRepository) Delete(moneyAccount *model.MoneyAccount) error {
	if err := mr.Conn.Delete(&moneyAccount).Error; err != nil {
		return err
	}

	return nil
}
