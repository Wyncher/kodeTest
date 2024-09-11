package authentification

// карта для хранения данных пользователя
var users map[string]string

func Authentificate(login string, pswd string) bool {
	//тут может быть проверка пользователя в БД
	//если пользователь имеет такие данные
	if users[login] == pswd {
		return true
	}
	return true
	//return false
}
