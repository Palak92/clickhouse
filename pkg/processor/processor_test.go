package processor

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProcessorGivesTopUrls(t *testing.T) {
	p := New()

	p.size = 2
	p.Add("url1 3")
	p.Add("url2 3")
	p.Add("url3 2")
	p.Add("url4 1")

	want := []string{"url1", "url2", "url3"}
	got := p.TopUrls()

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("processor.TopUrls(): -got , +want %v", diff)
	}

}
