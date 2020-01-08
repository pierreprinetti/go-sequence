package sequence_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pierreprinetti/go-sequence"
)

func TestInt(t *testing.T) {
	type checkFunc func([]int, error) error
	check := func(fns ...checkFunc) []checkFunc { return fns }

	wantError := func(want error) checkFunc {
		return func(_ []int, have error) error {
			if have != want {
				return fmt.Errorf("expected %v, found %v", want, have)
			}
			return nil
		}
	}

	wantSeq := func(want ...int) checkFunc {
		return func(have []int, _ error) error {
			if !reflect.DeepEqual(want, have) {
				return fmt.Errorf("expected %v, found %v", want, have)
			}
			return nil
		}
	}

	for _, tc := range [...]struct {
		in     string
		checks []checkFunc
	}{
		{
			"",
			check(
				wantSeq(),
				wantError(nil),
			),
		},
		{
			"1",
			check(
				wantSeq(1),
				wantError(nil),
			),
		},
		{
			"0",
			check(
				wantSeq(0),
				wantError(nil),
			),
		},
		{
			"1,9,3",
			check(
				wantSeq(1, 9, 3),
				wantError(nil),
			),
		},
		{
			" 1, 2,3 ,77, 5",
			check(
				wantSeq(1, 2, 3, 77, 5),
				wantError(nil),
			),
		},
		{
			"1,2,2,2,2,2,2,6,2",
			check(
				wantSeq(1, 2, 2, 2, 2, 2, 2, 6, 2),
				wantError(nil),
			),
		},
		{
			"576-580",
			check(
				wantSeq(576, 577, 578, 579, 580),
				wantError(nil),
			),
		},
		{
			"7-2",
			check(
				wantSeq(7, 6, 5, 4, 3, 2),
				wantError(nil),
			),
		},
		{
			"15-15",
			check(
				wantSeq(15),
				wantError(nil),
			),
		},
	} {
		t.Run(tc.in, func(t *testing.T) {
			s, e := sequence.Int(tc.in)
			for _, check := range tc.checks {
				if err := check(s, e); err != nil {
					t.Error(err)
				}
			}
		})
	}
}
