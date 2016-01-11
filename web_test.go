package main

import (
	"testing"
	"time"
	"fmt"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
)

func Test_Webserver_GET(t *testing.T) {

	//create the dummy request
	request, _ := http.NewRequest("GET", "localhost:8080/pipi", nil)
	response := httptest.NewRecorder()

	//execute the request
	processCall(response, request)

	//Check the status code
	checkStatus(t,response,200)
	
	//Check the content
	checkContent(t, response,"Prueba webserver") 
	
	//Check the Header
	checkHeader(t, response, "Go-Test", "Test 1.0")	
		
	fmt.Println("Test_Webserver_GET OK")

}

func Test_Webserver_POST(t *testing.T) {

	//create the dummy request
	request, _ := http.NewRequest("POST", "localhost:8080/pipi", nil)
	response := httptest.NewRecorder()

	//execute the request
	processCall(response, request)

	//Check the status code
	checkStatus(t,response,401)
		
	//Check the Header
	checkHeader(t, response, "Unauthorized", "You need to have a valid token")	
		
	fmt.Println("Test_Webserver_POST OK")

}

//check the body content of a response
func checkContent(t *testing.T, response *httptest.ResponseRecorder, content string) {

	body, _ := ioutil.ReadAll(response.Body)
	if string(body) != content {
		t.Fatalf("Non-expected content %s, expected %s", string(body), content)
	}
}

//Check the status code
func checkStatus(t *testing.T, response *httptest.ResponseRecorder, expected int) {
	if response.Code != expected {
		t.Fatalf("Non-expected status code %v :\n\tbody: %v", expected, response.Code)
	}
}

//Check the specific header value
func checkHeader(t *testing.T, response *httptest.ResponseRecorder, header string, value string) {
	if response.Header().Get(header) != value {
		t.Fatalf("Header: %s, get:%s expected:%s", header, response.Header().Get(header), value)
	}
}

func sleep() {
	time.Sleep(500 * time.Millisecond)
}
