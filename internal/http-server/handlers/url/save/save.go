package save

import (
	"log/slog"
	"net/http"
)

type Request struct {
	URL   string `json:"url" validate:"required, url"`
	Alias string `json:"alias,omitempty"`
}

// Если ты указываешь alias, то он будет сохранен
// А если не указываешь, то будет пустой
type Response struct {
	Status string `json:"status"` // Error, Ok
	Error  string `json:"error,omitempty"`
	Alias  string `json:"alias,omitempty"` // Alias только что сохраненного url
}
type URLSaver interface {
	SaveURL(urlTOSave string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
