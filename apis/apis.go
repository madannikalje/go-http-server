package apis

import (
	"encoding/json"
	"fmt"
	"http-test/config"
	"http-test/database/repository"
	"http-test/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRouter(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
			var bodyUser models.User
			if err := json.NewDecoder(r.Body).Decode(&bodyUser); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			existingUser := models.User{}
			if err := repository.FindUserByEmail(bodyUser.Email, &existingUser); err == nil {
				http.Error(w, "User already exists", http.StatusBadRequest)
				return
			}

			userToCreate := &models.User{
				Name:  bodyUser.Name,
				Email: bodyUser.Email,
			}

			if err := repository.CreateUser(userToCreate); err != nil {
				fmt.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email": bodyUser.Email,
			})

			if tokenString, err := token.SignedString([]byte(config.JWT_SECRET)); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			} else {
				w.Write([]byte(tokenString))
			}
		})
	})
}
