package tpl

type Tpl[T any, K any] interface {
	Sm(dat T, addr string) (K, error)
	//Rm() (K, error)
}
