package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"./rednosio/handlers"
)

func main() {
	args := os.Args
	port := 8080

	if 3 == len(args) && "--port" == args[1] {
		port, _ = strconv.Atoi(args[2])
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		http.HandleFunc("/", handlers.Error(handlers.Index))
		http.HandleFunc("/browse/downloads", handlers.Error(handlers.BrowseDownloads))
		http.HandleFunc("/browse/uploads", handlers.Error(handlers.BrowseUploads))
		http.HandleFunc("/delete/download", handlers.Error(handlers.DeleteDownload))
		http.HandleFunc("/image", handlers.Error(handlers.Image))
		http.HandleFunc("/rednosify", handlers.Error(handlers.Rednosify))
		http.HandleFunc("/save", handlers.Error(handlers.SaveImage))
		http.HandleFunc("/thumb/download", handlers.Error(handlers.DownloadThumb))
		http.HandleFunc("/thumb/upload", handlers.Error(handlers.UploadThumb))

		http.Handle("/downloads/", http.StripPrefix("/downloads/", http.FileServer(http.Dir("./downloads"))))
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		wg.Done()
	}()

	log.Printf("Listening at \":%d\"\n", port)
	log.Println("CTRL-C to exit...")
	wg.Wait()
}
