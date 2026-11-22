package appvalue_test

import (
	"log"

	"github.com/go-app-blazar/appvalue"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func ExampleUnmarshal() {
	// Create a simple app.Value that represents a user.
	jsValue := app.Window().Get("Object").New()
	jsValue.Set("name", "John")
	jsValue.Set("age", 30)
	jsValue.Set("isStudent", true)
	jsValue.Set("height", 1.75)
	jsValue.Set("emailAddresses", []any{"john1@example.com", "john2@example.com"})
	jsValue.Set("phoneNumbers", []any{map[string]any{"type": "home", "value": "111-555-1000"}, map[string]any{"type": "work", "value": "111-555-2000"}})
	jsValue.Set("createdAt", app.Window().Get("Date").New(app.Window().Get("Date").Call("UTC", 2026, 0, 1, 12, 30, 0, 0)))

	// Define a struct that matches the JavaScript object definition of a user.
	type User struct {
		RawValue       app.Value `javascript:"@this"` // Provide access to the original app.Value in case we need it (for example, for function calls, etc.).
		Name           string    `javascript:"name"`
		Age            int       `javascript:"age"`
		IsStudent      bool      `javascript:"isStudent"`
		Height         float64   `javascript:"height"`
		EmailAddresses []string  `javascript:"emailAddresses"`
		PhoneNumbers   []struct {
			Type  string `javascript:"type"`
			Value string `javascript:"value"`
		} `javascript:"phoneNumbers"`
		CreatedAt appvalue.Timestamp `javascript:"createdAt"`
	}

	// Unmarshal the JavaScript object into a User struct.
	var user User
	err := appvalue.Unmarshal(jsValue, &user)
	if err != nil {
		log.Fatal(err)
	}
}
