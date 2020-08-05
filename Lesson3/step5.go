package Lesson3

import (
	"encoding/json"
	"errors"
	"os"
)

//TODO На стандартный ввод подаются данные о студентах университетской группы в формате JSON (goto *1)
// В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
// Требуется прочитать данные и рассчитать среднее количество оценок, полученное студентами группы.
// Ответ на задачу требуется записать на стандартный вывод в формате JSON (goto *2).

type resultStruct struct {
	Average float32
}

type studentList struct {
	ID       uint
	Number   string
	Year     uint
	Students []student
}

type student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Bithday    string
	Address    string
	Phone      string
	Rating     []uint
}

func JSONWork(data []byte) []byte {
	var dataJSON studentList
	var result, count float32
	if !json.Valid(data) {
		panic(errors.New("Invalid JSON!"))
	}
	err := json.Unmarshal(data, &dataJSON)
	if err != nil {
		panic(err)
	}
	for _, value := range dataJSON.Students {
		result += float32(len(value.Rating))
		count++
	}
	resultBytes := resultStruct{result / count}
	resultJSON, err := json.MarshalIndent(resultBytes, " ", "\t")
	if err != nil {
		panic(err)
	}
	return resultJSON
}

//*1
// {
////    "ID":134,
////    "Number":"ИЛМ-1274",
////    "Year":2,
////    "Students":[
////        {
////            "LastName":"Вещий",
////            "FirstName":"Лифон",
////            "MiddleName":"Вениаминович",
////            "Birthday":"4апреля1970года",
////            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
////            "Phone":"+7(948)709-47-24",
////            "Rating":[1,2,3]
////        },
////        {
////            // ...
////        }
////    ]
//// }

//*2
//{
////    "Average": 14.1
//// }

//TODO На стандартный ввод подаются данные о студентах университетской группы в формате JSON:
// {
//    "ID":134,
//    "Number":"ИЛМ-1274",
//    "Year":2,
//    "Students":[
//        {
//            "LastName":"Вещий",
//            "FirstName":"Лифон",
//            "MiddleName":"Вениаминович",
//            "Birthday":"4апреля1970года",
//            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//            "Phone":"+7(948)709-47-24",
//            "Rating":[1,2,3]
//        },
//        {
//            // ...
//        }
//     ]
// }
// В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
// Требуется прочитать данные и рассчитать среднее количество оценок, полученное студентами группы.
// Ответ на задачу требуется записать на стандартный вывод в формате JSON.
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
