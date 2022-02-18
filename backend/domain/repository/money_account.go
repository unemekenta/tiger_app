// DBとのやりとりを定義
package repository

import (
	"backend/domain/model"
)

// MoneyAccountRepository moneyAccount repositoryのinterface
type MoneyAccountRepository interface {
	Create(moneyAccount *model.MoneyAccount) (*model.MoneyAccount, error)
	FindByID(id int) (*model.MoneyAccount, error)
	FindByUser(id int, year int, month int) (*[]model.MoneyAccount, error)
	Update(moneyAccount *model.MoneyAccount) (*model.MoneyAccount, error)
	Delete(moneyAccount *model.MoneyAccount) error
}
