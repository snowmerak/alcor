package wsstation

import "sync"

type Signal struct {
	state int8
}

type Sessions struct {
	signals map[chan<- Signal]struct{}
	mutex   sync.RWMutex
}

func (s *Sessions) Add(ch chan<- Signal) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.signals[ch] = struct{}{}
	return nil
}

func (s *Sessions) Delete(ch chan<- Signal) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.signals, ch)
	return nil
}

func (s *Sessions) BroadCast(signal Signal) error {
	for ch := range s.signals {
		go func(ch chan<- Signal) {
			ch <- signal
		}(ch)
	}
	return nil
}
