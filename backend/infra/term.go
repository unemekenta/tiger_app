// repositoryに記述してあったinterfaceに合わせて、中身のメソッドを実装
package infra

import (
	"backend/domain/model"
	"backend/domain/repository"

	"gorm.io/gorm"
)

// TermRepository term repositoryの構造体
type TermRepository struct {
	Conn *gorm.DB
}

// NewTermRepository term repositoryのコンストラクタ
func NewTermRepository(conn *gorm.DB) repository.TermRepository {
	return &TermRepository{Conn: conn}
}

// Create termの保存
func (tr *TermRepository) Create(term *model.Term) (*model.Term, error) {
	if err := tr.Conn.Create(&term).Error; err != nil {
		return nil, err
	}

	return term, nil
}

// FindByID termをIDで取得
func (tr *TermRepository) FindByID(id int) (*model.Term, error) {
	term := &model.Term{ID: id}

	if err := tr.Conn.First(&term).Error; err != nil {
		return nil, err
	}

	return term, nil
}

// Update termの更新
func (tr *TermRepository) Update(term *model.Term) (*model.Term, error) {
	if err := tr.Conn.Model(&term).Updates(&term).Error; err != nil {
		return nil, err
	}

	return term, nil
}

// Delete termの削除
func (tr *TermRepository) Delete(term *model.Term) error {
	if err := tr.Conn.Delete(&term).Error; err != nil {
		return err
	}

	return nil
}
