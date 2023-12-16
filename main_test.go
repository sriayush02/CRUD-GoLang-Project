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

	
	testProductName := "Pant"
	testAmount := 399.99

	
	reqBody := map[string]interface{}{"name": testProductName, "amount": testAmount}
	body, _ := json.Marshal(reqBody)

	// Create and send PUT request
	req, err := http.NewRequest("PUT", "http://localhost:9000/cart/update/4", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}

	
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Could not get response:", err)
	}

	
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
}
