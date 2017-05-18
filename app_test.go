package main

import "testing"

func TestRedisPing(t *testing.T) {
	if redisPingRoutine() != "PONG" {
		t.Error("Redis PING failed! not a PONG!")
	}
	t.Log("Redis PING Succeed! PING PONG.")
}
