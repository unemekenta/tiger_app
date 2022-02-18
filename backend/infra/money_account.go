// repositoryに記述してあったinterfaceに合わせて、中身のメソッドを実装
package infra

import (
	"backend/domain/model"
	"backend/domain/repository"

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

// FindByUser moneyAccountををUserで取得
func (mr *MoneyAccountRepository) FindByUser(id int, year int, month int) (*[]model.MoneyAccount, error) {
	moneyAccounts := &[]model.MoneyAccount{}

	if err := mr.Conn.Order("id").Where("user_id = ?", id).Where("user_id = ?", id).Where("year = ?", year).Where("month = ?", month).Find(&moneyAccounts).Error; err != nil {
		return nil, err
	}

	return moneyAccounts, nil
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
