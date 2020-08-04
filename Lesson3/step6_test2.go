package Lesson3

import "time"

//TODO На стандартный ввод подается строковое представление даты и времени
// определенного события в следующем формате: 2020-05-15 08:00:00
// Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату
// на стандартный вывод в том же формате. Если же событие должно произойти после обеда,
// необходимо перенести его на то же время на следующий день, а затем вывести на стандартный вывод в том же формате.
func RegisterBeforeLunch(scheduledDate *string) string {
	srcTime, err := time.Parse("2006-01-02 15:04:00", *scheduledDate)
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
