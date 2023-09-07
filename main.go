package main

import (
	"snowcart/snowcart"
	"time"

	"github.com/google/uuid"
)

func main() {
	s := snowcart.New("http://localhost:9090", 5, 3*time.Second)
	defer s.Close()

	for i := 0; i <= 10; i++ {
		e := snowcart.Event{
			Name:      "number_of_clicks",
			Namespace: "myapp",
			Id:        uuid.NewString(),
			Timestamp: time.Now().UnixMilli(),
			Value:     i,
		}

		time.Sleep(1 * time.Second)
		s.Emit(&e)
	}

}
