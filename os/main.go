package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {

	fmt.Println("=== Go os Package Deep Dive ===")

	// 1. Command-line Arguments
	fmt.Println("1. Command Line Arguments:")
	fmt.Printf("Executable name: %s\n", os.Args[0])
	fmt.Printf("Total arguments: %d\n", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("Arg[%d]: %s\n", i, arg)
	}

	// 2. Environment Variables (Very Important)
	fmt.Println("\n2. Environment Variables:")

	// Reading
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	fmt.Printf("ENVIRONMENT = %s\n", env)

	// Setting (only affects current process)
	os.Setenv("APP_VERSION", "1.2.3")
	fmt.Printf("APP_VERSION = %s\n", os.Getenv("APP_VERSION"))

	// 3. Working with Files - Modern way (Go 1.16+)
	fmt.Println("\n3. File Operations:")

	data := []byte("Hello, Go os package!\n")
	// Write file (simple)
	err := os.WriteFile("test.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("✅ File 'test.txt' created successfully")

	// Read file
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Printf("File content:\n%s\n", string(content))

	// 4. Directories
	fmt.Println("\n4. Directory Operations:")

	// Create directory
	os.MkdirAll("logs/app", 0755)
	fmt.Println("✅ Directory 'logs/app' created")

	// Get current working directory
	cwd, _ := os.Getwd()
	fmt.Printf("Current working directory: %s\n", cwd)

	// 5. File Information
	fmt.Println("\n5. File Info:")
	info, err := os.Stat("test.txt")
	if err == nil {
		fmt.Printf("File Name: %s\n", info.Name())
		fmt.Printf("File Size: %d bytes\n", info.Size())
		fmt.Printf("Is Directory: %t\n", info.IsDir())
		fmt.Printf("Permissions: %s\n", info.Mode())
	}

	// 6. Path Manipulation (Best Practice)
	fmt.Println("\n6. Path Handling (Cross-platform):")
	fullPath := filepath.Join(cwd, "logs", "app", "error.log")
	fmt.Printf("Full path: %s\n", fullPath)
	fmt.Printf("Base: %s\n", filepath.Base(fullPath))
	fmt.Printf("Dir: %s\n", filepath.Dir(fullPath))
	fmt.Printf("Extension: %s\n", filepath.Ext(fullPath))

	// 7. Host and User Information
	fmt.Println("\n7. System Information:")
	fmt.Printf("GOOS: %s\n", runtime.GOOS) // windows, linux, darwin
	fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
	fmt.Printf("COMPILER: %s\n", runtime.Compiler)
	hostname, _ := os.Hostname()
	fmt.Printf("Hostname: %s\n", hostname)
}
