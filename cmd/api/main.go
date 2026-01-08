package main

import (
	"app/internal/env"
	"log"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	r := app.mount()

	log.Fatal(app.run(r))

}
