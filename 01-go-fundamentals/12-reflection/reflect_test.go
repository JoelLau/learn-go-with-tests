package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("walk", func(t *testing.T) {
		expected := "Chris"
		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		have := len(got)
		want := 1
		if have != want {
			t.Errorf("wrong number of function calls, have %d want %d", have, want)
		}

		if got[0] != expected {
			t.Errorf("have %q want %q", got[0], expected)
		}
	})

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name:          "struct with one string field",
			Input:         struct{ Name string }{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two string field",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "nested fields",
			Input: Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pointer to tests",
			Input: &Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name:          "slices",
			Input:         []Profile{{33, "Chris"}, {33, "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var have []string
			walk(test.Input, func(input string) {
				have = append(have, input)
			})

			want := test.ExpectedCalls
			if !reflect.DeepEqual(have, want) {
				t.Errorf("have %v, want %v", have, want)
			}
		})
	}
}
