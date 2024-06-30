// Andre R. R. Costa *** github.com/andrerrcosta2

package nod

import (
	"crypto/sha1"
)

const (
	M = 160
)

type Id [M / 8]byte

func Nid(key string) Id {
	return Id(sha1.Sum([]byte(key)))
}

type Fge struct {
	str Id
	Nod *Nod
}

type Nod struct {
	ID  Id
	IP  string
	Pt  int
	Fgs []Fge
	Pre *Nod
	Suc *Nod
}

// Jn jn
func (n *Nod) Jn(exn *Nod) {
	if exn != nil {
		n.Suc = exn.fs(n.ID)
	} else {
		// Setup as a ring
		n.Suc = n
		n.Pre = n
	}
	n.ift(exn)
	n.unet()
}

// Stb Stab
func (n *Nod) Stb() {
	x := n.Suc.Pre
	if x != nil && btw(n.ID, x.ID, n.Suc.ID) {
		n.Suc = x
	}
	n.Suc.nfy(n)
}

// Sucf Suc
func (n *Nod) Sucf(i Id) *Nod {
	if btw(n.ID, i, n.Suc.ID) {
		return n.Suc
	} else {
		cpn := n.cpn(i)
		return cpn.fs(i)
	}
}

// Cpn cpn
func (n *Nod) cpn(i Id) *Nod {
	for k := M - 1; k >= 0; k-- {
		if btw(n.ID, n.Fgs[k].Nod.ID, i) {
			return n.Fgs[k].Nod
		}
	}
	return n
}

// fs fds
func (n *Nod) fs(i Id) *Nod {
	return nil
}

// ift ift
func (n *Nod) ift(t *Nod) {
	//t.Fgs[]
}

// nfy nfy
func (n *Nod) nfy(t *Nod) {

}

// unet unry
func (n *Nod) unet() {

}

func btw(pre, nod, suc Id) bool {
	return false
}
