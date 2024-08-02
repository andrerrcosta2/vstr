package crd

import (
	"github.com/andrerrcosta2/vstr/dht/cmn/dhtcmx"
	"github.com/andrerrcosta2/vstr/dht/crd/cid"
	"github.com/andrerrcosta2/vstr/dht/crd/lut"
	"github.com/andrerrcosta2/vstr/dht/crd/nod"
	"github.com/andrerrcosta2/vstr/dht/crd/nwu"
	"github.com/andrerrcosta2/vstr/dht/dhtcfg"
	"github.com/andrerrcosta2/vstr/nmm/hchk"
	"github.com/andrerrcosta2/vstr/nmm/srcfg"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg dhtcfg.CrdCfg
		msr dhtcmx.Nmr
	}
	tests := []struct {
		name string
		args args
		want *crdv
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg, tt.args.msr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crdv_Hbt(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			c.Hbt()
		})
	}
}

func Test_crdv_Jn(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
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
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			if err := c.Jn(); (err != nil) != tt.wantErr {
				t.Errorf("Jn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_crdv_Rcv(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		s srcfg.Srcfg
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
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			c.Rcv(tt.args.s)
		})
	}
}

func Test_crdv_Rt(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			got, got1 := c.Rt(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Rt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_crdv_St(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		key   string
		value []byte
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
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			c.St(tt.args.key, tt.args.value)
		})
	}
}

func Test_crdv_bsp(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	tests := []struct {
		name    string
		fields  fields
		want    *nod.Nod
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			got, err := c.bsp()
			if (err != nil) != tt.wantErr {
				t.Errorf("bsp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bsp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crdv_dpnu(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		s srcfg.Srcfg
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
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			c.dpnu(tt.args.s)
		})
	}
}

func Test_crdv_dpsu(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		s srcfg.Srcfg
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
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			c.dpsu(tt.args.s)
		})
	}
}

func Test_crdv_fp(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		id cid.Id
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nod.Nod
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			got, err := c.fp(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("fp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crdv_fs(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		id cid.Id
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nod.Nod
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			got, err := c.fs(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("fs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crdv_fwp(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		id cid.Id
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nod.Nod
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			got, err := c.fwp(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("fwp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fwp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crdv_fws(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		id cid.Id
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nod.Nod
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			got, err := c.fws(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("fws() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fws() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crdv_ife(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		i uint
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
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			c.ife(tt.args.i)
		})
	}
}

func Test_crdv_ift(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			c.ift()
		})
	}
}

func Test_crdv_ufe(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	type args struct {
		f *nod.Fge
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
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			c.ufe(tt.args.f)
		})
	}
}

func Test_crdv_uft(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			c.uft()
		})
	}
}

func Test_crdv_unet(t *testing.T) {
	type fields struct {
		cfg  dhtcfg.CrdCfg
		nod  nod.Nod
		msr  dhtcmx.Nmr
		lut  lut.LUT
		urgs *nwu.Urstk
		hsts *hchk.Hchks[*nwu.Ureg]
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crdv{
				cfg:  tt.fields.cfg,
				nod:  tt.fields.nod,
				msr:  tt.fields.msr,
				lut:  tt.fields.lut,
				urgs: tt.fields.urgs,
				hsts: tt.fields.hsts,
			}
			c.unet()
		})
	}
}
