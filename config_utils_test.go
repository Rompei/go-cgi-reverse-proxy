package main

import (
	"testing"
)

func TestLoadCondig(t *testing.T) {
	config, err := loadConfig("config/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config)
	assertArray(t, config.Path, "")
}

func assertArray(t *testing.T, path []interface{}, root string) {
	current := ""
	for i := range path {
		switch p := path[i].(type) {
		case string:
			current = p
			pp := root + "/" + p
			t.Log(pp)
		case []interface{}:
			assertArray(t, p, root+"/"+current)
		}
	}
}
