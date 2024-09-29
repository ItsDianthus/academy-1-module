package infrastructure

import "fmt"

func WelcomeWord() {
	fmt.Print("Всем привет. Это мой первый проект на Golang. \nДобро пожаловать в Виселицу!\n")
}

func EndGameWriter(gameWin bool, word string) {
	if gameWin {
		fmt.Println("Вы победили! :)")
	} else {
		fmt.Println("Вы проиграли :(")
	}
	fmt.Println("Загаданное слово: ", word)
}

func HangmanWriter(tries int) {
	switch tries {
	case 11:
		fmt.Println("\n" +
			"                \n" +
			"                \n" +
			"                \n" +
			"                \n" +
			"                \n" +
			"     ███        \n" +
			"    ══════╩═══")
	case 10:
		fmt.Println("\n" +
			"                \n" +
			"                \n" +
			"                \n" +
			"          ║   \n" +
			"          ║   \n" +
			"     ███  ║   \n" +
			"    ══════╩═══")
	case 9:
		fmt.Println("\n" +
			"                \n" +
			"          ║   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"     ███  ║   \n" +
			"    ══════╩═══")
	case 8:
		fmt.Println("\n" +
			"          ══╗   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"     ███  ║   \n" +
			"    ══════╩═══")
	case 7:
		fmt.Println("\n" +
			"      ╔═══╗   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"     ███  ║   \n" +
			"    ══════╩═══")
	case 6:
		fmt.Println("\n" +
			"      ╔═══╗   \n" +
			"      |   ║   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"     ███  ║   \n" +
			"    ══════╩═══")
	case 5:
		fmt.Println("\n" +
			"      ╔═══╗   \n" +
			"      |   ║   \n" +
			"      O   ║   \n" +
			"          ║   \n" +
			"          ║   \n" +
			"     ███  ║   \n" +
			"    ══════╩═══")
	case 4:
		fmt.Println("\n" +
			"      ╔═══╗   \n" +
			"      |   ║   \n" +
			"      O   ║   \n" +
			"      |   ║   \n" +
			"          ║   \n" +
			"     ███  ║   \n" +
			"    ══════╩═══")
	case 3:
		fmt.Println("\n" +
			"      ╔═══╗   \n" +
			"      |   ║   \n" +
			"      O   ║   \n" +
			"      |\\  ║   \n" +
			"          ║   \n" +
			"     ███  ║   \n" +
			"    ══════╩═══")
	case 2:
		fmt.Println("\n" +
			"      ╔═══╗   \n" +
			"      |   ║   \n" +
			"      O   ║   \n" +
			"     /|\\  ║   \n" +
			"          ║   \n" +
			"     ███  ║   \n" +
			"    ══════╩═══")
	case 1:
		fmt.Println("\n" +
			"      ╔═══╗   \n" +
			"      |   ║   \n" +
			"      O   ║   \n" +
			"     /|\\  ║   \n" +
			"       \\  ║   \n" +
			"     ███  ║   \n" +
			"    ══════╩═══")
	case 0:
		fmt.Println("\n" +
			"      ╔═══╗   \n" +
			"      |   ║   \n" +
			"      O   ║   \n" +
			"     /|\\  ║   \n" +
			"     / \\  ║   \n" +
			"          ║   \n" +
			"    ══════╩═══")
	default:
		fmt.Println("Некорректное значение попыток.")
	}

}
