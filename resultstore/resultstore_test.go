/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resultstore

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/genproto/googleapis/devtools/resultstore/v2"
)

func deepEqual(have, want interface{}) bool {
	return reflect.DeepEqual(have, want)
}

func diff(have, want interface{}) string {
	// TODO(fejta): something fancier
	return fmt.Sprintf("got %v, want %v", have, want)
}

func TestDur(t *testing.T) {
	cases := []struct {
		name     string
		dur      time.Duration
		expected *duration.Duration
	}{
		{
			name:     "basically works",
			expected: &duration.Duration{},
		},
		{
			name: "correct seconds",
			dur:  time.Minute,
			expected: &duration.Duration{
				Seconds: 60,
			},
		},
		{
			name: "correct nanos",
			dur:  300 * time.Nanosecond,
			expected: &duration.Duration{
				Nanos: 300,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := dur(tc.dur); !deepEqual(actual, tc.expected) {
				t.Errorf(diff(actual, tc.expected))
			}
		})
	}
}

func TestStamp(t *testing.T) {
	cases := []struct {
		name     string
		when     time.Time
		expected *timestamp.Timestamp
	}{
		{
			name: "basically works",
		},
		{
			name: "correct when only seconds",
			when: time.Unix(15, 0),
			expected: &timestamp.Timestamp{
				Seconds: 15,
			},
		},
		{
			name: "correct when only nanos",
			when: time.Unix(0, 22),
			expected: &timestamp.Timestamp{
				Nanos: 22,
			},
		},
		{
			name: "normal",
			when: time.Unix(33, 4444),
			expected: &timestamp.Timestamp{
				Seconds: 33,
				Nanos:   4444,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := stamp(tc.when); !deepEqual(actual, tc.expected) {
				t.Errorf(diff(actual, tc.expected))
			}
		})
	}
}

func TestTiming(t *testing.T) {
	now := time.Now()
	cases := []struct {
		name     string
		when     time.Time
		d        time.Duration
		expected *resultstore.Timing
	}{
		{
			name: "basically works",
		},
		{
			name: "only when",
			when: now,
			expected: &resultstore.Timing{
				StartTime: stamp(now),
			},
		},
		{
			name: "only duration",
			d:    time.Second,
			expected: &resultstore.Timing{
				Duration: dur(time.Second),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := timing(tc.when, tc.d)
			if !deepEqual(actual, tc.expected) {
				t.Errorf(diff(actual, tc.expected))
			}
		})
	}
}

func TestFromTiming(t *testing.T) {
	cases := []struct {
		name string
		t    *resultstore.Timing
		when time.Time
		dur  time.Duration
	}{
		{
			name: "basically works",
		},
		{
			name: "only StartTime works",
			t: &resultstore.Timing{
				StartTime: &timestamp.Timestamp{
					Seconds: 15,
					Nanos:   7,
				},
			},
			when: time.Unix(15, 7),
		},
		{
			name: "only Duration works",
			t: &resultstore.Timing{
				Duration: &duration.Duration{
					Seconds: 3,
					Nanos:   4,
				},
			},
			dur: 3*time.Second + 4*time.Nanosecond,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			when, dur := fromTiming(tc.t)
			if !when.Equal(tc.when) {
				t.Errorf("when: %v != expected %v", when, tc.when)
			}
			if dur != tc.dur {
				t.Errorf("dur: %v != expected %v", dur, tc.dur)
			}
		})
	}
}

