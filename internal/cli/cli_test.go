package cli_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/simurg-hq/cli/internal/cli"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string // substring; empty means stderr must be empty
	}{
		{"version", []string{"version"}, 0, "simurg v1.2.3\n", ""},
		{"no command", nil, 2, "", "usage: simurg <command>"},
		{"unknown command", []string{"run"}, 2, "", `unknown command "run"`},
		{"version with arguments", []string{"version", "extra"}, 2, "", "version takes no arguments"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := cli.Run(tt.args, "v1.2.3", &stdout, &stderr)

			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d", code, tt.wantCode)
			}
			if stdout.String() != tt.wantStdout {
				t.Errorf("stdout = %q, want %q", stdout.String(), tt.wantStdout)
			}
			if tt.wantStderr == "" && stderr.Len() > 0 {
				t.Errorf("stderr = %q, want empty", stderr.String())
			}
			if !strings.Contains(stderr.String(), tt.wantStderr) {
				t.Errorf("stderr = %q, want it to contain %q", stderr.String(), tt.wantStderr)
			}
		})
	}
}

// failingWriter fails every write, like stdout redirected to a full disk.
type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) {
	return 0, errors.New("disk full")
}

func TestRunVersionFailsWhenStdoutFails(t *testing.T) {
	var stderr bytes.Buffer
	if code := cli.Run([]string{"version"}, "v1.2.3", failingWriter{}, &stderr); code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
	if !strings.Contains(stderr.String(), "disk full") {
		t.Errorf("stderr = %q, want it to name the write error", stderr.String())
	}
}
