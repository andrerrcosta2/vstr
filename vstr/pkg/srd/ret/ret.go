package ret

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Rt[T any](att uint, dl time.Duration, fn func() (T, error)) (T, error) {
	var mpt T

	var err error
	var res T

	for i := 1; i <= int(att); i++ {
		res, err = fn()
		if err == nil {
			return res, nil
		}
		fmt.Printf("attmpt %d failed: %s\n", i, err)
		time.Sleep(dl)
	}
	return mpt, fmt.Errorf("all %d attpts has failed\n", att)
}

func Rtf[T any](att uint, dl time.Duration, fcr float64, fn func() (T, error)) (T, error) {
	var mpt T

	var err error
	var res T

	for i := 1; i <= int(att); i++ {
		res, err = fn()
		if err == nil {
			return res, nil
		}
		fmt.Printf("attmpt %d failed: %s\n", i, err)
		time.Sleep(dl)
		dl = time.Duration(float64(dl) * fcr)
	}
	return mpt, fmt.Errorf("all %d attpts has failed\n", att)
}

func RtSq[T any, K any](att uint, dl time.Duration, prms []T, fn func(T) (K, error)) (K, error) {
	var err error
	var res K
	if att <= 0 || len(prms) == 0 {
		return res, errors.New("invalid number of attempts or empty params list\n")
	}

	for i := 0; i < int(att); i++ {
		for j, prm := range prms {
			res, err = fn(prm)
			if err == nil {
				return res, nil
			}
			fmt.Printf("attempt %d with param %d failed: %s\n", i+1, j+1, err)
			time.Sleep(dl)
		}
	}
	return res, fmt.Errorf("all %d attempts failed\n", att)
}

func RtSqf[T any, K any](att uint, dl time.Duration, fcr float64, prms []T, fn func(T) (K, error)) (K, error) {
	var err error
	var res K
	if len(prms) == 0 {
		return res, errors.New("empty params list\n")
	}

	for i := 0; i < int(att); i++ {
		for j, prm := range prms {
			res, err = fn(prm)
			if err == nil {
				return res, nil
			}
			fmt.Printf("attempt %d with param %d failed: %s\n", i+1, j+1, err)
			time.Sleep(dl)
		}
		dl = time.Duration(float64(dl) * fcr)
	}
	return res, fmt.Errorf("all %d attempts failed\n", att)
}

func RtRd[T any, K any](att uint, dl time.Duration, prms []T, fn func(T) (K, error)) (K, error) {

	var err error
	var res K
	if att <= 0 || len(prms) == 0 {
		return res, errors.New("invalid number of attpts or mpt prms list\n")
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < int(att); i++ {
		prm := prms[rng.Intn(len(prms))]
		res, err = fn(prm)
		if err == nil {
			return res, nil
		}
		fmt.Printf("attpts %d failed: %s\n", i+1, err)
		time.Sleep(dl)
	}

	return res, fmt.Errorf("all %d attpts has failed\n", att)
}

func RtSh[T any, K any](att uint, dl time.Duration, params []T, fn func(T) (K, error)) (K, error) {
	var err error
	var res K
	if att <= 0 || len(params) == 0 {
		return res, errors.New("invalid number of attpts or mpt prms list\n")
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < int(att); i++ {
		prms := make([]T, len(params))
		copy(prms, params)
		rng.Shuffle(len(prms), func(i, j int) {
			prms[i], prms[j] = prms[j], prms[i]
		})

		for j, prm := range prms {
			res, err = fn(prm)
			if err == nil {
				return res, nil
			}
			fmt.Printf("attpt %d with prm %d failed: %s\n", i+1, j+1, err)
			time.Sleep(dl)
		}
	}
	return res, fmt.Errorf("all %d attpts failed\n", att)
}

func RtSt[T any, K any](att uint, dl time.Duration, prms []T, stf func(a, b T) bool, fn func(T) (K, error)) (K, error) {
	var err error
	var res K
	if att <= 0 || len(prms) == 0 {
		return res, errors.New("invalid number of attpts or empty params list\n")
	}

	for i := 0; i < int(att); i++ {
		sort.Slice(prms, func(x, y int) bool {
			return stf(prms[x], prms[y])
		})

		for j, prm := range prms {
			res, err = fn(prm)
			if err == nil {
				return res, nil
			}
			fmt.Printf("attempt %d with param %d failed: %s\n", i+1, j+1, err)
			time.Sleep(dl)
		}
	}
	return res, fmt.Errorf("all %d attempts failed\n", att)
}
