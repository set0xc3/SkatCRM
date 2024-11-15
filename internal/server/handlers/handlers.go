package handlers

import (
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func StatusOK(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func StatusNoContent(c echo.Context) error {
	return c.String(http.StatusNoContent, "")
}

func RedirectToDB(c echo.Context) error {
	// Создаем новый запрос к серверу http://localhost:8090
	req, err := http.NewRequest(c.Request().Method, "http://localhost:8090"+c.Request().URL.Path, c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer c.Request().Body.Close()

	// Копируем заголовки из оригинального запроса, исключая некоторые
	for key, value := range c.Request().Header {
		if key != "Content-Length" && key != "Transfer-Encoding" {
			req.Header[key] = value
		}
	}

	// Отправляем запрос на сервер http://localhost:8090
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close() // Закрываем тело ответа

	// Устанавливаем статус ответа
	c.Response().WriteHeader(resp.StatusCode)

	// Копируем заголовки ответа
	for key, value := range resp.Header {
		c.Response().Header()[key] = value
	}

	// Убедитесь, что Content-Type установлен правильно
	if resp.Header.Get("Content-Type") == "" {
		c.Response().Header().Set("Content-Type", "application/json")
	}

	// Копируем тело ответа клиенту
	if _, err := io.Copy(c.Response(), resp.Body); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}
