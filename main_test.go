package main

import "testing"
//testing high severity
//test 2
//test 3
//test 4
//test 5
//test 6
//test 7
//test 8
//test 9
//test 10
//tyestt 11
func TestForceFailure(t *testing.T) {
    t.Fatal("Panic: Forced CI failure to test high severity")
}
