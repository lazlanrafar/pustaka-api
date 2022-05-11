package loginservice

//Login credential
type LoginCredentials struct {
	Username    string `form:"username"`
	Password string `form:"password"`
}