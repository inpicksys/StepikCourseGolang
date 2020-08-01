package Lesson3

import (
	"encoding/json"
	"testing"
)

func TestJSONWork(t *testing.T) {
	dataRaw := []byte(`{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
        },
        {
            "LastName":"Вей",
            "FirstName":"Лион",
            "MiddleName":"Вениминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3,4]
        }
    ]
}`)
	result := JSONWork(dataRaw)
	var resultJSON resultStruct
	if err := json.Unmarshal(result, &resultJSON); err != nil {
		t.Error(err)
	}
	if resultJSON.Average != 3.5 {
		t.Error("Wrong answer! Expected 3.5, got ", resultJSON.Average)
	}
}
