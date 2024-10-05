package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sync"
	"time"
)

type PageData struct {
	Color string
	Name  string
}

var (
	colorCodes   = map[string]string{
		"red":      "#e74c3c",
		"green":    "#16a085",
		"blue":     "#2980b9",
		"blue2":    "#30336b",
		"pink":     "#be2edd",
		"darkblue": "#130f40",
	}
	currentColor = "blue" // default color
	mu           sync.Mutex // to ensure safe access to currentColor
)

// Handler for the root route
func mainHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	color := currentColor
	mu.Unlock()

	data := PageData{
		Color: colorCodes[color],
		Name:  os.Getenv("HOSTNAME"), // Get the hostname
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", mainHandler)

	// Start the HTTP server on port 80
	go func() {
		if err := http.ListenAndServe(":80", nil); err != nil {
			panic(err)
		}
	}()

	// Simple CLI to change the color
	for {
		var newColor string
		println("Enter a color (red, green, blue, blue2, pink, darkblue) or 'exit' to quit:")
		_, err := fmt.Scanln(&newColor)
		if err != nil {
			println("Error reading input:", err.Error())
			continue
		}

		if newColor == "exit" {
			break
		}

		mu.Lock()
		if _, exists := colorCodes[newColor]; exists {
			currentColor = newColor
			println("Background color changed to:", newColor)
		} else {
			println("Invalid color. Available colors: red, green, blue, blue2, pink, darkblue")
		}
		mu.Unlock()
	}

	println("Exiting...")
	time.Sleep(1 * time.Second) // Give some time to exit gracefully
}

