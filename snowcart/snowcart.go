package snowcart

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Snowcart struct {
	Url           string
	MaxQueueDepth int
	Timeout       time.Duration
	eventQueue    EventQueue
}

func New(url string, maxQueueDepth int, timeout time.Duration) *Snowcart {

	s := &Snowcart{
		Url:           url,
		MaxQueueDepth: maxQueueDepth,
		Timeout:       timeout,
		eventQueue:    EventQueue{},
	}

	go s.drainAfterInterval()
	return s
}

func (s *Snowcart) Emit(e *Event) {

	s.eventQueue.Append(e)
	length := s.eventQueue.Length()

	if length == s.MaxQueueDepth {
		s.drain()
	}
}

func (s *Snowcart) drain() {

	eventsBytes, err := json.Marshal(s.eventQueue.GetCurrentEventsList())
	if err != nil {
		log.Printf("Could not send events to Snowcart %s", err)
	}

	fmt.Printf("Sending events %s\n", eventsBytes)
	s.eventQueue.Reset()
}

func (s *Snowcart) Close() {
	s.drain()
}

func (s *Snowcart) drainAfterInterval() {
	for {
		time.Sleep(s.Timeout)
		s.drain()
	}
}
