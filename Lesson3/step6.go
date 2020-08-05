package Lesson3

import (
	"strings"
	"time"
)

const baseDate = 1589570165

//TODO На стандартный ввод подается строковое представление даты и времени в следующем формате:
// 1986-04-16T05:20:00+06:00
// Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:
func ConvertTimeToUnixFormat(dateAndTime string) string {
	srcTime, err := time.Parse(time.RFC3339, dateAndTime)
	if err != nil {
		panic(err)
	}
	return srcTime.Format(time.UnixDate)
}

//TODO На стандартный ввод подается строковое представление даты и времени
// определенного события в следующем формате: 2020-05-15 08:00:00
// Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату
// на стандартный вывод в том же формате. Если же событие должно произойти после обеда,
// необходимо перенести его на то же время на следующий день, а затем вывести на стандартный вывод в том же формате.
func RegisterBeforeLunch(scheduledDate *string) string {
	srcTime, err := time.Parse("2006-01-02 15:04:05", *scheduledDate)
	if err != nil {
		panic(err)
	}
	timeStamp := time.Date(srcTime.Year(), srcTime.Month(), srcTime.Day(), 13, 0, 0, 0, time.Local)
	if srcTime.After(timeStamp) {
		// добавить 12 часов
		return (srcTime.Add(time.Hour * 24)).Format("2006-01-02 15:04:00")
	}
	return *scheduledDate
}

//TODO На стандартный ввод подается строковое представление двух дат, разделенных запятой. Необходимо преобразовать
// полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.
func GetDifferenceBetweenDates(date1, date2 *string) string {
	time1, err := time.Parse("02.01.2006 15:04:05", *date1)
	if err != nil {
		panic(err)
	}
	time2, err := time.Parse("02.01.2006 15:04:05", *date2)
	if err != nil {
		panic(err)
	}
	if time1.After(time2) {
		dur, err := time.ParseDuration(time1.Sub(time2).String())
		if err != nil {
			panic(err)
		}
		return dur.String()
	}
	dur, err := time.ParseDuration(time2.Sub(time1).String())
	if err != nil {
		panic(err)
	}
	return dur.String()
}

//TODO На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере).
// Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей
// преобразования равны 0). Требуется считать данные о продолжительности периода, преобразовать их в тип Duration,
// а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.
func UnixDateFromBaseDatePlusDuration(durS *string) string {
	durCorrect := strings.ReplaceAll(*durS, " час. ", "h")
	durCorrect = strings.ReplaceAll(durCorrect, " мин. ", "m")
	durCorrect = strings.ReplaceAll(durCorrect, " сек.", "s")
	dur, err := time.ParseDuration(durCorrect)
	if err != nil {
		panic(err)
	}
	standardDate := time.Unix(baseDate, 0).UTC()
	result := standardDate.Add(dur)
	return result.Format(time.UnixDate)
}
