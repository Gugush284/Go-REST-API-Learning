package store

type Store interface {
	User() UserRepository
	Image() ImageRepository
}
