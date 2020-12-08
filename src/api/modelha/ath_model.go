package modelha

type User struct {
	username string
	passwd string
}

type sessionData struct {
	User
	LoogedIn bool
	loginFial bool
}