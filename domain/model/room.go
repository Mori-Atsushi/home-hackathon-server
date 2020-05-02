package model

import (
	"sync"
)

type Room struct {
	mutex    *sync.Mutex
	channels map[string]*Channel
}

func NewRoom() Room {
	return Room{
		mutex:    &sync.Mutex{},
		channels: map[string]*Channel{},
	}
}

func (r Room) AddChannel(user User) {
	channel := NewChannel()
	r.mutex.Lock()
	r.channels[user.id] = channel
	r.mutex.Unlock()
}

func (r Room) RemoveChannel(user User) {
	r.mutex.Lock()
	delete(r.channels, user.id)
	r.mutex.Unlock()
}