package authentification

// карта для хранения данных пользователя
var users map[string]string

func Authentificate(login string, pswd string) bool {
	//тут может быть проверка пользователя в БД
	//проверка на наличие пользователя
	users = make(map[string]string)
	//добавление двух пользователей
	users["user"] = "password"
	users["user1"] = "password"
	//проверка(или ее подобие) на наличие пользователя в списке
	if users[login] == pswd {
		//Пользователь найден, значит его заметки где-то должны быть...
		return true
	}
	//Если пользователя нет, значит он не оставлял заметку и продолжить не может
	return false
}
