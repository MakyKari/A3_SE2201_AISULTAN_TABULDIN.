package greenlight_test

import (
	"testing"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)


func printResponseBody(resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Print(sb)
}

func printRequestBody(req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Print(sb)

}

// Unit tests for the API
func TestNewAccountCreationRequest(t *testing.T){
	postBody, _ := json.Marshal(map[string]string{
		"Name":  "Ais",
		"Email": "tab@tab.a",
		"Password": "password123",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/users", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestNewAccountCreationRequest2(t *testing.T){
	postBody, _ := json.Marshal(map[string]string{
		"Name":  "Ais",
		"Email": "something not an email",
		"Password": "",})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/users", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestGettingAccountAuthenticationToken(t *testing.T){
	postBody, _ := json.Marshal(map[string]string{
		"Email": "tab@tab.a",
		"Password": "password123",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/tokens/authentication", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestGettingAccountAuthenticationToken2(t *testing.T){
	postBody, _ := json.Marshal(map[string]string{
		"Email": "tabtabtab@tab.a", // This email does not exist
		"Password": "password123",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/tokens/authentication", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

//FOR THE LAST 2 TESTS TO WORK, THE ACCOUNT MUST BE CREATED FIRST
//TO DO THIS, RUN THE FIRST TEST FIRST WITH YOUR EMAIL AND PASSWORD
//CHECK YOUR EMAIL FOR THE ACTIVATION TOKEN

func TestActivatingAccount(t *testing.T){
	postBody, _ := json.Marshal(map[string]string{
		"Token": "MQJNOQ3BTN6MHBENQBRROMA2FY", // ACTIVATION TOKEN
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.NewRequest(http.MethodPost,"http://localhost:4000/v1/users/activated", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printRequestBody(resp)
}

func TestActivatingAccount2(t *testing.T){
	postBody, _ := json.Marshal(map[string]string{
		"Token": "NOT A VALID TOKEN", // ACTIVATION TOKEN
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.NewRequest(http.MethodPost,"http://localhost:4000/v1/users/activated", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printRequestBody(resp)
}

//Integration tests for the API

type Movie struct {
	Title    string   `json:"title"`
	Year     int      `json:"year"`
	Runtime  string   `json:"runtime"`
	Genres   []string `json:"genres"`
}

func TestInsertingMoviesIntoDatabase(t *testing.T) {
	moviePayload := Movie{
		Title:   "Inception",
		Year:    2010,
		Runtime: "144 mins",
		Genres:  []string{"thriller"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "YXJXRFN44TZTZJ4OES3BVCR2RQ"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestInsertingMoviesIntoDatabaseWithWrongYear(t *testing.T) {
	moviePayload := Movie{
		Title:   "Inception",
		Year:    2029, // This is a future date
		Runtime: "144 mins",
		Genres:  []string{"thriller"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "YXJXRFN44TZTZJ4OES3BVCR2RQ"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestInsertingMoviesIntoDatabaseWithWrongRuntime(t *testing.T) {
	moviePayload := Movie{
		Title:   "Inception",
		Year:    2020, // This is a future date
		Runtime: "144",
		Genres:  []string{"thriller"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "YXJXRFN44TZTZJ4OES3BVCR2RQ"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestMovieDeletionById(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:4000/v1/movies/3", nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "YXJXRFN44TZTZJ4OES3BVCR2RQ"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()
	
	printResponseBody(resp)
}