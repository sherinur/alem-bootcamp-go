package bootcamp

func MorseToText(s string) string {
	morseCodeMap := map[rune]string{
		'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---", 'P': ".--.",
		'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--", 'Z': "--..", '1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....",
		'6': "-....", '7': "--...", '8': "---..", '9': "----.", '0': "-----", '.': ".-.-.-", ',': "--..--", '?': "..--..",
	}

	text := ""
	var morse []string
	code := ""
	var char rune

	for i := 0; i < len(s); i++ {
		char = rune(s[i])
		if char == ' ' {
			morse = append(morse, code)
			code = ""
			continue
		}

		code += string(char)
	}
	if code != "" {
		morse = append(morse, code)
	}

	for _, item := range morse {
		for key, value := range morseCodeMap {
			if value == item {
				text += string(key)
				break
			}
		}
	}

	return text
}

// func main() {
// 	fmt.Println(MorseToText("... --- ..."))                                                       // SOS
// 	fmt.Println(MorseToText("... .- .-.. . -- --..-- .... --- .-- .- .-. . -.-- --- ..- ..--..")) // SALEM,HOWAREYOU?
// }
