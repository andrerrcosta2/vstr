package nwu

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewUrstk(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name string
		args args
		want *Urstk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUrstk(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUrstk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrstk_Add(t *testing.T) {
	type fields struct {
		mtx sync.Mutex
		reg map[string][]*Ureg
	}
	type args struct {
		code string
		err  error
		data any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Urstk{
				mtx: tt.fields.mtx,
				reg: tt.fields.reg,
			}
			u.Add(tt.args.code, tt.args.err, tt.args.data)
		})
	}
}

func TestUrstk_Ctn(t *testing.T) {
	type fields struct {
		mtx sync.Mutex
		reg map[string][]*Ureg
	}
	type args struct {
		cod string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Urstk{
				mtx: tt.fields.mtx,
				reg: tt.fields.reg,
			}
			if got := u.Ctn(tt.args.cod); got != tt.want {
				t.Errorf("Ctn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrstk_Get(t *testing.T) {
	type fields struct {
		mtx sync.Mutex
		reg map[string][]*Ureg
	}
	type args struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Ureg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Urstk{
				mtx: tt.fields.mtx,
				reg: tt.fields.reg,
			}
			if got := u.Get(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrstk_Mpt(t *testing.T) {
	type fields struct {
		mtx sync.Mutex
		reg map[string][]*Ureg
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Urstk{
				mtx: tt.fields.mtx,
				reg: tt.fields.reg,
			}
			if got := u.Mpt(); got != tt.want {
				t.Errorf("Mpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrstk_RK(t *testing.T) {
	type fields struct {
		mtx sync.Mutex
		reg map[string][]*Ureg
	}
	type args struct {
		cod string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Urstk{
				mtx: tt.fields.mtx,
				reg: tt.fields.reg,
			}
			u.RK(tt.args.cod)
		})
	}
}

func TestUrstk_RR(t *testing.T) {
	type fields struct {
		mtx sync.Mutex
		reg map[string][]*Ureg
	}
	type args struct {
		reg *Ureg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Urstk{
				mtx: tt.fields.mtx,
				reg: tt.fields.reg,
			}
			u.RR(tt.args.reg)
		})
	}
}

func TestUrstk_Sz(t *testing.T) {
	type fields struct {
		mtx sync.Mutex
		reg map[string][]*Ureg
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Urstk{
				mtx: tt.fields.mtx,
				reg: tt.fields.reg,
			}
			if got := u.Sz(); got != tt.want {
				t.Errorf("Sz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrstk_Xtr(t *testing.T) {
	type fields struct {
		mtx sync.Mutex
		reg map[string][]*Ureg
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Ureg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Urstk{
				mtx: tt.fields.mtx,
				reg: tt.fields.reg,
			}
			if got := u.Xtr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Xtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
