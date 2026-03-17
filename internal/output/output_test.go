package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestPrintTable(t *testing.T) {
	var buf bytes.Buffer
	columns := []string{"NAME", "STATUS"}
	rows := [][]string{
		{"node1", "active"},
		{"node2", "inactive"},
	}

	if err := Fprint(&buf, "table", rows, columns); err != nil {
		t.Fatalf("Fprint table: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "NAME") {
		t.Error("expected header NAME in output")
	}
	if !strings.Contains(out, "node1") {
		t.Error("expected node1 in output")
	}
	if !strings.Contains(out, "node2") {
		t.Error("expected node2 in output")
	}
}

func TestPrintTableEmpty(t *testing.T) {
	var buf bytes.Buffer
	columns := []string{"NAME"}
	rows := [][]string{}

	if err := Fprint(&buf, "table", rows, columns); err != nil {
		t.Fatalf("Fprint table: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "NAME") {
		t.Error("expected header in output even with no rows")
	}
}

func TestPrintJSON(t *testing.T) {
	var buf bytes.Buffer
	data := map[string]string{"key": "value"}

	if err := Fprint(&buf, "json", data, nil); err != nil {
		t.Fatalf("Fprint json: %v", err)
	}

	var result map[string]string
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}
	if result["key"] != "value" {
		t.Errorf("expected key=value, got %q", result["key"])
	}
}

func TestPrintYAML(t *testing.T) {
	var buf bytes.Buffer
	data := map[string]string{"key": "value"}

	if err := Fprint(&buf, "yaml", data, nil); err != nil {
		t.Fatalf("Fprint yaml: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "key: value") {
		t.Errorf("expected 'key: value' in yaml output, got %q", out)
	}
}

func TestPrintTableInvalidData(t *testing.T) {
	var buf bytes.Buffer
	err := Fprint(&buf, "table", "not a slice", nil)
	if err == nil {
		t.Fatal("expected error for invalid table data")
	}
}

func TestPrintUnsupportedFormat(t *testing.T) {
	var buf bytes.Buffer
	err := Fprint(&buf, "xml", nil, nil)
	if err == nil {
		t.Fatal("expected error for unsupported format")
	}
}

func TestPrintDefaultFormat(t *testing.T) {
	var buf bytes.Buffer
	rows := [][]string{{"a", "b"}}
	if err := Fprint(&buf, "", rows, nil); err != nil {
		t.Fatalf("Fprint empty format: %v", err)
	}
	if !strings.Contains(buf.String(), "a") {
		t.Error("expected output with empty format (defaults to table)")
	}
}
