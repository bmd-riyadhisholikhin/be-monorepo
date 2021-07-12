// Code generated by candi v1.5.32.

package repository

import (
	"context"
	"database/sql"
	"time"

	"pkg.agungdp.dev/candi/candishared"
	shareddomain "monorepo/services/shark/pkg/shared/domain"

	"pkg.agungdp.dev/candi/tracer"
)

type accountRepoSQL struct {
	readDB, writeDB *sql.DB
	tx              *sql.Tx
}

// NewAccountRepoSQL mongo repo constructor
func NewAccountRepoSQL(readDB, writeDB *sql.DB, tx *sql.Tx) AccountRepository {
	return &accountRepoSQL{
		readDB, writeDB, tx,
	}
}

func (r *accountRepoSQL) FetchAll(ctx context.Context, filter *candishared.Filter) (data []shareddomain.Account, err error) {
	trace := tracer.StartTrace(ctx, "AccountRepoSQL:FetchAll")
	defer func() { trace.SetError(err); trace.Finish() }()

	if filter.OrderBy == "" {
		filter.OrderBy = "modified_at"
	}
	
	
	return
}

func (r *accountRepoSQL) Count(ctx context.Context, filter *candishared.Filter) (count int) {
	trace := tracer.StartTrace(ctx, "AccountRepoSQL:Count")
	defer trace.Finish()

	var total int64
	count = int(total)
	return
}

func (r *accountRepoSQL) Find(ctx context.Context, data *shareddomain.Account) (err error) {
	trace := tracer.StartTrace(ctx, "AccountRepoSQL:Find")
	defer func() { trace.SetError(err); trace.Finish() }()

	return
}

func (r *accountRepoSQL) Save(ctx context.Context, data *shareddomain.Account) (err error) {
	trace := tracer.StartTrace(ctx, "AccountRepoSQL:Save")
	defer func() { trace.SetError(err); trace.Finish() }()
	tracer.Log(ctx, "data", data)

	data.ModifiedAt = time.Now()
	if data.CreatedAt.IsZero() {
		data.CreatedAt = time.Now()
	}
	return
}

func (r *accountRepoSQL) Delete(ctx context.Context, id string) (err error) {
	trace := tracer.StartTrace(ctx, "AccountRepoSQL:Save")
	defer func() { trace.SetError(err); trace.Finish() }()

	return
}
