package version

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := String()
	if !strings.HasPrefix(s, "voxctl ") {
		t.Errorf("expected prefix 'voxctl ', got %q", s)
	}
	if !strings.Contains(s, "commit:") {
		t.Errorf("expected 'commit:' in version string, got %q", s)
	}
	if !strings.Contains(s, "built:") {
		t.Errorf("expected 'built:' in version string, got %q", s)
	}
}
