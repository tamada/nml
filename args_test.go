package main

import "testing"

func TestParseArguments(t *testing.T) {
	testdata := []struct {
		giveArgs       []string
		wontHeader     bool
		wontBackground bool
		wontHelp       bool
		wontError      bool
		message        string
	}{
		{[]string{"nml"}, false, false, false, true, "NUMBER is mandatory"},
		{[]string{"nml", "not_a_number"}, false, false, false, true, "not a number"},
		{[]string{"nml", "10"}, false, false, false, false, "success"},
		{[]string{"nml", "-h"}, false, false, true, false, "success with help"},
		{[]string{"nml", "-h", "--background"}, false, false, true, false, "success with help"},
		{[]string{"nml", "-H", "10"}, true, false, false, false, "success with header"},
		{[]string{"nml", "10", "--", "echo", "hoge"}, false, false, false, false, "success with command"},
		{[]string{"nml", "-u", "unknown_unit", "10"}, false, false, false, true, "parsing fail: the unknown unit"},
		{[]string{"nml", "10", "--"}, false, false, false, true, "parsing fail: no commands given"},
		{[]string{"nml", "-unknown-flag"}, false, false, false, true, "parsing fail: the unknown flag"},
	}
	for _, td := range testdata {
		opts, err := parseArgs(td.giveArgs)
		if (err == nil) && td.wontError {
			t.Errorf("parseArgs(%v) wont error, but got no error: %s", td.giveArgs, td.message)
		}
		if err != nil && !td.wontError {
			t.Errorf("parseArgs(%v) wont no error, but got error: %s (%s)", td.giveArgs, err.Error(), td.message)
		}
		if err != nil {
			continue
		}
		if opts.background != td.wontBackground {
			t.Errorf("parseArgs(%v) background did not match, wont %v, but got %v", td.giveArgs, td.wontBackground, opts.background)
		}
		if opts.header != td.wontHeader {
			t.Errorf("parseArgs(%v) header did not match, wont %v, but got %v", td.giveArgs, td.wontHeader, opts.header)
		}
		if opts.help != td.wontHelp {
			t.Errorf("parseArgs(%v) help did not match, wont %v, but got %v", td.giveArgs, td.wontHelp, opts.help)
		}
	}
}
