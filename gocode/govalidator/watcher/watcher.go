package watcher

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

// Watcher .
type Watcher struct {
	Event chan uint32
	Fd    int
}

// NewWatcher .
func NewWatcher() (*Watcher, error) {
	fd, err := unix.InotifyInit()
	if err != nil {
		return nil, err
	}
	watcher := &Watcher{
		Event: make(chan uint32),
		Fd:    fd,
	}
	watcher.getEvents()
	return watcher, nil
}

// AddWatcher .
func (w *Watcher) AddWatcher(file string, mask uint32) error {
	_, err := unix.InotifyAddWatch(w.Fd, file, mask)
	if err != nil {
		return err
	}
	return nil
}

func (w *Watcher) getEvents() {
	go func() {
		var buf [unix.SizeofInotifyEvent * 4096]byte
		for {
			n, err := unix.Read(w.fd, buf[:])
			if err != nil {
				n = 0
				continue
			}
			var offset uint32
			for offset <= uint32(n-unix.SizeofInotifyEvent) {
				raw := (*unix.InotifyEvent)(unsafe.Pointer(&buf[offset]))
				mask := uint32(raw.Mask)
				nameLen := uint32(raw.Len)
				// 塞到事件队列
				w.Event <- mask
				offset += unix.SizeofInotifyEvent + nameLen
			}
		}
	}()
}
