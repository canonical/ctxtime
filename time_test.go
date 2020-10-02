// Copyright 2020 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package ctxtime_test

import (
	"context"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"

	"github.com/canonical/ctxtime"
)

func TestNow(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()

	// Now returns the current time.
	start := time.Now()
	now := ctxtime.Now(ctx)
	c.Assert(now.After(start), qt.IsTrue)

	// Subsequent calls still return the current time.
	now2 := ctxtime.Now(ctx)
	c.Assert(now2.After(now), qt.IsTrue)
}

func TestContextWithTime(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()

	// TestingContextWithTime changes the default behavior of now.
	future := time.Now().Add(10 * time.Hour)
	ctx = ctxtime.ContextWithTime(ctx, future)
	now := ctxtime.Now(ctx)
	c.Assert(now, qt.Equals, future)
}

func TestContextWithTimeMultipleCalls(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()

	// Subsequent calls of ContextWithTime override the stored time.
	t1 := time.Now()
	t2 := t1.Add(-1 * time.Hour)
	ctx = ctxtime.ContextWithTime(ctx, t1)
	c.Assert(ctxtime.Now(ctx), qt.Equals, t1)
	ctx = ctxtime.ContextWithTime(ctx, t2)
	c.Assert(ctxtime.Now(ctx), qt.Equals, t2)
}

func TestNowZeroTimeInContext(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()

	// Now ignores zero time in context.
	ctx = ctxtime.ContextWithTime(ctx, time.Time{})
	now := ctxtime.Now(ctx)
	c.Assert(now.IsZero(), qt.IsFalse)
}

func TestUTCSeconds(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()

	tm := ctxtime.UTCSeconds(ctx)
	zone, offset := tm.Zone()
	c.Assert(zone, qt.Equals, "UTC")
	c.Assert(offset, qt.Equals, 0)
	c.Assert(tm.Nanosecond(), qt.Equals, 0)
}

func TestUTCMilliseconds(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()

	tm := ctxtime.UTCMilliseconds(ctx)
	zone, offset := tm.Zone()
	c.Assert(zone, qt.Equals, "UTC")
	c.Assert(offset, qt.Equals, 0)
	c.Assert(tm.Nanosecond()%int(time.Millisecond), qt.Equals, 0)
}
