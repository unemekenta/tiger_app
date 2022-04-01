// repositoryに記述してあったinterfaceに合わせて、中身のメソッドを実装
package infra

import (
	"github.com/unemekenta/tiger_app/backend/domain/model"
	"github.com/unemekenta/tiger_app/backend/domain/repository"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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

// FindByEmail userをemailで取得
func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	if err := ur.Conn.Debug().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Update userの更新
func (ur *UserRepository) Update(user *model.User) (*model.User, error) {
	if err := ur.Conn.Model(&user).Updates(&user).Error; err != nil {
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

// CreateSession sessionの作成
func (ur *UserRepository) CreateSession(uuid string, value string) error {
	// コネクションプールの作成
	pool := model.RedisNewPool()

	// コネクションの取得
	conn := pool.Get()
	defer conn.Close()

	// TODO セッション時間変更
	_, err := conn.Do("SET", uuid, value, "NX", "EX", 36000)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (ur *UserRepository) LoginCheck(uuid string) (string, error) {
	// コネクションプールの作成
	pool := model.RedisNewPool()

	// コネクションの取得
	conn := pool.Get()
	defer conn.Close()

	su, err := redis.String(conn.Do("GET", uuid))
	if err != nil {
		return "", errors.WithStack(err)
	}

	return su, nil
}
