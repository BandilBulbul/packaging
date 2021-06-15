package helloo

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. Dev.to is awesome" {
		t.Errorf("Greet() = %s; want ello GitHub Actions. Dev.to is awesome", result)
	}
}
