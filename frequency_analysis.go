package main

import (
	"fmt"
	"sort"
	"strings"
)

const counter = 10

type key_val struct {
	Key   string
	Value int
}

//Написать функцию, которая получает на вход текст и возвращает
//10 самых часто встречающихся слов без учета словоформ

func Analysis(s_in string) []key_val {

	dct := make(map[string]int)
	words := strings.Fields(s_in)
	for _, word := range words {
		dct[word] += 1
	}

	var ss []key_val
	for k, v := range dct {
		ss = append(ss, key_val{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	if len(ss) >= counter {
		return ss[:counter]
	} else {
		return ss
	}
}

func main() {

	someString := "привет всем сегодня мы будем говорить о том. как хорошо же жить на этом свете и о том, что лучше всего помогает нам выживать. Далее, вы все говорите привет и вам в ответ говорят привет."

	result := Analysis(someString)
	for i, r := range result {
		fmt.Println("i:", i, "Frequency:", r.Value, "Word:", r.Key)
	}
}
