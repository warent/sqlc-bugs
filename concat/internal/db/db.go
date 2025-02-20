// Code generated by sqlc. DO NOT EDIT.

package gendb

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.testStmt, err = db.PrepareContext(ctx, test); err != nil {
		return nil, fmt.Errorf("error preparing query Test: %w", err)
	}
	if q.test3Stmt, err = db.PrepareContext(ctx, test3); err != nil {
		return nil, fmt.Errorf("error preparing query Test3: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.testStmt != nil {
		if cerr := q.testStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing testStmt: %w", cerr)
		}
	}
	if q.test3Stmt != nil {
		if cerr := q.test3Stmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing test3Stmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db        DBTX
	tx        *sql.Tx
	testStmt  *sql.Stmt
	test3Stmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:        tx,
		tx:        tx,
		testStmt:  q.testStmt,
		test3Stmt: q.test3Stmt,
	}
}
