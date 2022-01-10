// ユースケースに沿った処理の流れを実装。
// データの取得や保存などでDBにアクセスするときもdomain層のrepositoryを介してアクセスすることによって、
// infrastructure層ではなくdomain層のみに依存させる。
package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"time"
)

// UserUsecase user usecaseのinterface
type UserUsecase interface {
	Create(email string, name string, password string, roleid int, updatedAt time.Time) (*model.User, error)
	FindByID(id int) (*model.User, error)
	Update(id int, email string, name string, password string, roleid int, updatedAt time.Time) (*model.User, error)
	Delete(id int) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase user usecaseのコンストラクタ
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

// Create userを保存するときのユースケース
func (uu *userUsecase) Create(email string, name string, password string, roleid int, updatedAt time.Time) (*model.User, error) {
	user, err := model.NewUser(email, name, password, roleid, updatedAt)
	if err != nil {
		return nil, err
	}

	createdUser, err := uu.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// FindByID userをIDで取得するときのユースケース
func (uu *userUsecase) FindByID(id int) (*model.User, error) {
	foundUser, err := uu.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

// Update userを更新するときのユースケース
func (uu *userUsecase) Update(id int, email string, name string, password string, roleid int, updatedAt time.Time) (*model.User, error) {
	targetUser, err := uu.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetUser.Set(email, name, password, roleid, updatedAt)
	if err != nil {
		return nil, err
	}

	updatedUser, err := uu.userRepo.Update(targetUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// Delete userを削除するときのユースケース
func (uu *userUsecase) Delete(id int) error {
	user, err := uu.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = uu.userRepo.Delete(user)
	if err != nil {
		return err
	}

	return nil
}
