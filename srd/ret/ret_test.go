package ret

import (
	"errors"
	"testing"
	"time"
)

func TestRt(t *testing.T) {
	tcs := []struct {
		name    string
		att     uint
		dl      time.Duration
		fn      func() (int, error)
		want    int
		wantErr bool
	}{
		{
			name:    "Success on first attempt",
			att:     3,
			dl:      time.Millisecond,
			fn:      func() (int, error) { return 1, nil },
			want:    1,
			wantErr: false,
		},
		{
			name:    "Fail all attempts",
			att:     3,
			dl:      time.Millisecond,
			fn:      func() (int, error) { return 0, errors.New("fail") },
			want:    0,
			wantErr: true,
		},
		{
			name: "Success on second attempt",
			att:  3,
			dl:   time.Millisecond,
			fn: func() func() (int, error) {
				attempt := 0
				return func() (int, error) {
					attempt++
					if attempt == 2 {
						return 1, nil
					}
					return 0, errors.New("fail")
				}
			}(),
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Rt(tt.att, tt.dl, tt.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRtf(t *testing.T) {
	tcs := []struct {
		name    string
		att     uint
		dl      time.Duration
		fcr     float64
		fn      func() (int, error)
		want    int
		wantErr bool
	}{
		{
			name:    "\nSuccess on first attempt",
			att:     3,
			dl:      time.Millisecond,
			fcr:     1.5,
			fn:      func() (int, error) { return 1, nil },
			want:    1,
			wantErr: false,
		},
		{
			name:    "\nFail all attempts",
			att:     3,
			dl:      time.Millisecond,
			fcr:     1.5,
			fn:      func() (int, error) { return 0, errors.New("fail") },
			want:    0,
			wantErr: true,
		},
		{
			name: "\nSuccess on second attempt",
			att:  3,
			dl:   time.Millisecond,
			fcr:  1.5,
			fn: func() func() (int, error) {
				attempt := 0
				return func() (int, error) {
					attempt++
					if attempt == 2 {
						return 1, nil
					}
					return 0, errors.New("fail")
				}
			}(),
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Rtf(tt.att, tt.dl, tt.fcr, tt.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rtf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rtf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRtSq(t *testing.T) {
	prms := []int{1, 2, 3}
	tcs := []struct {
		name    string
		att     uint
		dl      time.Duration
		prms    []int
		fn      func(int) (int, error)
		want    int
		wantErr bool
	}{
		{
			name:    "Success on first param",
			att:     3,
			dl:      time.Millisecond,
			prms:    prms,
			fn:      func(n int) (int, error) { return n, nil },
			want:    1,
			wantErr: false,
		},
		{
			name:    "Fail all attempts",
			att:     3,
			dl:      time.Millisecond,
			prms:    prms,
			fn:      func(n int) (int, error) { return 0, errors.New("fail") },
			want:    0,
			wantErr: true,
		},
		{
			name: "Success on second param, first attempt",
			att:  3,
			dl:   time.Millisecond,
			prms: prms,
			fn: func(n int) (int, error) {
				if n == 2 {
					return n, nil
				}
				return 0, errors.New("fail")
			},
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RtSq(tt.att, tt.dl, tt.prms, tt.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("RtSq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RtSq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRtRd(t *testing.T) {
	prms := []int{1, 2, 3}
	tcs := []struct {
		name    string
		att     uint
		dl      time.Duration
		prms    []int
		fn      func(int) (int, error)
		want    int
		wantErr bool
	}{
		{
			name:    "Success on first param",
			att:     3,
			dl:      time.Millisecond,
			prms:    prms,
			fn:      func(n int) (int, error) { return n, nil },
			want:    1,
			wantErr: false,
		},
		{
			name:    "Fail all attempts",
			att:     3,
			dl:      time.Millisecond,
			prms:    prms,
			fn:      func(n int) (int, error) { return 0, errors.New("fail") },
			want:    0,
			wantErr: true,
		},
		{
			name: "Success on second param, first attempt",
			att:  3,
			dl:   time.Millisecond,
			prms: prms,
			fn: func(n int) (int, error) {
				if n == 2 {
					return n, nil
				}
				return 0, errors.New("fail")
			},
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RtRd(tt.att, tt.dl, tt.prms, tt.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("RtRd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RtRd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRtSh(t *testing.T) {
	prms := []int{1, 2, 3}
	tcs := []struct {
		name    string
		att     uint
		dl      time.Duration
		prms    []int
		fn      func(int) (int, error)
		want    int
		wantErr bool
	}{
		{
			name:    "Success on first param",
			att:     3,
			dl:      time.Millisecond,
			prms:    prms,
			fn:      func(n int) (int, error) { return n, nil },
			want:    1,
			wantErr: false,
		},
		{
			name:    "Fail all attempts",
			att:     3,
			dl:      time.Millisecond,
			prms:    prms,
			fn:      func(n int) (int, error) { return 0, errors.New("fail") },
			want:    0,
			wantErr: true,
		},
		{
			name: "Success on second param, first attempt",
			att:  3,
			dl:   time.Millisecond,
			prms: prms,
			fn: func(n int) (int, error) {
				if n == 2 {
					return n, nil
				}
				return 0, errors.New("fail")
			},
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RtSh(tt.att, tt.dl, tt.prms, tt.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("RtSh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RtSh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRtSt(t *testing.T) {
	prms := []int{3, 1, 2}
	tcs := []struct {
		name    string
		att     uint
		dl      time.Duration
		prms    []int
		stf     func(a, b int) bool
		fn      func(int) (int, error)
		want    int
		wantErr bool
	}{
		{
			name:    "Success on sorted param",
			att:     3,
			dl:      time.Millisecond,
			prms:    prms,
			stf:     func(a, b int) bool { return a < b },
			fn:      func(n int) (int, error) { return n, nil },
			want:    1,
			wantErr: false,
		},
		{
			name:    "Fail all attempts",
			att:     3,
			dl:      time.Millisecond,
			prms:    prms,
			stf:     func(a, b int) bool { return a < b },
			fn:      func(n int) (int, error) { return 0, errors.New("fail") },
			want:    0,
			wantErr: true,
		},
		{
			name: "Success on second param, first attempt after sort",
			att:  3,
			dl:   time.Millisecond,
			prms: prms,
			stf:  func(a, b int) bool { return a < b },
			fn: func(n int) (int, error) {
				if n == 2 {
					return n, nil
				}
				return 0, errors.New("fail")
			},
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RtSt(tt.att, tt.dl, tt.prms, tt.stf, tt.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("RtSt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RtSt() = %v, want %v", got, tt.want)
			}
		})
	}
}
