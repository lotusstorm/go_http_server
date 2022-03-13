package src

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Config Configuration
	Router *mux.Router
	DB     *gorm.DB
}

func (app *App) Configure() {
	app.configureEnv()
	app.configureDB()
	app.configureRoutes()
	app.configureMiddleware()
}

func (app *App) Run() {
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%s", app.Config.Host, app.Config.Port),
			handlers.LoggingHandler(os.Stdout, app.Router),
		),
	)
}

func (app *App) configureEnv() {
	app.Config.loadEnv()
}

func (app *App) configureDB() {
	var err error
	app.DB, err = gorm.Open(postgres.Open(app.Config.DBString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	_ = app.DB.AutoMigrate(&Book{})
}

func (app *App) configureMiddleware() {
	app.Router.Use(commonMiddleware)
}

func (app *App) configureRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/book", app.getAllBooks).Methods(http.MethodGet)
	app.Router.HandleFunc("/book", app.addBook).Methods(http.MethodPost)
	app.Router.HandleFunc("/book/{book_id:[0-9]+}", app.getBook).Methods(http.MethodGet)
	app.Router.HandleFunc("/book/{book_id:[0-9]+}", app.updateBook).Methods(http.MethodPut)
	app.Router.HandleFunc("/book/{book_id:[0-9]+}", app.deleteBook).Methods(http.MethodDelete)
}
