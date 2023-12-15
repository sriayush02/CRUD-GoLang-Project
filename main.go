package main

import (
	"encoding/json"

	"gofr.dev/pkg/gofr"
)

type Cart struct {
	PID     int     `json:"id"`
	PName   string  `json:"name"`
	PAmount float64 `json:"amount"`
}

func main() {
	// initialise gofr object
	app := gofr.New()

	app.GET("/greet", greet)

	app.POST("/cart/add/{name}/{amount}", createcart)

	app.GET("/cart/view", viewcart)

	app.DELETE("/cart/delete/{id}", deleteproduct)

	app.PUT("/cart/update/{id}", updatecart)

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
func greet(ctx *gofr.Context) (interface{}, error) {
	// Get the value using the redis instance
	value, err := ctx.Redis.Get(ctx.Context, "greeting").Result()

	return value, err
}
func createcart(ctx *gofr.Context) (interface{}, error) {
	name := ctx.PathParam("name")
	amount := ctx.PathParam("amount")

	// Inserting a customer row in database using SQL
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO carts (name, amount) VALUES (?, ?)", name, amount)

	return nil, err
}
func viewcart(ctx *gofr.Context) (interface{}, error) {
	var carts []Cart

	// Getting the customer from the database using SQL
	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM carts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cart Cart
		if err := rows.Scan(&cart.PID, &cart.PName, &cart.PAmount); err != nil {
			return nil, err
		}

		carts = append(carts, cart)
	}

	// return the customer
	return carts, nil
}

func deleteproduct(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	// Inserting a customer row in database using SQL
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM carts WHERE id =?", id)

	return nil, err

}
func updatecart(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	var newcart struct {
		PName   string  `json:"name"`
		PAmount float64 `json:"amount"`
	}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&newcart); err != nil {
		return nil, err
	}
	if newcart.PName == "" {
		return map[string]string{"ERR": "empty string"}, nil
	}

	// Inserting a customer row in database using SQL
	_, err := ctx.DB().ExecContext(ctx, "UPDATE carts SET name = ?, amount = ? WHERE id = ?", newcart.PName, newcart.PAmount, id)

	return nil, err

}