func TestProperties(t *testing.T) {
	cases := []struct {
		name     string
		shock    bool
		input    []string
		expected []Property
	}{
		{
			name: "basically works",
		},
		{
			name:  "one pair works",
			input: []string{"hello", "world"},
			expected: []Property{
				{
					Key:   "hello",
					Value: "world",
				},
			},
		},
		{
			name:  "two pairs work",
			input: []string{"key1", "value1", "key2", "value2"},
			expected: []Property{
				{
					Key:   "key1",
					Value: "value1",
				},
				{
					Key:   "key2",
					Value: "value2",
				},
			},
		},
		{
			name:  "unbalanced pairs panic",
			input: []string{"key1", "value1", "panic"},
			shock: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var shocked bool
			func() {
				defer func() {
					if r := recover(); r != nil {
						shocked = true
					}
				}()
				actual := Properties(tc.input...)
				if !deepEqual(actual, tc.expected) {
					t.Errorf(diff(actual, tc.expected))
				}
			}()
			if shocked != tc.shock {
				t.Errorf("shock %t != expected %t", shocked, tc.shock)
			}
		})
	}
}

func TestFromTarget(t *testing.T) {
	now := time.Now()
	cases := []struct {
		name     string
		t        *resultstore.Target
		expected Target
	}{
		{
			name: "empty fromTarget works",
			t: &resultstore.Target{
				Name: "Empty Target",
			},
			expected: Target{
				Name: "Empty Target",
			},
		},
		{
			name: "fromTarget with empty properties",
			t: &resultstore.Target{
				Name: "Empty Properties",
				TargetAttributes: &resultstore.TargetAttributes{
					Tags: []string{
						"attr1",
					},
				},
			},
			expected: Target{
				Name: "Empty Properties",
				Tags: []string{
					"attr1",
				},
			},
		},
		{
			name: "fromTarget with properties",
			t: &resultstore.Target{
				Name: "Contain Properties",
				TargetAttributes: &resultstore.TargetAttributes{
					Tags: []string{
						"attr1",
						"attr2",
					},
				},
				Properties: []*resultstore.Property{
					{
						Key:   "key1",
						Value: "val1",
					},
					{
						Key:   "key2",
						Value: "val2",
					},
				},
			},
			expected: Target{
				Name: "Contain Properties",
				Tags: []string{
					"attr1",
					"attr2",
				},
				Properties: []Property{
					{
						Key:   "key1",
						Value: "val1",
					},
					{
						Key:   "key2",
						Value: "val2",
					},
				},
			},
		},
		{
			name: "test all attributes of target",
			t: &resultstore.Target{
				Name: "All Attributes",
				TargetAttributes: &resultstore.TargetAttributes{
					Tags: []string{
						"attr1",
					},
				},
				Properties: []*resultstore.Property{
					{
						Key:   "key1",
						Value: "val1",
					},
				},
				Timing: &resultstore.Timing{
					StartTime: stamp(now),
					Duration:  dur(time.Second),
				},
				StatusAttributes: &resultstore.StatusAttributes{
					Status:      resultstore.Status_STATUS_UNSPECIFIED,
					Description: "description1",
				},
			},
			expected: Target{
				Name: "All Attributes",
				Tags: []string{
					"attr1",
				},
				Properties: []Property{
					{
						Key:   "key1",
						Value: "val1",
					},
				},
				Start:       now,
				Duration:    time.Second,
				Status:      resultstore.Status_STATUS_UNSPECIFIED,
				Description: "description1",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tgt := fromTarget(tc.t)
			if !deepEqual(tgt.Name, tc.expected.Name) {
				t.Errorf(diff(tgt.Name, tc.expected.Name))
			}
			if !deepEqual(tgt.Tags, tc.expected.Tags) {
				t.Errorf(diff(tgt.Tags, tc.expected.Tags))
			}
			if !deepEqual(tgt.Properties, tc.expected.Properties) {
				t.Errorf(diff(tgt.Properties, tc.expected.Properties))
			}
			if !deepEqual(tgt.Duration, tc.expected.Duration) {
				t.Errorf(diff(tgt.Duration, tc.expected.Duration))
			}
			if !deepEqual(tgt.Status, tc.expected.Status) {
				t.Errorf(diff(tgt.Status, tc.expected.Status))
			}
			if !deepEqual(tgt.Description, tc.expected.Description) {
				t.Errorf(diff(tgt.Description, tc.expected.Description))
			}
		})
	}
}
