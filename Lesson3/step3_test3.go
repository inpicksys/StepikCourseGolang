package Lesson3

// TODO реализуем объект, удовлетворяющий интерфейсу fmt.Stringer. Назовем его "Батарейка".
//  Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.
//  Ваш тип должен предусматривать, что на печати он будет выглядеть так:
//  [      XXXX]: где пробелы - "опустошенная" емкость батареи, а X - "заряженная".
//  Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр:
//  0 или 1 (порядок 0/1 случайный). Ваша задача считать эту строку любым возможным
//  способом и создать на основе этой строки объект объявленного вами на первом этапе типа:
//  надеюсь, вы понимаете, что строка символизирует емкость батарейки: 0 - это "опустошенная"
//  часть, а 1 - "заряженная". В-третьих, созданный вами объект должен называться
//  batteryForTest (использование этого имени обязательно).
type Battery struct {
	capacity string
}

func (b Battery) String() string {
	return b.capacity
}

func InterfaceWork2(capacityRaw string) string {
	maker := func(cap string) string {
		lenCap := len(cap)
		counter := 0
		for _, i := range cap {
			if i == 49 {
				counter++
			}
		}
		result := make([]rune, lenCap+2)
		result[0] = '['
		result[lenCap+1] = ']'
		for i := 1; i < lenCap+1; i++ {
			if i <= lenCap-counter {
				result[i] = ' '
			} else {
				result[i] = 'X'
			}
		}
		return string(result)
	}
	batteryForTest := Battery{maker(capacityRaw)}
	return batteryForTest.capacity
}
