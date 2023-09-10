package counter

import (
	"context"
	"log"
	"sync"

	"github.com/BioSphereGame/server-core/proto/counter"
)

type Counter struct {
	counter.UnimplementedCounterServer
	counterValue int32
	mu           sync.Mutex
}

func (s *Counter) Get(ctx context.Context, _ *counter.Empty) (*counter.Value, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Received 'Get', counter: %d\n", s.counterValue)
	return &counter.Value{Value: s.counterValue}, nil
}

func (s *Counter) Increment(ctx context.Context, _ *counter.Empty) (*counter.Value, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Received 'Increment', counter: %d\n", s.counterValue)
	s.counterValue++
	return &counter.Value{Value: s.counterValue}, nil
}

func (s *Counter) Decrement(ctx context.Context, _ *counter.Empty) (*counter.Value, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Received 'Decrement', counter: %d\n", s.counterValue)
	s.counterValue--
	return &counter.Value{Value: s.counterValue}, nil
}
