package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler показывает главную страницу с формой
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загруженный файл
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Читаем форму (ограничиваю размер 10 МБ)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка при парсинге формы", http.StatusInternalServerError)
		return
	}

	// Получаем файл из формы
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Обрабатываем содержимое файла в память
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	// Конвертируем данные
	result, err := service.AutoConvert(string(data))
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаю локальный файл для результата (беру расширение оригинала и текущее время)
	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("result_%s%s", time.Now().UTC().Format("20060102_150405"), ext)

	// Записываем результат
	err = os.WriteFile(fileName, []byte(result), 0644)
	if err != nil {
		http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ в браузер (что всё прошло успешно)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "%s", result)
}