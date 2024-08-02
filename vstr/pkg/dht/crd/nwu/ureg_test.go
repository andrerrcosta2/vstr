package nwu

import (
	"github.com/andrerrcosta2/vstr/nmm/srs/reg"
	"reflect"
	"testing"
	"time"
)

func TestNewUreg(t *testing.T) {
	type args struct {
		cod string
		err error
		dat any
	}
	tests := []struct {
		name string
		args args
		want *Ureg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUreg(tt.args.cod, tt.args.err, tt.args.dat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUreg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUreg_Equal(t *testing.T) {
	type fields struct {
		err error
		cod string
		dat any
		ts  time.Time
	}
	type args struct {
		ot reg.Srrg
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
			u := &Ureg{
				err: tt.fields.err,
				cod: tt.fields.cod,
				dat: tt.fields.dat,
				ts:  tt.fields.ts,
			}
			if got := u.Equal(tt.args.ot); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUreg_Gdt(t *testing.T) {
	type fields struct {
		err error
		cod string
		dat any
		ts  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Ureg{
				err: tt.fields.err,
				cod: tt.fields.cod,
				dat: tt.fields.dat,
				ts:  tt.fields.ts,
			}
			if got := u.Gdt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gdt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUreg_Ger(t *testing.T) {
	type fields struct {
		err error
		cod string
		dat any
		ts  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Ureg{
				err: tt.fields.err,
				cod: tt.fields.cod,
				dat: tt.fields.dat,
				ts:  tt.fields.ts,
			}
			if err := u.Ger(); (err != nil) != tt.wantErr {
				t.Errorf("Ger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUreg_Grc(t *testing.T) {
	type fields struct {
		err error
		cod string
		dat any
		ts  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Ureg{
				err: tt.fields.err,
				cod: tt.fields.cod,
				dat: tt.fields.dat,
				ts:  tt.fields.ts,
			}
			if got := u.Grc(); got != tt.want {
				t.Errorf("Grc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUreg_Gts(t *testing.T) {
	type fields struct {
		err error
		cod string
		dat any
		ts  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Ureg{
				err: tt.fields.err,
				cod: tt.fields.cod,
				dat: tt.fields.dat,
				ts:  tt.fields.ts,
			}
			if got := u.Gts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gts() = %v, want %v", got, tt.want)
			}
		})
	}
}
