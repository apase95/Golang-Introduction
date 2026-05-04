package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const MaxUploadSize = 10 << 20

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		http.Error(w, "File is too large or invalid request", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("upload_file")
	if err != nil {
		http.Error(w, "Failed to retrieve the file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Printf("[UPLOAD INFO] Recieved file: %s (Size; %d bytes)\n", fileHeader.Filename, fileHeader.Size)

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
		return
	}
	dstPath := filepath.Join(uploadDir, fileHeader.Filename)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
		return
	}
	defer dstFile.Close()
	if _, err := io.Copy(dstFile, file); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "Success",
		"message": "File uploaded successfully",
		"filename": fileHeader.Filename,
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/upload", uploadFileHandler)
	port := ":8088"
	fmt.Printf("🚀 Upload Server is running at http://localhost%s\n", port)
	
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}