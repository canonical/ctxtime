// Copyright 2020 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

// Package ctxtime provides access to either the current time, or a time
// that has been cached in a context. It is intended to provide consistent
// time information in tests and other similar situation whilst being as
// unintrusive as possible.
package ctxtime

import (
	"context"
	"time"
)

// Now determines the current time. If a non-zero time has been added to
// the given context using ContextWithTime then this time will be
// returned, otherwise the value of time.Now will be used.
func Now(ctx context.Context) time.Time {
	if t := timeFromContext(ctx); !t.IsZero() {
		return t
	}
	return time.Now()
}

// ContextWithTime attaches the given time to a context. All subsequent
// calls to Now will return this time. If the given time is the zero time,
// then Now will return the system time.
//
// It is expected that ContextWithTime is most useful for having consistent
// times in tests.
func ContextWithTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, timeKey{}, t)
}

// TimeFromContext returns the time attached to the given context, if any.
func timeFromContext(ctx context.Context) time.Time {
	t, _ := ctx.Value(timeKey{}).(time.Time)
	return t
}

type timeKey struct{}

// UTCSeconds is a convenience funtion that returns the time returned by
// Now, but in the UTC time zone and rounded to the nearest second.
func UTCSeconds(ctx context.Context) time.Time {
	return Now(ctx).UTC().Round(time.Second)
}

// UTCMilliseconds is a convenience funtion that returns the time returned
// by Now, but in the UTC time zone and rounded to the nearest millisecond.
func UTCMilliseconds(ctx context.Context) time.Time {
	return Now(ctx).UTC().Round(time.Millisecond)
}
