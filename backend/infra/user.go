// repositoryに記述してあったinterfaceに合わせて、中身のメソッドを実装
package infra

import (
	"backend/domain/model"
	"backend/domain/repository"

	"github.com/jinzhu/gorm"
)

// UserRepository user repositoryの構造体
type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository user repositoryのコンストラクタ
func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

// Create userの保存
func (ur *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := ur.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindByID userをIDで取得
func (ur *UserRepository) FindByID(id int) (*model.User, error) {
	user := &model.User{ID: id}

	if err := ur.Conn.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Update userの更新
func (ur *UserRepository) Update(user *model.User) (*model.User, error) {
	if err := ur.Conn.Model(&user).Update(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Delete userの削除
func (ur *UserRepository) Delete(user *model.User) error {
	if err := ur.Conn.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
