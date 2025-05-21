package main

import "testing"

func TestHello(t *testing.T) {
    got := "Hello, user!"
    want := "Hello, user!"

    if got != want {
        t.Errorf("Expected %s, got %s", want, got)
    }
}
