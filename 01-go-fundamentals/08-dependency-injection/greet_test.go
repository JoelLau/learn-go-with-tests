package dependencyinjection

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	for idx, tt := range []struct {
		input  string
		expect string
	}{
		{input: "boo", expect: "Hello, boo!"},
	} {
		name := fmt.Sprintf("test %d", idx)

		t.Run(name, func(t *testing.T) {
			buffer := bytes.Buffer{}
			Greet(&buffer, tt.input)

			have := buffer.String()
			want := tt.expect

			if have != want {
				t.Errorf("have %q want %q", have, want)
			}
		})
	}
}
