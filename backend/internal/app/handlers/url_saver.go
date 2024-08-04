package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/goriiin/myapp/backend/internal/service"
	resp "github.com/goriiin/myapp/backend/pkg/api/response"
	"github.com/goriiin/myapp/backend/pkg/sl"
	"log/slog"
	"net/http"
)

type Response struct {
	resp.Response
	URL   string `json:"url"`
	Alias string `json:"alias,omitempty"`
}

func Save(log *slog.Logger, serv service.URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "app.handlers.save"

		log = log.With(
			slog.String("op", op),
			//"request_id", r.Header.Get("X-Request-ID"),
		)

		var req service.Url

		err := json.NewDecoder(r.Body).Decode(&req)

		fmt.Println(req)

		if err != nil {
			log.Error("fail to decode request body", sl.Err(err))
			w.Header().Set("Content-Type", "application/json")

			_ = json.NewEncoder(w).Encode(resp.Error("fail to decode"))

			return
		}

		log.Info("request body decoded")

		// валидация обязательного поля
		if err = validator.New().Struct(req); err != nil {
			var validationErrors validator.ValidationErrors
			errors.As(err, &validationErrors)
			log.Error("fail to validate request body", sl.Err(err))

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(resp.ValidationError(validationErrors))

			return
		}

		if newAlias, err := serv.SaveURL(req.Url, req.Alias); err != nil {
			if newAlias != nil {
				req.Alias = *newAlias
			} else {
				log.Error("fail to decode request body", sl.Err(err))
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(resp.Error("fail to decode"))

				return
			}

		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Response{
			Response: *resp.OK(),
			URL:      req.Url,
			Alias:    req.Alias,
		})
	}
}

func Edit(log *slog.Logger, serv service.URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.handlers.edit"

		log = log.With(
			slog.String("op", op),
		)
		var req service.Url

		err := json.NewDecoder(r.Body).Decode(&req)

		fmt.Println(req)

		if err != nil {
			log.Error("fail to decode request body", sl.Err(err))
			w.Header().Set("Content-Type", "application/json")

			_ = json.NewEncoder(w).Encode(resp.Error("fail to decode"))

			return
		}

		log.Info("request body decoded")

		// валидация обязательного поля
		if err = validator.New().Struct(req); err != nil {
			var validationErrors validator.ValidationErrors
			errors.As(err, &validationErrors)
			log.Error("fail to validate request body", sl.Err(err))

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(resp.ValidationError(validationErrors))

			return
		}

		if newAlias, err := serv.EditURL(req.Url, req.Alias); err != nil {
			if newAlias != nil {
				req.Alias = *newAlias
			} else {
				log.Error("fail to decode request body", sl.Err(err))
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(resp.Error("fail to decode"))

				return
			}

		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Response{
			Response: *resp.OK(),
			URL:      req.Url,
			Alias:    req.Alias,
		})
	}
}

func Delete(log *slog.Logger, serv service.URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.handlers.delete"

		log = log.With(
			slog.String("op", op))

		query := r.URL.Query()
		url := query.Get("url")

		if url == "" {
			log.Error("fail to get url from query-params", sl.Err(errors.New("url is required")))
			_ = json.NewEncoder(w).Encode(Response{
				Response: *resp.Error("empty query params"),
			})
			return
		}

		err := serv.RemoveURL(url)
		if err != nil {
			log.Error("fail to remove url", sl.Err(err))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Response{
			Response: *resp.OK(),
			URL:      url,
		})
	}
}

func Get(log *slog.Logger, serv service.URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.handlers.get"
		log = log.With(
			slog.String("op", op),
		)

		query := r.URL.Query()
		alias := query.Get("alias")
		if alias == "" {
			log.Error("fail to get alias from query-params", sl.Err(errors.New("url is required")))
			_ = json.NewEncoder(w).Encode(Response{
				Response: *resp.Error("empty query params"),
			})

			return
		}

		url, err := serv.GetURL(alias)
		if err != nil {
			log.Error("fail to get url from DB", sl.Err(err))
			return
		}
		if url.Url == "" {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(Response{
				Response: *resp.Error("no url"),
				Alias:    url.Alias,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Response{
			Response: *resp.OK(),
			URL:      url.Url,
			Alias:    alias,
		})
	}
}
