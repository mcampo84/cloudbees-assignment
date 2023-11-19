//go:generate mockgen -source=train.go -destination=train.mockgen.go -package=service

package service

type Train interface {
	GetID() uint
	GetSections() []Section
}
