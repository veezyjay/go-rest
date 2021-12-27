package main

import "fmt"

// App - the struct which contains things like pointers to
// database connections
type App struct{}

// Run - sets up the app
func (app *App) Run() error {
	fmt.Println("Setting up the app")
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
