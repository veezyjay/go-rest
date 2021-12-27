package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/veezyjay/go-rest/internal/transport/http"
)

// App - the struct which contains things like pointers to
// database connections
type App struct{}

// Run - sets up the app
func (app *App) Run() error {
	fmt.Println("Setting up the app")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go Rest API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up the Rest API")
		fmt.Println(err)
	}
}
