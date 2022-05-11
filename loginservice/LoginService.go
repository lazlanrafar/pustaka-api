package loginservice

type LoginService interface {
	LoginUser(username, password string) bool
}

type loginInfomation struct {
	username string
	password string
}

func StaticLoginService() LoginService {
	return &loginInfomation{
		username: "admin",
		password: "pass1234",
	}
}

func (l *loginInfomation) LoginUser(username, password string) bool {
	if username == l.username && password == l.password {
		return true
	}
	return false
}