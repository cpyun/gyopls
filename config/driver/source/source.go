package source

import "time"

type Source interface {
	Read() (*ChangeSet, error)
	Watch() (Watcher, error)
	String() string
}

// ChangeSet represents a set of changes from a source
type ChangeSet struct {
	Data      []byte
	Checksum  string
	Format    string
	Source    string
	Timestamp time.Time
}

// Watcher watches a source for changes
type Watcher interface {
}