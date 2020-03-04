package config

import (
	"context"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type fileSource struct {
	file   string
	format string
	sync.RWMutex
	cancel context.CancelFunc

	// current changeset
	current *Snapshot
}

// Load read initial change set
func (s *fileSource) Load() (*Snapshot, error) {
	s.RLock()
	current := s.current
	s.RUnlock()

	if current != nil {
		return current, nil
	}

	snap, err := s.readFile()
	if err != nil {
		return nil, err
	}

	s.Lock()
	s.current = snap
	s.Unlock()

	return snap, nil
}

// readFile read configuration file and put into current snapshot
func (s *fileSource) readFile() (*Snapshot, error) {
	b, err := ioutil.ReadFile(s.file)
	if err != nil {
		return nil, err
	}

	snap := &Snapshot{Data: b}

	return snap, nil
}

func (s *fileSource) watchChanges() {
	// create context and cancelation
	ctx, cancel := context.WithCancel(context.Background())

	s.cancel = cancel
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	watcher.Add(s.file)

	for {
		select {
		case <-ctx.Done():
			watcher.Close()
			return
		case event, ok := <-watcher.Events:
			if !ok {
				break
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				snap, err := s.readFile()
				if err != nil {
					log.Fatal(err)
				}

				s.Lock()
				if s.current.Checksum() != snap.Checksum() {
					s.current = snap
				}
				s.Unlock()
			}
		}
	}
}

// File create config source from give file
func File(file string, watch ...bool) Loader {
	s := &fileSource{
		file:   file,
		format: filepath.Ext(file)[1:],
	}

	if len(watch) > 0 && watch[0] {
		go s.watchChanges()
	}

	return s
}
