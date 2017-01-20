// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package etcd

import (
	"errors"
	"time"

	"github.com/m3db/m3x/instrument"
	"github.com/m3db/m3x/retry"
)

var (
	defaultRequestTimeout         = 10 * time.Second
	defaultDialTimeout            = 10 * time.Second
	defaultWatchChanCheckInterval = 10 * time.Second
	defaultWatchChanResetInterval = 10 * time.Second
	defaultRetryOptions           = xretry.NewOptions().SetMaxRetries(5)
	defaultKeyFn                  = KeyFn(
		func(key string) string {
			return key
		},
	)
)

// KeyFn is a function that wraps a key
type KeyFn func(key string) string

// Options are options for the client of the kv store
type Options interface {
	// RequestTimeout is the timeout for etcd requests
	RequestTimeout() time.Duration
	// SetRequestTimeout sets the RequestTimeout
	SetRequestTimeout(t time.Duration) Options

	// KeyFn is the function to wrap a key
	KeyFn() KeyFn
	// SetKeyFn sets the KeyFn
	SetKeyFn(f KeyFn) Options

	// InstrumentsOptions is the instrument options
	InstrumentsOptions() instrument.Options
	// SetInstrumentsOptions sets the InstrumentsOptions
	SetInstrumentsOptions(iopts instrument.Options) Options

	// RetryOptions is the retry options
	RetryOptions() xretry.Options
	// SetRetryOptions sets the RetryOptions
	SetRetryOptions(ropts xretry.Options) Options

	// WatchChanCheckInterval will be used to periodicaly check if a watch chan
	// is no longer being subscribed and should be closed
	WatchChanCheckInterval() time.Duration
	// SetWatchChanCheckInterval sets the WatchChanCheckInterval
	SetWatchChanCheckInterval(t time.Duration) Options

	// WatchChanResetInterval is the delay before resetting the etcd watch chan
	WatchChanResetInterval() time.Duration
	// SetWatchChanResetInterval sets the WatchChanResetInterval
	SetWatchChanResetInterval(t time.Duration) Options

	// CacheFilePath is the file path to persist in-memory cache.
	// If not provided, not file persisting will happen
	CacheFilePath() string
	// SetCacheFilePath sets the CacheFilePath
	SetCacheFilePath(c string) Options

	// Validate validates the Options
	Validate() error
}

type options struct {
	requestTimeout         time.Duration
	keyFn                  KeyFn
	iopts                  instrument.Options
	ropts                  xretry.Options
	watchChanCheckInterval time.Duration
	watchChanResetInterval time.Duration
	cacheFilePath          string
}

// NewOptions creates a sane default Option
func NewOptions() Options {
	o := options{}
	return o.SetRequestTimeout(defaultRequestTimeout).
		SetInstrumentsOptions(instrument.NewOptions()).
		SetRetryOptions(defaultRetryOptions).
		SetWatchChanCheckInterval(defaultWatchChanCheckInterval).
		SetWatchChanResetInterval(defaultWatchChanResetInterval).
		SetKeyFn(defaultKeyFn)
}

func (o options) Validate() error {
	if o.iopts == nil {
		return errors.New("no instrument options")
	}

	if o.ropts == nil {
		return errors.New("no retry options")
	}

	if o.watchChanCheckInterval <= 0 {
		return errors.New("invalid watch channel check interval")
	}

	if o.keyFn == nil {
		return errors.New("no keyFn set")
	}

	return nil
}

func (o options) RequestTimeout() time.Duration {
	return o.requestTimeout
}

func (o options) SetRequestTimeout(t time.Duration) Options {
	o.requestTimeout = t
	return o
}

func (o options) KeyFn() KeyFn {
	return o.keyFn
}

func (o options) SetKeyFn(f KeyFn) Options {
	o.keyFn = f
	return o
}

func (o options) InstrumentsOptions() instrument.Options {
	return o.iopts
}

func (o options) SetInstrumentsOptions(iopts instrument.Options) Options {
	o.iopts = iopts
	return o
}

func (o options) RetryOptions() xretry.Options {
	return o.ropts
}

func (o options) SetRetryOptions(ropts xretry.Options) Options {
	o.ropts = ropts
	return o
}

func (o options) WatchChanCheckInterval() time.Duration {
	return o.watchChanCheckInterval
}

func (o options) SetWatchChanCheckInterval(t time.Duration) Options {
	o.watchChanCheckInterval = t
	return o
}

func (o options) WatchChanResetInterval() time.Duration {
	return o.watchChanResetInterval
}

func (o options) SetWatchChanResetInterval(t time.Duration) Options {
	o.watchChanResetInterval = t
	return o
}

func (o options) CacheFilePath() string {
	return o.cacheFilePath
}

func (o options) SetCacheFilePath(c string) Options {
	o.cacheFilePath = c
	return o
}