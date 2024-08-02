package obs

import (
	"github.com/google/uuid"
)

type Sub[T any] struct {
	id  string
	obs Obs[T]
}

func NewSub[T any](obs Obs[T]) *Sub[T] {
	return &Sub[T]{
		id:  uuid.New().String(),
		obs: obs,
	}
}

func (s *Sub[T]) Gid() string {
	return s.id
}

func (s *Sub[T]) Uns() {
	s.obs.Rmo(s)
}
