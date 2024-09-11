package speller

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"main/models"
	"net/http"
	"net/url"
	"strings"
)

func CheckText(text string) string {
	// Создаем HTTP-клиент с отключенной проверкой сертификатов(с включенной проверкой сертификат не действителен)
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Отключение проверки TLS
		},
	}
	// Формируем параметры запроса
	data := url.Values{}
	data.Set("text", text)
	//Язык Русский
	data.Set("lang", "ru")

	// Отправляем запрос на Yandex Speller API
	resp, err := httpClient.Post(
		"https://speller.yandex.net/services/spellservice.json/checkText",
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	// Декодирирование ответа
	var spellErrors []models.SpellError
	if err := json.NewDecoder(resp.Body).Decode(&spellErrors); err != nil {
		log.Fatalf("Ошибка при декодировании ответа: %v", err)
	}
	// Если ошибок нет, возвращается оригинальный текст
	if len(spellErrors) == 0 {
		return text
	}
	// Если ошибки есть, то выполняются замены
	updatedText := text
	for i := len(spellErrors) - 1; i >= 0; i-- { // подсчёт с конца, чтобы не сместить индексы символов в начале строки.
		spellError := spellErrors[i]
		runesUpdatedText := []rune(updatedText)
		if len(spellError.Suggest) > 0 {
			// Выбор первого предложенного исправления
			suggestion := spellError.Suggest[0]
			// Замена ошибочного слово на исправленное(с конвертацией в руны)
			updatedText = strings.Replace(updatedText, string(runesUpdatedText[spellError.Pos:(spellError.Pos+spellError.Len)]), suggestion, 1)
		}
	}
	return updatedText
}
