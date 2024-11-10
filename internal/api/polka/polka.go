package polka

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spossner/go-chirpy/internal/config"
	"github.com/spossner/go-chirpy/internal/utils"
	"net/http"
)

func HandleWebhook(cfg *config.ApiConfig) http.Handler {
	type request struct {
		Event string `json:"event"` //: "user.upgraded",
		Data  struct {
			UserId pgtype.UUID `json:"user_id"`
		} `json:"data"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := utils.GetPolkaToken(r)
		if !ok || token != cfg.PolkaKey {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		req, err := utils.Decode[request](r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if req.Event != "user.upgraded" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		_, err = cfg.Queries.SetRedUserById(r.Context(), req.Data.UserId)
		if err != nil {
			fmt.Println("error upgrading user:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
