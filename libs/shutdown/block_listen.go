package shutdown

import "context"

func BlockListen(ctx context.Context, fn func() error) error {
	lisErr := make(chan error, 1)
	go func() {
		if e := fn(); e != nil {
			lisErr <- e
		} else {
			close(lisErr)
		}
	}()
	for {
		select {
		case err := <-lisErr:
			return err
		case <-ctx.Done():
			return nil
		}
	}
}
