// Copyright (C) 2022 Storx Labs, Inc.
// See LICENSE for copying information.

package tokenprice_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"common/currency"
	"common/testcontext"
	"storxscan/storxscandb/storxscandbtest"
)

func TestPriceQuoteDBBefore(t *testing.T) {
	storxscandbtest.Run(t, func(ctx *testcontext.Context, t *testing.T, db *storxscandbtest.DB) {
		tokenPriceDB := db.TokenPrice()
		now := time.Now().Truncate(time.Second).UTC()

		const priceCount = 10
		for i := 0; i < priceCount; i++ {
			require.NoError(t, tokenPriceDB.Update(ctx, now.Add(time.Duration(i)*time.Second), int64(i)*1000000))
		}

		pq, err := tokenPriceDB.Before(ctx, now.Add(priceCount*time.Second))
		require.NoError(t, err)
		require.Equal(t, now.Add((priceCount-1)*time.Second), pq.Timestamp.UTC())
		require.EqualValues(t, currency.AmountFromBaseUnits((priceCount-1)*1000000, currency.USDollarsMicro), pq.Price)
	})
}
