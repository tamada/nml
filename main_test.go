package main

import "testing"

func Test_main(t *testing.T) {
	status := goMain([]string{"nml", "-h"})
	if status != 0 {
		t.Errorf("status code wont 0, but got %d", status)
	}
}
