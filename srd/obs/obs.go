package obs

type Obs[T any] interface {
	Sub(obv Obv[T]) (*Sub[T], error)
	Nxt(d T)
	Rmo(s *Sub[T])
}
