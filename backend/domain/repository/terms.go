// DBとのやりとりを定義
package repository

import (
	"backend/domain/model"
)

// TermRepository term repositoryのinterface
type TermRepository interface {
	Create(term *model.Term) (*model.Term, error)
	FindByID(id int) (*model.Term, error)
	Update(term *model.Term) (*model.Term, error)
	Delete(term *model.Term) error
}
