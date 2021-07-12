// Code generated by candi v1.5.32.

package usecase

import (
	"context"

	shareddomain "monorepo/services/shark/pkg/shared/domain"
	"monorepo/services/shark/pkg/shared/repository"
	"monorepo/services/shark/pkg/shared/usecase/common"

	"github.com/google/uuid"
	"pkg.agungdp.dev/candi/candishared"
	"pkg.agungdp.dev/candi/codebase/factory/dependency"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

type contactUsecaseImpl struct {
	sharedUsecase common.Usecase
	cache         interfaces.Cache
	repoSQL   repository.RepoSQL
	
}

// NewContactUsecase usecase impl constructor
func NewContactUsecase(deps dependency.Dependency) (ContactUsecase, func(sharedUsecase common.Usecase)) {
	uc := &contactUsecaseImpl{
		cache: deps.GetRedisPool().Cache(),
		repoSQL:   repository.GetSharedRepoSQL(),
		
	}
	return uc, func(sharedUsecase common.Usecase) {
		uc.sharedUsecase = sharedUsecase
	}
}

func (uc *contactUsecaseImpl) GetAllContact(ctx context.Context, filter candishared.Filter) (data []shareddomain.Contact, meta candishared.Meta, err error) {
	trace := tracer.StartTrace(ctx, "ContactUsecase:GetAllContact")
	defer trace.Finish()
	ctx = trace.Context()

	data, err = uc.repoSQL.ContactRepo().FetchAll(ctx, &filter)
	if err != nil {
		return data, meta, err
	}
	count := uc.repoSQL.ContactRepo().Count(ctx, &filter)
	meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	return
}

func (uc *contactUsecaseImpl) GetDetailContact(ctx context.Context, id string) (data shareddomain.Contact, err error) {
	trace := tracer.StartTrace(ctx, "ContactUsecase:GetDetailContact")
	defer trace.Finish()
	ctx = trace.Context()

	data.ID = id
	err = uc.repoSQL.ContactRepo().Find(ctx, &data)
	return
}

func (uc *contactUsecaseImpl) CreateContact(ctx context.Context, data *shareddomain.Contact) (err error) {
	trace := tracer.StartTrace(ctx, "ContactUsecase:CreateContact")
	defer trace.Finish()
	ctx = trace.Context()

	data.ID = uuid.NewString()
	return uc.repoSQL.ContactRepo().Save(ctx, data)
}

func (uc *contactUsecaseImpl) UpdateContact(ctx context.Context, id string, data *shareddomain.Contact) (err error) {
	trace := tracer.StartTrace(ctx, "ContactUsecase:UpdateContact")
	defer trace.Finish()
	ctx = trace.Context()

	existing := &shareddomain.Contact{ID: id}
	if err := uc.repoSQL.ContactRepo().Find(ctx, existing); err != nil {
		return err
	}
	data.ID = existing.ID
	data.CreatedAt = existing.CreatedAt
	return  uc.repoSQL.ContactRepo().Save(ctx, data)
}

func (uc *contactUsecaseImpl) DeleteContact(ctx context.Context, id string) (err error) {
	trace := tracer.StartTrace(ctx, "ContactUsecase:DeleteContact")
	defer trace.Finish()
	ctx = trace.Context()

	return uc.repoSQL.ContactRepo().Delete(ctx, id)
}
