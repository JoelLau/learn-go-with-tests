package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}

func TestRepeat(t *testing.T) {
	t.Run("repeats 'a' 5 times", func(t *testing.T)  {
		have := Repeat("a")
		want := "aaaaa"

		assertString(t, have, want)
	})
	
	t.Run("repeats 'x' 5 times", func(t *testing.T)  {
		have := Repeat("x")
		want := "xxxxx"

		assertString(t, have, want)
	})
}


func assertString(t testing.TB, have, want string) {
	t.Helper()

	if have != want {
		t.Errorf("have '%s' want '%s'", have, want)
	}
}
