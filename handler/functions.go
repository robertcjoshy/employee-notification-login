package handler

func Authenticated(user, pass string) bool {
	if user == "robert" && pass == "1234" {
		return true
	}
	return false
}
