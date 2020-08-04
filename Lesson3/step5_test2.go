package Lesson3

import (
	"encoding/json"
	"os"
)

/*
type Props struct {
	SystemObjectId string `json:"system_object_id,omitempty"`
	Kod            string `json:"Kod,omitempty"`
	IsDeleted      string `json:"is_deleted,omitempty"`
	SignatureDay   string `json:"signature_day,omitempty"`
	Nomdescr       string `json:"Nomdescr,omitempty"`
	GlobalId       int64 `json:"global_id"`
	Idx            string `json:"Idx,omitempty"`
	Razdel         string `json:"Razdel,omitempty"`
	Name           string `json:"Name,omitempty"`
}
*/
type Props struct {
	SystemObjectId string `json:"-"`
	Kod            string `json:"-"`
	IsDeleted      string `json:"-"`
	SignatureDay   string `json:"-"`
	Nomdescr       string `json:"-"`
	GlobalId       int64  `json:"global_id"`
	Idx            string `json:"-"`
	Razdel         string `json:"-"`
	Name           string `json:"-"`
}

func EncodeFromFileDecodeToFileCountSumAllGlobalIDs(srcFileName string) int64 {
	var (
		reader, err  = os.Open(srcFileName)
		resultObject = []Props{}
		result       int64
	)
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(reader)
	dec.Decode(&resultObject)
	for _, val2 := range resultObject {
		result += val2.GlobalId
	}
	return result
}
