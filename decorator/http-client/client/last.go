package client

import (
	"log"
	"net/http"
)

// Audit will create a client decorator with auditing concerns.
func Last(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.Println("last")
			return c.Do(r)
		})
	}
}
