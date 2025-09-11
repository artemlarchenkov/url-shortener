package delete

import (
	"errors"
	"log/slog"
	"net/http"
	"url-shortener/internal/storage/sqllite"

	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

// Request структура запроса на удаление
type Request struct {
	Alias string `json:"alias" validate:"required"` // Alias обязателен
}

// Response структура ответа
type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=URLDeleter
type URLDeleter interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urldeleter *sqllite.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}

		// Валидация
		v := validator.New()
		if err := v.Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("failed to validate request", sl.Err(err))
			render.JSON(w, r, resp.ValidationError(validateErr))
			return
		}

		// Удаление URL
		err := urldeleter.DeleteURL(req.Alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("alias not found", slog.String("alias", req.Alias))
			render.JSON(w, r, resp.Error("alias not found"))
			return
		}
		if err != nil {
			log.Error("failed to delete url", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to delete url"))
			return
		}

		log.Info("url deleted", slog.String("alias", req.Alias))
		render.JSON(w, r, Response{
			Response: resp.OK(),
			Alias:    req.Alias,
		})
	}
}
