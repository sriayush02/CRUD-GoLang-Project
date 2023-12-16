// main_test.go

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCart(t *testing.T) {
	go main()
	req, _ := http.NewRequest("POST", "http://localhost:9000/cart/add/TestProduct/19.99", nil)
	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		t.Error("Could not get response", err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Expected status code 201")

}

func TestViewCart(t *testing.T) {
	go main()
	req, _ := http.NewRequest("GET", "http://localhost:9000/cart/view", nil)
	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		t.Error("Could not get response", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

}

func TestDeleteProduct(t *testing.T) {
	go main()
	req, _ := http.NewRequest("DELETE", "http://localhost:9000/cart/delete/2", nil)
	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		t.Error("Could not get response", err)
	}

	assert.Equal(t, http.StatusNoContent, resp.StatusCode, "Expected status code 204")

}

func TestUpdateCart(t *testing.T) {
	go main()
	testProductname := "Pant"
	testAmount := "399.99"
	reqBody := map[string]interface{}{"name": testProductname, "amount": testAmount}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("PUT", "http://localhost:9000/cart/update/4", bytes.NewBuffer(body))
	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		t.Error("Could not get response", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

}

// func TestViewCart(t *testing.T) {
// 	app := gofr.New()

// 	// Set up a test server
// 	ts := httptest.NewServer(app)
// 	defer ts.Close()

// 	// Perform a GET request
// 	resp, err := http.Get(ts.URL + "/cart/view")
// 	assert.NoError(t, err)
// 	defer resp.Body.Close()

// 	// Check if the request was successful
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	// Parse the response
// 	var result []Cart
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	assert.NoError(t, err)

// 	// Check the response
// 	assert.NotEmpty(t, result)
// }

// func TestDeleteProduct(t *testing.T) {
// 	app := gofr.New()

// 	// Set up a test server
// 	ts := httptest.NewServer(app)
// 	defer ts.Close()

// 	// Test data
// 	productID := 1

// 	// Perform a DELETE request
// 	req, err := http.NewRequest("DELETE", ts.URL+"/cart/delete/"+string(rune(productID)), nil)
// 	assert.NoError(t, err)

// 	client := http.Client{}
// 	resp, err := client.Do(req)
// 	assert.NoError(t, err)
// 	defer resp.Body.Close()

// 	// Check if the request was successful
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	// Parse the response
// 	var result map[string]string
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	assert.NoError(t, err)

// 	// Check the response
// 	assert.Equal(t, "Deleted Successfully", result["err"])
// }

// func TestUpdateCart(t *testing.T) {
// 	app := gofr.New()

// 	// Set up a test server
// 	ts := httptest.NewServer(app)
// 	defer ts.Close()

// 	// Test data
// 	productID := 1
// 	payload := `{"name": "UpdatedProduct", "amount": 29.99}`

// 	// Perform a PUT request
// 	req, err := http.NewRequest("PUT", ts.URL+"/cart/update/"+string(rune(productID)), strings.NewReader(payload))
// 	assert.NoError(t, err)
// 	req.Header.Set("Content-Type", "application/json")

// 	client := http.Client{}
// 	resp, err := client.Do(req)
// 	assert.NoError(t, err)
// 	defer resp.Body.Close()

// 	// Check if the request was successful
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	// Parse the response
// 	var result map[string]string
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	assert.NoError(t, err)

// 	// Check the response
// 	assert.Equal(t, "Updated Successfully", result["err"])
// }
