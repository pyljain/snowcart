package snowcart

import "sync"

type EventQueue struct {
	queue []*Event
	lock  sync.RWMutex
}

func (eq *EventQueue) Append(event *Event) {
	eq.lock.Lock()
	defer eq.lock.Unlock()

	eq.queue = append(eq.queue, event)
}

func (eq *EventQueue) Reset() {
	eq.lock.Lock()
	defer eq.lock.Unlock()

	eq.queue = []*Event{}
}

func (eq *EventQueue) GetCurrentEventsList() []*Event {
	eq.lock.RLock()
	defer eq.lock.RUnlock()
	currentEvents := []*Event{}
	currentEvents = append(currentEvents, eq.queue...)

	return currentEvents
}

func (eq *EventQueue) Length() int {
	eq.lock.RLock()
	defer eq.lock.RUnlock()
	return len(eq.queue)
}
