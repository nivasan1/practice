package crypto

import "fmt"

const cipher = "BZQODS BDLDMS ZEEZHQ FNZS"


func SolveCaesar() {
	for i := 0; i <= 26; i++ {
		fmt.Println(shiftChars(i))
	}
}

func shiftChars(i int) string {
	newCipher := make([]rune, 0)
	for _, char := range cipher {
		newCipher = append(newCipher, rune(int(char) + i))
	}
	return string(newCipher)
}
