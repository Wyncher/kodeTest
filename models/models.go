package models

// Структура заметки
type Note struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Заметка,которая отдается пользователю по его запросу
type NoteResponse struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Структура,которая приходит от пользователя вместе с запросом на получение заметок
type GetRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// Speller структура ответа от API
type SpellError struct {
	Code    int      `json:"code"`
	Pos     int      `json:"pos"`
	Row     int      `json:"row"`
	Col     int      `json:"col"`
	Len     int      `json:"len"`
	Word    string   `json:"word"`
	Suggest []string `json:"s"`
}
