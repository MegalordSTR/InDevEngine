package w32api

type RegisterClassExError struct{}

func (r RegisterClassExError) Error() string {
	return "RegisterClassExError error"
}
