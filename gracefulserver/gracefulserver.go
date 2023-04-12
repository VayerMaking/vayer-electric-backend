package gracefulserver

import (
	"context"
	"log"
	"net/http"

	"vayer-electric-backend/constants"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type GracefulServer struct {
	server  *http.Server
	cancel  context.CancelFunc
	wait    func() error
	running bool
}

// New graceful server wrapper
func New(srv *http.Server) *GracefulServer {
	return &GracefulServer{
		server: srv,
	}
}

func (s *GracefulServer) StartListening(ctx context.Context) error {
	if s.server == nil {
		return errors.Errorf("http.Server required")
	}

	ctx, cancel := context.WithCancel(ctx)
	g, gCtx := errgroup.WithContext(ctx)

	s.wait = g.Wait
	s.cancel = cancel

	g.Go(func() error {
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
		return nil
	})
	g.Go(func() error {
		<-gCtx.Done()

		// Shutdown signal with grace period of constants.ShutdownTimeout seconds
		timeout, _ := context.WithTimeout(context.Background(), constants.ShutdownTimeout)
		go func() {
			<-timeout.Done()
			if timeout.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		return s.server.Shutdown(context.Background())
	})

	s.running = true

	return nil
}

func (s *GracefulServer) Shutdown() error {
	if !s.running {
		return nil
	}
	s.cancel()
	return s.wait()
}
