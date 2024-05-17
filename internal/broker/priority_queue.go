package broker

import "sync"

type message struct {
	content  string
	priority int32
}

type priorityQueue struct {
	high   []message
	medium []message
	low    []message
	sync.Mutex
}

func (pq *priorityQueue) Enqueue(msg message) {
	pq.Lock()
	defer pq.Unlock()
	switch msg.priority {
	case 1:
		pq.high = append(pq.high, msg)
	case 2:
		pq.medium = append(pq.medium, msg)
	default:
		pq.low = append(pq.low, msg)
	}
}

func (pq *priorityQueue) Dequeue() *message {
	pq.Lock()
	defer pq.Unlock()
	if len(pq.high) > 0 {
		msg := pq.high[0]
		pq.high = pq.high[1:]
		return &msg
	} else if len(pq.medium) > 0 {
		msg := pq.medium[0]
		pq.medium = pq.medium[1:]
		return &msg
	} else if len(pq.low) > 0 {
		msg := pq.low[0]
		pq.low = pq.low[1:]
		return &msg
	}
	return nil
}
