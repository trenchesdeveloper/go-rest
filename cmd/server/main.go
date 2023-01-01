package main

import "fmt"

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run(addr string) error {
	fmt.Println("Setting Up Our App")
	return nil
}

func main() {
	fmt.Println("Go rest api server")
	app := App{}

	if err := app.Run(":3000"); err != nil {
		fmt.Println("Error starting server", err)
	}
}