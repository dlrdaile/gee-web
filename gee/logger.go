package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		// record the request info and route
		log.Printf("Started %s %s", c.Req.Method, c.Req.RequestURI)
		t := time.Now()
		// Process request
		c.Next()

		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
