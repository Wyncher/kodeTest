package apis

import (
	"encoding/json"
	authentification "main/auth"
	"main/models"
	"main/speller"
	"net/http"
	"sync"
)

var (
	notes         = []models.Note{}
	notesResponse = []models.NoteResponse{}
	nextID        = 1
	notesMu       sync.Mutex
)

// Обработчик для получения списка заметок
func GetNotes(w http.ResponseWriter, r *http.Request) {
	var auth models.GetRequest
	// извлекаем данные из запроса
	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	notesMu.Lock()
	defer notesMu.Unlock()
	if authentification.Authentificate(auth.User, auth.Password) {
		for noteNum := range notes {
			if notes[noteNum].User == auth.User && notes[noteNum].Password == auth.Password {
				notesResponse = append(notesResponse, models.NoteResponse{ID: notes[noteNum].ID, Text: notes[noteNum].Text})
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notesResponse)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("error")
	}

}

// Обработчик для создания новой заметки
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var newNote models.Note

	// извлекаем данные из запроса
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	notesMu.Lock()
	defer notesMu.Unlock()

	// Присваиваем ID и добавляем заметку в список
	newNote.ID = nextID
	newNote.Text = speller.CheckText(newNote.Text)
	nextID++
	notes = append(notes, newNote)

	// Возвращаем добавленную заметку в ответе
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newNote)
}
