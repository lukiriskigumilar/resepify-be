package auth

import "github.com/lukiriskigumilar/resepify-be/internal/users"

type AuthModule struct {
	Handler *AuthHandler
	Service AuthService
}

func InitAuthModule(usersModule *users.UserModule) *AuthModule {
	service := NewAuthService(usersModule.Repo)
	handler := NewAuthHandler(service)

	return &AuthModule{
		Handler: handler,
		Service: service,
	}
}
