// DBとのやりとりを定義
package repository

import (
	"github.com/unemekenta/tiger_app/backend/domain/model"
)

// MoneyAccountRepository moneyAccount repositoryのinterface
type MoneyAccountRepository interface {
	Create(moneyAccount *model.MoneyAccount) (*model.MoneyAccount, error)
	FindByID(id int) (*model.MoneyAccount, error)
	FindByUser(id int, year int, month int) (*[]model.MoneyAccount, error)
	FindSubscriptionsByUser(id int) (*[]model.MoneyAccount, error)
	Update(moneyAccount *model.MoneyAccount) (*model.MoneyAccount, error)
	Delete(moneyAccount *model.MoneyAccount) error
}
