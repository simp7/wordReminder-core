package uid

import (
	"sync"
	"wordReminder-core/user"
)

var instance *creator
var once sync.Once

type creator struct {
	current user.UID
	dataChan chan user.UID
	lock sync.Mutex
}

func New(last user.UID) user.UID {

	once.Do(func() {
		initInstance(last)
	})

	instance.increment()

	return <- instance.dataChan

}

func initInstance(last user.UID) {
	instance = new(creator)
	instance.current = last + 1
	instance.dataChan = make(chan user.UID, 1)
}

func (c *creator) increment() {
	instance.lock.Lock()
	instance.dataChan <- instance.current
	instance.current ++
	instance.lock.Unlock()
}