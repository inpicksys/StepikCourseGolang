package Lesson3

import (
	"encoding/json"
	"os"
)

type MyStruct struct {
	MinItems    uint   `json:"-"`
	Scheme      string `json:"-"`
	Description string `json:"-"`
	Title       string `json:"-"`
	Type        string `json:"-"`
	Items       []ItemsInOKVED
}

type ItemsInOKVED struct {
	Description string `json:"-"`
	Type        string `json:"-"`
	Properties  Props
}

type Props struct {
	system_object_id string `json:"-"`
	Kod              string `json:"-"`
	is_deleted       int    `json:",omitempty"`
	signature_day    string `json:"-"`
	Nomdescr         string `json:"-"`
	global_id        int64  `json:"global_id"`
	Idx              string `json:",omitempty"`
	Razdel           string `json:"-"`
	Name             string `json:"-"`
}

func EncodeFromFileDecodeToFileCountSumAllGlobalIDs(srcFileName string) int64 {
	reader, err := os.Open(srcFileName)
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(reader)
	var resultObject MyStruct
	enc.Encode(resultObject)
	return 0
}
