package models

type Note struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	User     string `json:"user"`
	Password string `json:"password"`
}
type NoteResponse struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type GetRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// speller структура ответа
type SpellError struct {
	Code    int      `json:"code"`
	Pos     int      `json:"pos"`
	Row     int      `json:"row"`
	Col     int      `json:"col"`
	Len     int      `json:"len"`
	Word    string   `json:"word"`
	Suggest []string `json:"s"`
}
