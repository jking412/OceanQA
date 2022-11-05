package main

import (
	"OceanQA/boot"
	"testing"
)

func TestMain(m *testing.M) {
	boot.Initialize()
	m.Run()
}

func TestTemp(t *testing.T) {

}
