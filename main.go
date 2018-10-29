package main

import (
	"fmt"
	"github.com/dan-ancora/apiteamamerica"
	"google.golang.org/appengine" // Required external App Engine library
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, "<h1>Hello, Google Api Engine!</h1>")
}

func getCitiesList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start testing Team America Package")

	taClient := &apiteamamerica.Taclient{
		Username: "darkcomputer",
		Password: "12345",
		URL:      "https://javatest.teamamericany.com:8443/TADoclit/services/TADoclit?wsdl"}
	//URL: "http://82.77.142.41:9999/ancoraerp/index.jsp"}

	/*
		ok, err := taClient.Connect()
		if err != nil {
			fmt.Println("Error connecting to url: ", taClient.URL)
		} else {
			fmt.Println("Response: ", ok)
		}
	*/

	response, err := taClient.ListCities(r)
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write([]byte(response))
	}

}

func main() {
	http.HandleFunc("/city_list", getCitiesList)

	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}
