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

package observer_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func assertEmpty(t testing.TB, logs *ObservedLogs) {
	assert.Equal(t, 0, logs.Len(), "Expected empty ObservedLogs to have zero length.")
	assert.Equal(t, []LoggedEntry{}, logs.All(), "Unexpected LoggedEntries in empty ObservedLogs.")
}

func TestObserver(t *testing.T) {
	for _, includeContext := range []bool{false, true} {
		t.Run(fmt.Sprint("withContext ", includeContext), func(t *testing.T) {
			logs := &ObservedLogs{}
			assertEmpty(t, logs)

			obs := zap.New(New(zap.InfoLevel, logs.Add, includeContext)).With(zap.Int("i", 1))
			obs.Info("foo")
			obs.Debug("bar")
			want := []LoggedEntry{{
				Entry:   zapcore.Entry{Level: zap.InfoLevel, Message: "foo"},
				Context: nil,
			}}
			if includeContext {
				want[0].Context = []zapcore.Field{zap.Int("i", 1)}
			}

			assert.Equal(t, 1, logs.Len(), "Unexpected observed logs Len.")
			assert.Equal(t, want, logs.AllUntimed(), "Unexpected contents from AllUntimed.")

			all := logs.All()
			require.Equal(t, 1, len(all), "Unexpected numbed of LoggedEntries returned from All.")
			assert.NotEqual(t, time.Time{}, all[0].Time, "Expected non-zero time on LoggedEntry.")

			// copy & zero time for stable assertions
			untimed := append([]LoggedEntry{}, all...)
			untimed[0].Time = time.Time{}
			assert.Equal(t, want, untimed, "Unexpected LoggedEntries from All.")

			assert.Equal(t, all, logs.TakeAll(), "Expected All and TakeAll to return identical results.")
			assertEmpty(t, logs)
		})
	}
}

func TestObserverWith(t *testing.T) {
	var logs ObservedLogs
	sf1 := New(zap.InfoLevel, logs.Add, true)

	// need to pad out enough initial fields so that the underlying slice cap()
	// gets ahead of its len() so that the sf3/4 With append's could choose
	// not to copy (if the implementation doesn't force them)
	sf1 = sf1.With([]zapcore.Field{zap.Int("a", 1), zap.Int("b", 2)})

	sf2 := sf1.With([]zapcore.Field{zap.Int("c", 3)})
	sf3 := sf2.With([]zapcore.Field{zap.Int("d", 4)})
	sf4 := sf2.With([]zapcore.Field{zap.Int("e", 5)})
	ent := zapcore.Entry{Level: zap.InfoLevel, Message: "hello"}

	for i, core := range []zapcore.Core{sf2, sf3, sf4} {
		if ce := core.Check(ent, nil); ce != nil {
			ce.Write(zap.Int("i", i))
		}
	}

	assert.Equal(t, []LoggedEntry{
		{
			Entry: ent,
			Context: []zapcore.Field{
				zap.Int("a", 1),
				zap.Int("b", 2),
				zap.Int("c", 3),
				zap.Int("i", 0),
			},
		},
		{
			Entry: ent,
			Context: []zapcore.Field{
				zap.Int("a", 1),
				zap.Int("b", 2),
				zap.Int("c", 3),
				zap.Int("d", 4),
				zap.Int("i", 1),
			},
		},
		{
			Entry: ent,
			Context: []zapcore.Field{
				zap.Int("a", 1),
				zap.Int("b", 2),
				zap.Int("c", 3),
				zap.Int("e", 5),
				zap.Int("i", 2),
			},
		},
	}, logs.All(), "expected no field sharing between With siblings")
}
