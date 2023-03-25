// Copyright (C) 2022 Storx Labs, Inc.
// See LICENSE for copying information.

package storxscandbtest_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"common/testcontext"
	"storxscan/storxscandb/storxscandbtest"
)

func TestRun(t *testing.T) {
	storxscandbtest.Run(t, func(ctx *testcontext.Context, t *testing.T, db *storxscandbtest.DB) {
		tableCmd := `CREATE TABLE test ( 
			number bigint NOT NULL, 
			PRIMARY KEY (number)
		)`
		_, err := db.Exec(ctx, tableCmd)
		require.NoError(t, err)

		_, err = db.Exec(ctx, "INSERT INTO test (number) VALUES ($1)", int64(1))
		require.NoError(t, err)

		row := db.QueryRowContext(ctx, "SELECT number FROM test")
		require.NoError(t, row.Err())
		var num int64
		require.NoError(t, row.Scan(&num))
		require.Equal(t, int64(1), num)
	})
}
