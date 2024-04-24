package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type streamReader struct {
	reader io.ReadCloser
}

func (s *streamReader) Read(p []byte) (n int, err error) {
	return s.reader.Read(p)
}

func handleStream(w http.ResponseWriter, r *http.Request) {
	// Generate dynamic timestamp
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// Build original URL
	originalURL := fmt.Sprintf("https://an.cdn.eurozet.pl/ant-web.mp3?t=%d", timestamp)

	// Create a new HTTP client
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	// Send request with automatic redirection disabled
	req, err := http.NewRequest(http.MethodGet, originalURL, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error fetching stream: %d", resp.StatusCode), http.StatusBadGateway)
		return
	}

	// Get the final URL (no redirection)
	finalURL := resp.Request.URL.String()

	fmt.Printf("Playing stream from: %s\n", finalURL)

	// Set content type header for mpv compatibility
	w.Header().Set("Content-Type", "audio/mpeg")

	// Create a reader from the response body
	reader := &streamReader{reader: resp.Body}

	// Allocate a buffer for data transfer (adjust size as needed)
	buffer := make([]byte, 1024)

	// Continuously read from radio stream and copy to response
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Write the read data to response writer
		_, err = w.Write(buffer[:n])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handleStream)
	http.ListenAndServe(":8088", nil)
}
