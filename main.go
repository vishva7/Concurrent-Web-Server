package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchDataFromDatabase(requestID int) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Data for request #%d", requestID)
}

func processData(data string) string {
	return fmt.Sprintf("Processed: %s", data)
}

func handleRequest(requestID int, dataChannel chan string) {
	rawData := fetchDataFromDatabase(requestID)
	processedData := processData(rawData)
	dataChannel <- processedData
}

func main() {
	requestID := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestID++
		dataChannel := make(chan string)

		go handleRequest(requestID, dataChannel)
		processedData := <-dataChannel
		fmt.Fprintf(w, "Response: %s", processedData)
		fmt.Printf("Response: %s \n", processedData)
	})

	fmt.Println("Server is starting...")
	http.ListenAndServe(":8080", nil)
}
