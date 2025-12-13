package users

import "gorm.io/gorm"

type UserModule struct {
	Repo UserRepository
}

func InitUserModule(db *gorm.DB) *UserModule {
	repo := NewUserRepository(db)

	return &UserModule{
		Repo: repo,
	}

}
