// Copyright (C) 2022 Storx Labs, Inc.
// See LICENSE for copying information.

package dbx

import (
	// make sure we load our cockroach driver so dbx.Open can find it.
	_ "private/dbutil/cockroachutil"
)
