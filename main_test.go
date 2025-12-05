package main

import "testing"
//testing high severity
//test 2
//test 3
//test 4
//test 5
//test 6
//test 7
func TestForceFailure(t *testing.T) {
    t.Fatal("fail: Forced CI failure to test high severity")
}
