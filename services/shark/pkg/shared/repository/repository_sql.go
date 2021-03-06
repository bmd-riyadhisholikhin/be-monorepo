// Code generated by candi v1.5.32.

package repository

import (
	"context"
	"database/sql"
	"fmt"

	// @candi:repositoryImport

	"pkg.agungdp.dev/candi/tracer"
)

type (
	// RepoSQL abstraction
	RepoSQL interface {
		WithTransaction(ctx context.Context, txFunc func(ctx context.Context, repo RepoSQL) error) (err error)
		Free()

		// @candi:repositoryMethod
	}

	repoSQLImpl struct {
		readDB, writeDB *sql.DB
		tx    *sql.Tx
	
		// register all repository from modules
		// @candi:repositoryField
	}
)

var (
	globalRepoSQL RepoSQL
)

// setSharedRepoSQL set the global singleton "RepoSQL" implementation
func setSharedRepoSQL(readDB, writeDB *sql.DB) {
	
	globalRepoSQL = NewRepositorySQL(readDB, writeDB, nil)
}

// GetSharedRepoSQL returns the global singleton "RepoSQL" implementation
func GetSharedRepoSQL() RepoSQL {
	return globalRepoSQL
}

// NewRepositorySQL constructor
func NewRepositorySQL(readDB, writeDB *sql.DB, tx *sql.Tx) RepoSQL {

	return &repoSQLImpl{
		readDB: readDB, writeDB: writeDB, tx: tx,

		// @candi:repositoryConstructor
	}
}

// WithTransaction run transaction for each repository with context, include handle canceled or timeout context
func (r *repoSQLImpl) WithTransaction(ctx context.Context, txFunc func(ctx context.Context, repo RepoSQL) error) (err error) {
	trace := tracer.StartTrace(ctx, "RepoSQL:Transaction")
	defer trace.Finish()
	ctx = trace.Context()

	tx, err := r.writeDB.Begin()
	if err != nil {
		return err
	}

	// reinit new repository in different memory address with tx value
	manager := NewRepositorySQL(r.readDB, r.writeDB, tx)
	defer func() {
		if err != nil {
			tx.Rollback()
			trace.SetError(err)
		} else {
			tx.Commit()
		}
		manager.Free()
	}()

	errChan := make(chan error)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic: %v", r)
			}
			close(errChan)
		}()

		if err := txFunc(ctx, manager); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("Canceled or timeout: %v", ctx.Err())
	case e := <-errChan:
		return e
	}
}

func (r *repoSQLImpl) Free() {
	// make nil all repository
	// @candi:repositoryDestructor
}

// @candi:repositoryImplementation
