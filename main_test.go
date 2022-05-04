package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T){
  t.Log("Yeah Ici");
  fmt.Println("YEE AH");
}

func TestHello(t *testing.T){
  t.Log("testing hello world")
  assert.Equal(t,"Hello World", hello())
}

