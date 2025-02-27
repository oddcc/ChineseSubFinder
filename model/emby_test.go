package model

import (
	"github.com/allanpk716/ChineseSubFinder/common"
	"testing"
)

func init() {
	configViper, err := InitConfigure()
	if err != nil {
		return
	}
	config, err = ReadConfig(configViper)
	if err != nil {
		return
	}
}

func TestEmbyHelper_GetRecentlyItems(t *testing.T) {

	em := NewEmbyHelper(config.EmbyConfig)
	items, err := em.GetRecentlyItems()
	if err != nil {
		t.Fatal(err)
	}

	println(items.Items[0].Name, items.Items[0].SeriesName, items.Items[0].Type)
}

func TestEmbyHelper_GetItemsAncestors(t *testing.T) {
	em := NewEmbyHelper(config.EmbyConfig)
	items, err := em.GetItemAncestors("96564")
	if err != nil {
		t.Fatal(err)
	}

	if len(items) < 1 {
		t.Fatal("less than 1")
	}

	println(items[0].Name, items[0].Path)
}

func TestEmbyHelper_GetItemVideoInfo(t *testing.T) {
	em := NewEmbyHelper(config.EmbyConfig)
	// 95813 -- 命运夜
	// 96564 -- The Bad Batch - S01E11
	videoInfo, err := em.GetItemVideoInfo("95813")
	if err != nil {
		t.Fatal(err)
	}

	println(videoInfo.Name, videoInfo.Path)
}

func TestEmbyHelper_UpdateVideoSubList(t *testing.T) {
	em := NewEmbyHelper(config.EmbyConfig)
	// 95813 -- 命运夜
	// 96564 -- The Bad Batch - S01E11
	err := em.UpdateVideoSubList("95813")
	if err != nil {
		t.Fatal(err)
	}
}

var (
	config *common.Config
)