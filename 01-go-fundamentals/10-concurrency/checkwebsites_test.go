package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockCheckWebsite(website string) bool {
	return website != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	t.Run("check websites", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}

		have := CheckWebsites(mockCheckWebsite, websites)
		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}

		if !reflect.DeepEqual(have, want) {
			t.Errorf("have %v want %v", have, want)
		}
	})
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	websites := make([]string, 100)
	for i := 0; i < len(websites); i++ {
		websites[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, websites)
	}
}
