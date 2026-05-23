package middleware

import(
	"errors"
	"net/http"

	"github.com/vector15-05/GoCoin/api"
	"github.com/vector15-05/GoCoin/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or Token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		database, err := tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w, err)
			return
		}

		loginDetails := database.GetUserLoginDetails(username)

		if loginDetails == nil || token != loginDetails.AuthToken {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}