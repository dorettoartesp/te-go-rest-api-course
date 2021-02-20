package main

import "fmt"

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

// Runs - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our app...")
	return nil
}

func main() {
	fmt.Println("GO Rest API")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
