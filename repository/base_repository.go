package repository

type BaseRepository[T any] interface {
	Create(payload T) error
	List() ([]T, error)
}