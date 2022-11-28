package generators

import "migueleliasweb.github.io/api/v1beta1"

type Generator[T any] interface {
	GenerateFromApp(v1beta1.App) T
}
