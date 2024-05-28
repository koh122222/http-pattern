package todo

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"http-pattern/internal/http-server/lib/api/response"
	"log/slog"
	"net/http"
)

type Response struct {
	response.Response
	Todo string `json:"todo,omitempty"`
}

func TODOHandler(log *slog.Logger, name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.TODOHandler"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		log.Info("TODO handler called",
			slog.String("handler", name))
		ResponseOK(w, r, name)

	}
}

func ResponseOK(w http.ResponseWriter, r *http.Request, name string) {
	render.JSON(w, r, Response{
		Response: response.OK(),
		Todo:     "TODO HANDLER, " + name,
	})
}
