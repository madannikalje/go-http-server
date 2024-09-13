package bootstraps

import (
	"http-test/apis"
	"http-test/config"
	"http-test/database"
	"sync"

	"github.com/fatih/color"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var wg = sync.WaitGroup{}

func initializeHttpServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	apis.AuthRouter(r)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API"))
	})
	http.ListenAndServe(":3000", r)
	wg.Done()
}

func Init() {
	config.Init()
	database.InitializeDatabase()
	wg.Add(1)
	go initializeHttpServer()
	color.Cyan("Server started at http://localhost:3000")
	wg.Wait()
}
