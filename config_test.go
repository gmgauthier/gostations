package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//str2int
func Test_str2int(t *testing.T){
	t.Run("Testing normal integer:", teststr2intFunc("5", 5))
	t.Run("Testing large integer:", teststr2intFunc("4294967296",4294967296))
	t.Run("Testing negative integer:", teststr2intFunc("-5", -5))
	t.Run("Testing invalid input:", teststr2intFunc("a", 9999))
}

func teststr2intFunc(numstr string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		if !assert.Equal(t, expected, str2int(numstr)) {
			t.Error(fmt.Sprintf("%d does not equal %d", str2int(numstr), expected))
		}
	}
}

func Test_api(t *testing.T) {
	expectedApi := "all.api.radio-browser.info"
	api := api()
	pass := assert.Equal(t, expectedApi, api)
	if !pass {
		t.Error(fmt.Sprintf("`%s does not match %s", api, expectedApi))
	}
}

func Test_player(t *testing.T) {
	expectedPlayer := "mpv"
	player := player()
	pass := assert.Equal(t, expectedPlayer, player)
	if !pass {
		t.Error(fmt.Sprintf("`%s does not match %s", player, expectedPlayer))
	}
}

func Test_options(t *testing.T) {
	expectedOptions := "--no-video"
	options := options()
	pass := assert.Equal(t, expectedOptions, options)
	if !pass {
		t.Error(fmt.Sprintf("`%s does not match %s", options, expectedOptions))
	}
}

func Test_maxitems(t *testing.T) {
	expectedMaxItems := 9999
	maxItems := maxitems()
	pass := assert.Equal(t, expectedMaxItems, maxItems)
	if !pass {
		t.Error(fmt.Sprintf("`%d does not match %d", maxItems, expectedMaxItems))
	}
}