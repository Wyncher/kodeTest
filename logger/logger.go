package logger

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

// Создание экземпляра logrus
var log = logrus.New()
var File *os.File

// CustomLoggerMiddleware - логирование запросов с использованием logrus
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Получение статуса запроса
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		// Выполнение запроса
		next.ServeHTTP(ww, r)

		// Логирование информации о запросе
		log.WithFields(logrus.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   ww.Status(),
			"duration": time.Since(start),
			"ip":       r.RemoteAddr,
		}).Info("Handled request")
	})
}

// logError - логгирование пользовательских ошибок с использованием logrus
func LogError(r *http.Request, err error) {
	if r != nil {
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"error":  err,
			"ip":     r.RemoteAddr,
		}).Error(err)
	} else {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error(err)
	}

}

// TableFormatter кастомный форматтер для logrus, который выводит логи в виде таблицы
type TableFormatter struct{}

// Format реализует интерфейс logrus для вывода логов в формате таблицы
func (f *TableFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Получаем текущую временную метку
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	// Форматирование лога в виде строки таблицы
	logMessage := fmt.Sprintf(
		"| %-20s | %-7s | %-10s | %-30s | %-20s | %-20s |\n",
		timestamp,
		entry.Level.String(),
		entry.Data["method"],
		entry.Message,
		entry.Data["ip"],
		entry.Data["path"],
	)
	File.WriteString(logMessage)

	return []byte(logMessage), nil
}
func Setup() {
	//Задание табличного форматирование
	log.SetFormatter(&TableFormatter{})
	// Получение времени
	timeNow := time.Now()
	logFileName := "logs/" + timeNow.Format("15-04-05") + ".log"
	//Создание файла с логом, который будет иметь имя означающее
	//время запуска сервера и иметь формат log
	f, err := os.Create(logFileName)
	if err != nil {
		fmt.Println(err)
		LogError(nil, err)
	}
	f.Close()
	//Открытие файла в глобальную переменную
	File, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		LogError(nil, err)
	}
	File.WriteString("| Время                | Уровень | Метод      | Сообщение                      | IP адрес             | Путь                 |")
	File.WriteString("|----------------------|---------|------------|--------------------------------|----------------------|----------------------|")
	// Заголовок таблиц
	fmt.Println("| Время                | Уровень | Метод      | Сообщение                      | IP адрес             | Путь                 |")
	fmt.Println("|----------------------|---------|------------|--------------------------------|----------------------|----------------------|")

}
