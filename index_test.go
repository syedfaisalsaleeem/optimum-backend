package main

import (
     "testing"	 
)

var a App

func TestNew(t *testing.T) {
    a.Initialize(
		"postgres",
		"postgres",
		"testing")
}

func TestFail(t *testing.T) {
	 t.Errorf("Expected error")

}
