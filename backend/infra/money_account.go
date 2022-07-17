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

// CreateSubscription subscriptionの保存
func (mr *MoneyAccountRepository) CreateSubscription(subscriptionWithMoneyAccount *model.SubscriptionWithMoneyAccount) (*model.SubscriptionWithMoneyAccount, error) {

	sm := &model.MoneyAccount{
		UserID:              subscriptionWithMoneyAccount.MoneyAccount.UserID,
		MoneyAccountLabelID: subscriptionWithMoneyAccount.MoneyAccount.MoneyAccountLabelID,
		Amount:              subscriptionWithMoneyAccount.MoneyAccount.Amount,
		Title:               subscriptionWithMoneyAccount.MoneyAccount.Title,
		Contents:            subscriptionWithMoneyAccount.MoneyAccount.Contents,
		Year:                subscriptionWithMoneyAccount.MoneyAccount.Year,
		Month:               subscriptionWithMoneyAccount.MoneyAccount.Month,
		SubscriptionsFlg:    subscriptionWithMoneyAccount.MoneyAccount.SubscriptionsFlg,
	}

	if err := mr.Conn.Create(&sm).Error; err != nil {
		return nil, err
	}

	moneyAccountID := &sm.ID

	ss := &model.Subscription{
		MoneyAccountID: *moneyAccountID,
		StartYear:      subscriptionWithMoneyAccount.Subscription.StartYear,
		StartMonth:     subscriptionWithMoneyAccount.Subscription.StartMonth,
		EndYear:        subscriptionWithMoneyAccount.Subscription.EndYear,
		EndMonth:       subscriptionWithMoneyAccount.Subscription.EndMonth,
	}

	if err := mr.Conn.Debug().Create(&ss).Error; err != nil {
		return nil, err
	}

	swm := &model.SubscriptionWithMoneyAccount{
		MoneyAccount: *sm,
		Subscription: *ss,
	}

	return swm, nil
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
			Or("subscriptions_flg = ? AND((((start_year <= ?) OR (start_year = ? AND start_month <= ?)) OR (start_year IN (0, NULL) AND start_month IN (0, NULL)) ) AND ( (? <= end_year AND ? <= end_month) OR (end_year IN (0, NULL) AND end_month IN (0, NULL))))", true, year, year, month, year, month),
		).
		Find(&moneyAccounts).Error; err != nil {
		return nil, err
	}

	return moneyAccounts, nil
}

// FindSubscriptionsByUser moneyAccountのうちSubscriptionsをUserで取得
func (mr *MoneyAccountRepository) FindSubscriptionWithMoneyAccountByUser(id int) (*[]model.SubscriptionWithMoneyAccount, error) {
	subscriptions := &[]model.SubscriptionWithMoneyAccount{}

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
