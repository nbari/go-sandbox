package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kr/beanstalk"
)

var conn, _ = beanstalk.Dial("tcp", "127.0.0.1:11300")

func main() {
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err == nil {
		go func() {
			for i := 0; i < 10; i++ {

				_, err := c.Put([]byte("hello"), 1, 0, 120*time.Second)
				if err != nil {
					log.Fatal(err)
				}
			}
		}()
	}
	c, err = beanstalk.Dial("tcp", "127.0.0.1:11300")
	for {
		id, body, err := c.Reserve(5 * time.Second)
		if err == nil {
			fmt.Printf("id: %d body: %s\n", id, body)
			c.Delete(id)
		}
	}
}
