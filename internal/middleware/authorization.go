package middleware

import (
	"errors"
	"net/http"

	"github.com/SaiAungMinKhant/goAPI/api"
	"github.com/SaiAungMinKhant/goAPI/internal/tools"
	log "github.com/sirupsen/logrus"
)

var ErrUnAuthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(ErrUnAuthorized)
			api.RequestErrorHandler(w, ErrUnAuthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		// var loginDetails *tools.LoginDetails
		// loginDetails =  (*database).GetUserLoginDetails(username)
		loginDetails := (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(ErrUnAuthorized)
			api.RequestErrorHandler(w, ErrUnAuthorized)
			return
		}

		next.ServeHTTP(w, r)

	})
}
