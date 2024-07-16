package rfc

func Drf(v any) any {
	if ptr, ok := v.(**any); ok {
		return *ptr
	}
	return v
}
