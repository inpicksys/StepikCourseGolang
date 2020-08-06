package Lesson3

import "sync"

//TODO Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
func ChannelWork(channel chan int, n int) {
	channel <- n + 1
}

//TODO Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз,
// добавив к ней пробел.
func SendModifiedStringToChannel(channel chan string, str string) {
	str += " "
	for i := 0; i < 5; i++ {
		channel <- str
	}
}

//TODO Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения
// на следующий этап конвейера только если оно отличается от того, что пришло ранее. Ваша фукнция должна
// принимать два канала - inputStream и outputStream, в первый вы будете получать строки, во второй вы должны
// отправлять значения без повторов. В итоге в outputStream должны остаться значения, которые не повторяются подряд.
func RemoveDuplicates(inputStream chan string, outputStream chan string) {
	currentValue := ""
	defer close(outputStream)
	for {
		candidate, isOk := <-inputStream
		if !isOk {
			break
		}
		if candidate != currentValue {
			currentValue = candidate
		} else {
			continue
		}
		outputStream <- currentValue
	}
}

//TODO необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения.
// Функция work() ничего не принимает и не возвращает.
//Won't be tested due to the realization details
func SimulateWork() {
	done := make(chan struct{})
	go func(c chan struct{}) {
		work(5)
		close(c)
	}(done)
	<-done
}

//TODO Необходимо в отдельных горутинах вызвать функцию work 10 раз и дождаться результатов выполнения вызванных функций.
//Won't be tested due to the realization details
func SimulateWork2() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			work(5)
		}(wg)
	}
	wg.Wait()
}

//TODO Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.
// В случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал
// вы должны отправить квадрат аргумента. В случае, если аргумент будет получен из канала secondChan,
// в выходной (возвращенный) канал вы должны отправить результат умножения аргумента на 3. В случае,
// если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
// Функция calculator должна быть неблокирующей, сразу возвращая управление. Ваша функция получит всего одно значение
// в один из каналов - получили значение, обработали его, завершили работу. После завершения работы необходимо
// освободить ресурсы, закрыв выходной канал, если вы этого не сделаете, то превысите предельное время работы.
//Won't be tested due to the realization details
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	result := make(chan int)
	defer close(result)
	go func() {
		for {
			select {
			case data := <-firstChan:
				result <- data * data
			case data := <-secondChan:
				result <- data * 3
			case _ = <-stopChan:

			}
		}
	}()
	return result
}
