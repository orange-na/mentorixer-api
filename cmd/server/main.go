package server

import (
	"net/http"

	"main/internal/handler"
	"main/internal/repository"
	"main/pkg/db"

	"github.com/rs/cors"
)

func Run() {
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// リポジトリの初期化
	taskRepo := repository.NewTaskRepository(db)
	userRepo := repository.NewUserRepository(db)

	// ハンドラーの初期化
	taskHandler := handler.NewTaskHandler(taskRepo)
	userHandler := handler.NewUserHandler(userRepo)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	http.Handle("/tasks", c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetTasks(w, r)
		case http.MethodPost:
			taskHandler.AddTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))
	http.Handle("/tasks/", c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			taskHandler.EditTask(w, r)
		case http.MethodDelete:
			taskHandler.DeleteTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	http.Handle("/sign-up", c.Handler(http.HandlerFunc(userHandler.SignUp)))
	http.Handle("/sign-in", c.Handler(http.HandlerFunc(userHandler.SignIn)))
	http.Handle("/users", c.Handler(http.HandlerFunc(userHandler.GetUsers)))

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	server.ListenAndServe()
}