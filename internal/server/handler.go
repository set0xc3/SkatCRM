package server

import (
	"io"
	"net/http"
)

func RedirectToAPI(w http.ResponseWriter, r *http.Request) {
	// Создаем новый запрос к серверу http://localhost:8090
	req, err := http.NewRequest(r.Method, "http://localhost:8090"+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Копируем заголовки из оригинального запроса
	for key, value := range r.Header {
		req.Header[key] = value
	}

	// Отправляем запрос на сервер http://localhost:8090
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close() // Закрываем тело ответа

	// Устанавливаем заголовки ответа
	for key, value := range resp.Header {
		w.Header()[key] = value
	}

	// Устанавливаем статус ответа
	w.WriteHeader(resp.StatusCode)

	// Копируем тело ответа клиенту
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
