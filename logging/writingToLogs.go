package logging

import (
	"fmt"
	"log"
	"strconv"
)

/*
Существует три вида логов:
	Print - пишет лог в логгер
	Fatal - пишет лог в логгер и вызывает выход из системы когда код на выводе будет 1
	Panic - пишет в логгер и  вызывает panic

Каждый из этих видов может иметь рпзличные функции; например
y Print есть Printf для форматтирования, а Println добавляет новую строку
после написания
*/

func logPrintln() {
	str := "awdwada"
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Println("Cannot parse string", err)
	}
	fmt.Println("Number is,", num)
}

func logFatal() {
	str := "awdwada"
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatalln("Cannot parse string", err)
	}
	fmt.Println("Number is,", num)
}

func logPanic() {
	str := "awddadawd"
	num := conv(str)
	fmt.Println("Number is", num)
}

func conv(str string) (num int64) {
	defer fmt.Println("deferred code in function conv")
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Panicln("Cannot parse string", err)
	}
	fmt.Println("End of function conv")
	return
}
