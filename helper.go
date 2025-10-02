package main

import (
	"reflect"
	"strings"
)

func old(name string, msg *Message) string {
	if name == "" {
		return ""
	}

	if msg == nil {
		return "" // or handle error
	}

	//fmt.Printf("%#v\n", msg) // raw Go representation
	//os.Exit(1)

	words := strings.Split(name, "_")

	// Validation
	if len(words) > 2 {
		errors = append(errors, "input can have at most two words separated by '_'")
	} else if len(words) == 2 && (words[0] == "" || words[1] == "") {
		errors = append(errors, "words must be non-empty")
	}

	// Capitalize each word
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}

	readableWords := strings.Join(words, "_")

	val := reflect.ValueOf(msg).Elem()
	field := val.FieldByName(readableWords)

	if field.IsValid() && field.Kind() == reflect.String {
		//println("\n" + field.String())
		return field.String() // return field value (e.g., msg.Name)
	}

	return ""
}
