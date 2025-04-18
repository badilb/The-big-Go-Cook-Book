package handling_errors

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

/*
Когда у тебя много функций, и все из них нужно обработать на ошибки,
сталшиваешься с такой фигней-зачастую обработка этих ошибок однообразна
и повторяется. Это вызывает огромное количество посторяющихся строк кода,
что не круто. Есть классное решение. Сделай функции-хелперы(helpers),
которые ты можешь использовать при обработке ошибок.

Например:
*/

type Person struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	SkinColor string `json:"skin_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
}

func unmarshal() (person Person) {
	r, err := http.Get("https://swapi.dev/api/people/1")
	if err != nil {
		// handle error
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		// handle error
	}
	err = json.Unmarshal(data, &person)
	if err != nil {
		// handle error
	}
	return person
}

/*
в функции unmarshal есть целых три обработки ошибок-когда мы вызываем
API-запрос http.Get, когда мы вызываем io.ReadAll и когда мы хотим
десериализовать JSON текст в структуру Person.

И все эти error-handling одинаковые. Но теперь посмотрим на код ниже:
*/

func helperUnmarshal() (person Person) {
	r, err := http.Get("https://swapi.dev/api/people/1")
	check(err, "Calling people API")
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	check(err, "Read JSON from response")

	err = json.Unmarshal(data, &person)
	check(err, "Unmarshalling")
	return person
}

func check(err error, msg string) {
	if err != nil {
		log.Println("Error encountered", msg)
		// do common error-handling stuff
	}
}
