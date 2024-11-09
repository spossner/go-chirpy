package ping

import (
	"github.com/spossner/go-chirpy/internal/utils"
	"net/http"
)

func HandlePing() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := utils.Encode(w, struct {
			Result string `json:"result"`
		}{"OK"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})
}
