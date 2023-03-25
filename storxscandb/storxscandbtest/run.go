// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package storxscandbtest

import (
	"context"
	"strings"
	"testing"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"common/testcontext"
	"private/dbutil"
	"private/dbutil/pgtest"
	"private/dbutil/pgutil"
	"private/dbutil/tempdb"
	"storxscan"
	"storxscan/storxscandb"
)

// Checks that test db implements storxscan.DB.
var _ storxscan.DB = (*DB)(nil)

// Run creates new storxscan test database, create tables and execute test function against that db.
func Run(t *testing.T, test func(ctx *testcontext.Context, t *testing.T, db *DB)) {
	t.Run("Postgres", func(t *testing.T) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		connStr := pgtest.PickPostgres(t)

		db, err := OpenDB(ctx, zaptest.NewLogger(t), connStr, t.Name(), "T")
		if err != nil {
			t.Fatal(err)
		}
		defer ctx.Check(db.Close)

		err = db.MigrateToLatest(ctx)
		if err != nil {
			t.Fatal(err)
		}

		test(ctx, t, db)
	})

	t.Run("Cockroach", func(t *testing.T) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		connStr := pgtest.PickCockroach(t)

		db, err := OpenDB(ctx, zaptest.NewLogger(t), connStr, t.Name(), "T")
		if err != nil {
			t.Fatal(err)
		}
		defer ctx.Check(db.Close)

		err = db.MigrateToLatest(ctx)
		if err != nil {
			t.Fatal(err)
		}

		test(ctx, t, db)
	})
}

// DB is test storxscan database with unique schema which performs cleanup on close.
type DB struct {
	*storxscandb.DB
	tempDB *dbutil.TempDatabase
}

// OpenDB opens new unique temp storxscan test database.
func OpenDB(ctx context.Context, log *zap.Logger, connStr, testName, category string) (*DB, error) {
	schemaSuffix := pgutil.CreateRandomTestingSchemaName(6)
	schemaName := schemaName(testName, schemaSuffix, category)

	tempDB, err := tempdb.OpenUnique(ctx, connStr, schemaName)
	if err != nil {
		return nil, err
	}
	storxscanDB, err := storxscandb.Open(ctx, log, tempDB.ConnStr)
	if err != nil {
		return nil, errs.Combine(err, tempDB.Close())
	}

	return &DB{
		DB:     storxscanDB,
		tempDB: tempDB,
	}, nil
}

// Close closes test database and performs cleanup.
func (db *DB) Close() error {
	return errs.Combine(db.DB.Close(), db.tempDB.Close())
}

// schemaName create new postgres db schema name for testing.
func schemaName(testName, suffix, category string) string {
	maxTestNameLength := 64 - len(suffix) - len(category)
	if len(testName) > maxTestNameLength {
		testName = testName[:maxTestNameLength]
	}
	return strings.ToLower(testName + "/" + suffix + "/" + category)
}
