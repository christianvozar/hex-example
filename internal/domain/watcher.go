package domain

import "context"

type Watcher interface {
	Watch(ctx context.Context) error
}
