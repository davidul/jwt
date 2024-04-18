package pkg

import (
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
			if jwt == "" {
				logger.Error("Missing jwt")
				writer.WriteHeader(400)
				write, err := writer.Write([]byte("Missing jwt"))
				if err != nil {
					logger.Error("Error writing response", zap.Error(err))
					return
				}
				logger.Info("Missing jwt", zap.Int("bytes", write))
				return
			}

			if secret == "" {
				secret = pkg.DEFAULT_SECRET
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
			writer.WriteHeader(200)
			writer.Header().Set("Content-Type", "application/json")
			printJWT := pkg.PrintJWT(parse, "json")
			write, err := writer.Write([]byte(printJWT))
			if err != nil {
				return
			}
			logger.Info("Decoded jwt", zap.Int("bytes", write))
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
