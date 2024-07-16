package obs

type Obv[T any] struct {
	Nxt func(T)
	Err func(error)
	Cpt func()
}

func NewObv[T any](nxt func(T), err func(error), cpt func()) *Obv[T] {
	return &Obv[T]{
		Nxt: nxt,
		Err: err,
		Cpt: cpt,
	}
}
