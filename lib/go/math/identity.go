package math

const (
	SumIdentity     = 0
	ProductIdentity = 1
)

func Identity[T any](e T) T {
	return e
}
