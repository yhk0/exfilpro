// this code is very important (not exclude)

package util

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type DataLeak struct {
	Type       string // Type of data found (e.g., Email, Password, API Key)
	Match      string // The data found
	LineNumber int    // Number of the line where the data was found
	Path       string // Parsed file path or URL
}

// Loading displays a progress animation while the operation occurs.
func Loading(message string, stopChan chan bool) {
	ticker := time.NewTicker(3 * time.Second) // Animation interval
	defer ticker.Stop()

	frames := []string{"-", "\\", "|", "/"} // Animation frames
	i := 0

	for {
		select {
		case <-stopChan:
			fmt.Printf("\r%s...\n", message) // Final message
			os.Stdout.Sync()                 // Force Terminal Refresh
			return
		case <-ticker.C:
			fmt.Printf("\r%s %s", message, frames[i%len(frames)]) // View loading
			os.Stdout.Sync()                                      // Force Terminal Refresh
			i++
		}
	}
}

func AnalyzeWeb(url string, filter string) {
	log.SetFlags(0) // This removes timestamps from records
	stopChan := make(chan bool)
	go Loading("[INFO] Analyzing content", stopChan)

	log.Println("[INFO] Starting Web Mode Analysis...")
	content, err := FetchContent(url)
	if err != nil {
		log.Fatalf("[ERROR] Failed to fetch content from URL: %v", err)
	}

	// Stop the spinner after fetching the content
	stopChan <- true
	close(stopChan)
	log.Println("[INFO] Successfully obtained content. Starting Analysis...")

	// Call the analysis with the URL as the path
	leaks := ScanContent(content, url)

	// Apply filter if needed
	if filter != "" {
		leaks = FilterLeaks(leaks, filter)
	}

	// View the analysis results
	if len(leaks) == 0 {
		log.Println("[INFO] No leaks found.")
		return
	}

	fmt.Println("[#] Leaks found:")
	for _, leak := range leaks {
		fmt.Printf("[#] Type: %s\n", leak.Type)
		fmt.Printf("[#] Data: %s\n", leak.Match)
		fmt.Printf("[#] Line: %d\n", leak.LineNumber)
		fmt.Printf("[#] Local: %s\n", leak.Path)
		fmt.Println()
	}
	time.Sleep(3 * time.Second)
}

func AnalyzeFile(filePath string) {
	log.SetFlags(0)
	log.Println("\n[INFO] Starting File Mode Analysis...")

	stopChan := make(chan bool)
	go Loading("\n[INFO] Analyzing content", stopChan)
	// Get file information
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("[ERROR] Error accessing the file: %v", err)
	}

	// Stop the spinner after fetching the content
	stopChan <- true
	close(stopChan)

	fmt.Printf("\n[INFO] File Information:\n")
	fmt.Printf("- Name: %s\n", fileInfo.Name())
	fmt.Printf("- Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("- Last Modified: %s\n", fileInfo.ModTime().Format(time.RFC1123))
	fmt.Printf("- It is a directory: %v\n", fileInfo.IsDir())
	time.Sleep(3 * time.Second)
}

// FetchContent gets the HTML content from a URL
func FetchContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("\n[ERR] erro ao acessar o site: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("\n[ERR] erro: status HTTP %d", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return "", fmt.Errorf("\n[ERR] erro ao ler o corpo da resposta: %v", err)
	}

	return buf.String(), nil
}

// ScanContent analyzes content and returns possible classified leaks
func ScanContent(content string, path string) []DataLeak {
	// Map patterns for each sensitive data type
	patterns := map[string]*regexp.Regexp{
		"Email":     regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`),
		"Senha":     regexp.MustCompile(`(?i)(password|senha)\s*[:=]\s*["']?[\w!@#$%^&*()-+=]{6,}["']?`),
		"API Key":   regexp.MustCompile(`(?i)(api_key|token|key)\s*[:=]\s*["']?[a-zA-Z0-9]{20,}["']?`),
		"reCAPTCHA": regexp.MustCompile(`[A-Za-z0-9_-]{40}`),
	}

	var leaks []DataLeak

	// Separate content into rows for analysis
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		for label, pattern := range patterns {
			matches := pattern.FindAllString(line, -1)
			for _, match := range matches {
				// Add the leak found with line number and path
				leaks = append(leaks, DataLeak{
					Type:       label,
					Match:      match,
					LineNumber: i + 1, // Line number starts at 1
					Path:       path,  // Directory or URL parsed
				})
			}
		}
	}

	return leaks
}

// FilterLeaks filters the list of leaks based on the specified type
func FilterLeaks(leaks []DataLeak, filter string) []DataLeak {
	var filteredLeaks []DataLeak

	for _, leak := range leaks {
		// Checks that the type of leak corresponds to the filter (case insensitive)
		if strings.EqualFold(leak.Type, filter) {
			filteredLeaks = append(filteredLeaks, leak)
		}
	}

	return filteredLeaks
}
