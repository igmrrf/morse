package morse

import (
	"errors"
	"fmt"
)

func InterpretS(sentence string) (string, error) {
	if sentence == "" {
		return "", errors.New("word required")
	}
	morse_code, err := GetCodes()
	if err != nil {
		return "", err
	}

	coded := ""

	for _, char := range sentence {
		letter := fmt.Sprintf("%c", char)
		code := morse_code[letter]
		coded = coded + code + " "

	}
	return coded, nil
}

func InterpretW(word string) (string, error) {
	if word == "" {
		return "", errors.New("word required")
	}
	morse_code, err := GetCodes()
	if err != nil {
		return "", err
	}

	coded := ""
	for _, char := range word {
		letter := fmt.Sprintf("%c", char)
		code := morse_code[letter]
		coded = coded + code + " "
	}
	return coded, nil
}

func GetCodes() (map[string]string, error) {
	morse_codes := map[string]string{
		"a":  "•–",
		"b":  "–••",
		"c":  "-•-•",
		"d":  "–••",
		"e":  "•",
		"f":  "••-•",
		"g":  "--•",
		"h":  "••••",
		"i":  "••",
		"j":  "•---",
		"k":  "-•-",
		"l":  "•-••",
		"m":  "--",
		"n":  "-•",
		"o":  "---",
		"p":  "•--•",
		"q":  "--•-",
		"r":  "•-•",
		"s":  "•••",
		"t":  "-",
		"u":  "••-",
		"v":  "•••-",
		"w":  "•--",
		"x":  "-••-",
		"y":  "-•--",
		"z":  "--••",
		"A":  "•–",
		"B":  "–••",
		"C":  "-•-•",
		"D":  "–••",
		"E":  "•",
		"F":  "••-•",
		"G":  "--•",
		"H":  "••••",
		"I":  "••",
		"J":  "•---",
		"K":  "-•-",
		"L":  "•-••",
		"M":  "--",
		"N":  "-•",
		"O":  "---",
		"P":  "•--•",
		"Q":  "--•-",
		"R":  "•-•",
		"S":  "•••",
		"T":  "-",
		"U":  "••-",
		"V":  "•••-",
		"W":  "•--",
		"X":  "-••-",
		"Y":  "-•--",
		"Z":  "--••",
		"1":  ".----",
		"2":  "..---",
		"3":  "...--",
		"4":  "....-",
		"5":  ".....",
		"6":  "-....",
		"7":  "--...",
		"8":  "---..",
		"9":  "----.",
		"0":  "-----",
		".":  ".-.-.-",  // period
		":":  "---...",  // colon
		",":  "--..--",  // comma
		";":  "-.-.-",   // semicolon
		"?":  "..--..",  // question
		"=":  "-...-",   // equals
		"'":  ".----.",  // apostrophe
		"/":  "-..-.",   // slash
		"!":  "-.-.--",  // exclamation
		"-":  "-....-",  // dash
		"_":  "..--.-",  // underline
		"\"": ".-..-.",  // quotation marks
		"(":  "-.--.",   // parenthesis (open)
		")":  "-.--.-",  // parenthesis (close)
		"()": "-.--.-",  // parentheses
		"$":  "...-..-", // dollar
		"&":  ".-...",   // ampersand
		"@":  ".--.-.",  // at
		"+":  ".-.-.",   // plus
		"Á":  ".--.-",   // A with acute accent
		"Ä":  ".-.-",    // A with diaeresis
		"É":  "..-..",   // E with acute accent
		"Ñ":  "--.--",   // N with tilde
		"Ö":  "---.",    // O with diaeresis
		"Ü":  "..--",    // U with diaeresis
		" ":  ".......", // word interval
	}
	return morse_codes, nil
}
