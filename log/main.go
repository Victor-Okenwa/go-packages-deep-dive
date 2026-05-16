package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("=== Go log Package Deep Dive ===")

	// 1. Default Logger (writes to stderr)
	log.Println("This is a basic log message")
	log.Printf("User %s logged in at %s", "Victor", time.Now().Format("15:04:05"))

	// 2. Custom Logger - Most Important Pattern
	fmt.Println("\n--- Custom Loggers ---")

	// Create different loggers for different purposes
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger := log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLogger.Println("Application started successfully")
	warnLogger.Println("Database connection is slow")
	errorLogger.Println("Failed to connect to payment gateway")

	// 3. Common Log Flags (Very Important)
	fmt.Println("\n--- Log Flags Explained ---")

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	log.Println("This log has date, time with microseconds, filename, and UTC")

	// 4. Different Output Destinations
	fmt.Println("\n--- Writing Logs to File ---")

	// Create a log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a logger that writes to file
	fileLogger := log.New(file, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)

	// 5. Fatal and Panic
	fmt.Println("\n--- log.Fatal and log.Panic ---")

	// log.Fatal("This will print and then call os.Exit(1)")  // Uncomment to test
	// log.Panic("This will print and then panic")            // Uncomment to test

	// 6. Best Practice: Global Logger Setup
	fmt.Println("\n--- Recommended Pattern ---")
	setupLogger(fileLogger)
	log.Println("This uses our configured global logger")
}

// Recommended way to setup logger in real projects
func setupLogger(fileLogger *log.Logger) {
	log.SetPrefix("DEEP DIVE: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout) // You can change to multi-writer (file + stdout)
	fileLogger.Println("This message is written to app.log file")
}
