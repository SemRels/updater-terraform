package plugin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpdaterUpdateTerraformFile(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "updater-terraform-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "versions.tf")
	if err := os.WriteFile(file, []byte("module \"x\" {\n  version = \"1.2.3\"\n}\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := NewUpdater().Update(file, "1.3.0"); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), `version = "1.3.0"`) {
		t.Fatalf("updated file = %s", got)
	}
}

func TestUpdaterMissingFile(t *testing.T) {
	t.Parallel()

	err := NewUpdater().Update(filepath.Join(t.TempDir(), "versions.tf"), "1.3.0")
	if err == nil || !strings.Contains(err.Error(), "read") {
		t.Fatalf("expected read error, got %v", err)
	}
}

func TestUpdaterMissingVersion(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "updater-terraform-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "versions.tf")
	if err := os.WriteFile(file, []byte("terraform {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	err = NewUpdater().Update(file, "1.3.0")
	if err == nil || !strings.Contains(err.Error(), "version assignment not found") {
		t.Fatalf("expected version error, got %v", err)
	}
}
