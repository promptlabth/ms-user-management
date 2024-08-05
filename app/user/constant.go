package user

const freePlanType string = "Free"

func TypeToPrt[T any](val T) *T {
	return &val
}

func PtrToType[T any](val *T) T {
	if val == nil {
		return *new(T)
	}
	return *val
}
