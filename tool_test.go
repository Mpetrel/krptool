package krptool

import (
	"encoding/xml"
	"fmt"
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
	// to xml
	data, err := xml.Marshal(krpano.Scene[0])
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
}
