package main

import (
	"fmt"
	"log"
	"net/http"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/download" {
		http.Error(w, r.URL.Path+" not found!", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// f, err := os.Open(yourFile)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()

	// w.Header().Set("Content-Disposition", "attachment; filename=YourFile")
	// w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	// io.Copy(w, f)

	fmt.Fprintf(w, "Here is a list of available images!")
}

func main() {
	http.HandleFunc("/download", downloadHandler) // Update this line of code

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
