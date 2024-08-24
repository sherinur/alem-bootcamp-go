package bootcamp

func TextToMorse(s string) string {
	morse := ""
	morseCodeMap := map[rune]string{
		'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---", 'P': ".--.",
		'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--", 'Z': "--..", '1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....",
		'6': "-....", '7': "--...", '8': "---..", '9': "----.", '0': "-----", '.': ".-.-.-", ',': "--..--", '?': "..--..",
	}
	var char rune
	for i := 0; i < len(s); i++ {
		char = rune(s[i])

		if char >= 'a' && char <= 'z' {
			char -= 32
		}

		for key, value := range morseCodeMap {
			if rune(char) == key {
				morse += value + " "
			}
		}
	}

	return morse[:len(morse)-1]
}

// func main() {
// 	fmt.Println(TextToMorse("SOS"))                 // ... --- ...
// 	fmt.Println(TextToMorse("salem, how are you?")) // ... .- .-.. . -- --..-- .... --- .-- .- .-. . -.-- --- ..- ..--..
// }
