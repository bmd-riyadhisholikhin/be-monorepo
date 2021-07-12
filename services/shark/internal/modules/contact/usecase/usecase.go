// Code generated by candi v1.5.32.

package usecase

import (
	"context"

	shareddomain "monorepo/services/shark/pkg/shared/domain"

	"pkg.agungdp.dev/candi/candishared"
)

// ContactUsecase abstraction
type ContactUsecase interface {
	GetAllContact(ctx context.Context, filter candishared.Filter) (data []shareddomain.Contact, meta candishared.Meta, err error)
	GetDetailContact(ctx context.Context, id string) (data shareddomain.Contact, err error)
	CreateContact(ctx context.Context, data *shareddomain.Contact) (err error)
	UpdateContact(ctx context.Context, id string, data *shareddomain.Contact) (err error)
	DeleteContact(ctx context.Context, id string) (err error)
}
