// --------  DO NOT EDIT --------
// this is an autogenerated file

package main

import (
	"net/http"
	"time"
)

// Version autogenerated
var Version = "2015-06-11T04:16:07.806453642Z"

// getVersion -
func getVersion() string {
	t, _ := time.Parse(time.RFC3339Nano, Version)
	if t.IsZero() {
		return ""
	}
	return t.Format(http.TimeFormat)
}
