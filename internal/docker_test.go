package internal

import "testing"

func TestDocker_GetLinks(t *testing.T) {
	docker, err := NewDockerIntegration()
	if err != nil {
		t.Fatalf("Error NewDockerIntegration: %s", err)
	}
	links, err := docker.GetLinks()
	if err != nil {
		t.Fatalf("Error GetLinks: %s", err)
	}
	t.Log("List of link in docker")
	for _, link := range links {
		t.Logf("[%s - %s] %s", link.Icon, link.Name, link.Url)
	}
}
