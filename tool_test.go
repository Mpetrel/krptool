package krptool

import (
	"testing"
)

func TestFromPath(t *testing.T) {
	krpano, err := FromPath("res/tour.xml")
	if err != nil {
		t.Errorf("unexpected result: %v", err)
	}
	if len(krpano.Scene) != 19 {
		t.Errorf("invalid scene count: %d, expected: 19", len(krpano.Scene))
	}
}
