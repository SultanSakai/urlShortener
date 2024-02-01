package redirect

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	resp "urlShortener/internal/lib/api/response"
	"urlShortener/internal/storage"
)

type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	const fn = "handlers.redirect.New"
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(context.Background())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Error("alias is empty")
			render.JSON(w, r, resp.Error("invalid request"))
			return
		}

		resURL, err := urlGetter.GetURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Error("url not found", "alias", alias)
			render.JSON(w, r, resp.Error("url not found"))
			return
		}
		if err != nil {
			log.Error("failed to get url", "alias", alias)
			render.JSON(w, r, resp.Error("internal error"))
			return
		}
		log.Info("url found", "alias", alias, "url", resURL)

		http.Redirect(w, r, resURL, http.StatusTemporaryRedirect)
	}
}
