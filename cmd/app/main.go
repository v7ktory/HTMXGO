package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
	"github.com/v7ktory/htmx+go/handler"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	supabase := supa.CreateClient(os.Getenv("SUPA_URL"), os.Getenv("SUPA_KEY"))
	h := handler.NewHandler(supabase)
	h.InitRoute().Run(":8000")
}
