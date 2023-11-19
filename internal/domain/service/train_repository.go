//go:generate mockgen -source=train_repository.go -destination=train_repository.mockgen.go -package=service

package service

type TrainRepository interface {
	Find(from string, to string) (Train, error)
	Create(from string, to string) Train
}
