package pkg

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"jwt/pkg"
	"net/http"
)

func Exec() {
	logger := zap.NewExample()
	defer logger.Sync()

	http.HandleFunc("/decode", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			//jwt, secret, public key
			jwt := request.URL.Query().Get("jwt")
			secret := request.URL.Query().Get("secret")
			if secret == "" {
				secret = pkg.DEFAULT_SECRET
			}
			err := validateParams(jwt)
			if err != nil {
				writer.WriteHeader(400)
				logger.Error("Error validating params", zap.Error(err))
				_, err := writer.Write([]byte(err.Error()))
				if err != nil {
					return
				}
			}

			logger.Info("Decoding jwt", zap.String("jwt", jwt))
			parse, err := pkg.Parse(jwt, secret)
			if err != nil {
				logger.Error("Error decoding jwt", zap.Error(err))
				writer.WriteHeader(400)
				_, err := writer.Write([]byte(err.Error()))
				if err != nil {
					return
				}
				return
			}
			jwtOk(writer, parse)
		}
	})

	http.HandleFunc("/encode", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("This is foo"))

	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		return
	}
}

func validateParams(jwt string) error {
	if jwt == "" {
		return errors.New("missing jwt")
	}
	return nil
}

func jwtOk(writer http.ResponseWriter, parse *jwt.Token) {
	writer.WriteHeader(200)
	writer.Header().Set("Content-Type", "application/json")
	printJWT := pkg.PrintJWT(parse, "json")
	_, err := writer.Write([]byte(printJWT))
	if err != nil {
		return
	}

}
