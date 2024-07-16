package reg

import "time"

type Srrg interface {
	Gts() time.Time
	Grc() string
	Ger() error
	Gdt() any
	Equal(t Srrg) bool
}
