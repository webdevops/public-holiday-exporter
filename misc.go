package main

import "strconv"

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}
