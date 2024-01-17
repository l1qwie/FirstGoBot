package client

import "fmt"

func FormulationOfTheAnswer(str string, name string) string {
	return fmt.Sprintf("%s, %s", str, name)
}

func ComStart(com string, name *string) (answer string) {
	if com == "/start" {
		answer = fmt.Sprint("Hello, World! Hello, ", *name)
	} else {
		answer = fmt.Sprint("Sorry, I couldn't understand you, ", *name)
	}
	return answer
}
