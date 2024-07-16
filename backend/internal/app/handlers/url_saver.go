package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	service "github.com/goriiin/myapp/backend/internal/service "
	resp "github.com/goriiin/myapp/backend/pkg/api/response"
	"github.com/goriiin/myapp/backend/pkg/sl"
	"log/slog"
	"net/http"
)

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

func New(log *slog.Logger, serv service.URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "app.handlers.save.New"

		log = log.With(
			slog.String("op", op),
			//"request_id", r.Header.Get("X-Request-ID"),
		)

		var req service.Url

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Error("fail to decode request body", sl.Err(err))
			w.Header().Set("Content-Type", "application/json")

			_ = json.NewEncoder(w).Encode(resp.Error("fail to decode"))

			return
		}

		log.Info("request body decoded")

		// валидация обязательного поля
		if err = validator.New().Struct(req); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			log.Error("fail to validate request body", sl.Err(err))

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(resp.ValidationError(validationErrors))

			return
		}

		// TODO возврат alias
		if err = serv.SaveURL(req.Url, req.Alias); err != nil {
			log.Error("fail to decode request body", sl.Err(err))
			w.Header().Set("Content-Type", "application/json")

			_ = json.NewEncoder(w).Encode(resp.Error("fail to decode"))

			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Response{
			Response: *resp.OK(),
			//Alias:    req.Alias,
		})
	}
}
