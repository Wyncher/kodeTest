package speller

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"main/models"
	"net/http"
	"net/url"
	"strings"
)

func CheckText(text string) string {
	// Создаем HTTP-клиент с отключенной проверкой сертификатов
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Отключение проверки TLS
		},
	}

	// Текст для проверки
	//text = "Привет, как дила?"

	// Формируем параметры запроса
	data := url.Values{}
	data.Set("text", text)
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

	// Декодируем ответ
	var spellErrors []models.SpellError
	if err := json.NewDecoder(resp.Body).Decode(&spellErrors); err != nil {
		log.Fatalf("Ошибка при декодировании ответа: %v", err)
	}
	// Если ошибок нет, выводим оригинальный текст
	if len(spellErrors) == 0 {
		fmt.Println("Ошибки не найдены!")
		fmt.Println("Текст: ", text)
		return text
	}

	// Исправляем текст, заменяя слова на первые предложения
	updatedText := text
	for i := len(spellErrors) - 1; i >= 0; i-- { // Идем с конца, чтобы не сместить индексы
		spellError := spellErrors[i]
		if len(spellError.Suggest) > 0 {
			// Берем первое предложенное исправление
			suggestion := spellError.Suggest[0]

			// Заменяем ошибочное слово на исправление
			updatedText = updatedText[:spellError.Pos] + suggestion + updatedText[(spellError.Pos+spellError.Len)*2:]
		}
	}
	return updatedText
}
