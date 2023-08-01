package main

import (
	"sync"
	"time"
)

type Publisher struct {
	subscribers map[Subscriber]TopicFunc
	buffer      int
	timeout     time.Duration
	m           sync.RWMutex
}
type (
	Subscriber chan interface{}
	TopicFunc  func(v interface{}) bool
)

func (p *Publisher) Subscribe() Subscriber {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic TopicFunc) Subscriber {
	ch := make(Subscriber, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()

	return ch
}
func (p *Publisher) Delete(sub Subscriber) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.Unlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func (p *Publisher) sendTopic(sub Subscriber, topic TopicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

//算了，没写完
