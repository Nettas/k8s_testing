package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	// MAIN SECTION HTML CODE
	fmt.Fprintf(w, "<h1> k8s with CCP on HX a Better Together Story!</h1>")
	fmt.Fprintf(w, "<img src='cisco.jpg' alt='CISCO' > ")

}

func about_handler(w http.ResponseWriter, r *http.Request) {
	// ABOUT SECTION HTML CODE
	fmt.Fprintf(w, "<title>Go/about/</title>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.Handle("/code", http.StripPrefix("/code", http.FileServer(http.Dir("./cisco.jpg"))))
	http.ListenAndServe(":8000", nil)

	// 	existingImageFile, err := os.Open("cisco.png")
	// 	if err != nil {

	// 		fmt.Println(cisco)
	// 		fmt.Println()
	// 		existingImageFile.Seek(0, 0)
	// 		loadedImage, err := png.Decode(existingImageFile)
	// 		if err != nil {
	// 			// Handle error
	// 		}
	// 		fmt.Println(loadedImage)
	// 	}
}
