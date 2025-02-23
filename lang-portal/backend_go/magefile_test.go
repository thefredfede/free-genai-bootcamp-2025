package main

import (
	"os/exec"
	"testing"
)

func TestBuild(t *testing.T) {
	cmd := exec.Command("mage", "build")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Build failed: %s\nOutput: %s", err, string(output))
	}
	t.Logf("Build output: %s", string(output))
}

func TestRun(t *testing.T) {
	cmd := exec.Command("mage", "run")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Run failed: %s\nOutput: %s", err, string(output))
	}
	t.Logf("Run output: %s", string(output))
}

func TestMigrate(t *testing.T) {
	cmd := exec.Command("mage", "migrate")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Migrate failed: %s\nOutput: %s", err, string(output))
	}
	t.Logf("Migrate output: %s", string(output))
}

func TestSeed(t *testing.T) {
	cmd := exec.Command("mage", "seed")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Seed failed: %s\nOutput: %s", err, string(output))
	}
	t.Logf("Seed output: %s", string(output))
}
