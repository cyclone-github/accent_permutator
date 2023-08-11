package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// accent_permutator is a tool to transform characters to accented characters such as "o" to "ò"
// tool will accept ASCII, UTF-8 and $HEX[] formatted wordlists
// ex input: plaintext "password" or "$HEX[70617373776f7264]"
// ex usage: cat wordlist.txt | ./accent_permutator.bin | hashcat.bin...
// written by cyclone
// v2023-08-10.2330

var accentReplacements = map[rune][]rune{
	'a': {'á', 'à', 'â', 'ã', 'ä', 'å', 'ā', 'ă', 'α', 'ά', 'a'},
	'e': {'é', 'è', 'ê', 'ë', 'ę', 'ē', 'ė', 'ě', 'ё', 'е', 'e'},
	'i': {'í', 'ì', 'î', 'ï', 'į', 'ī', 'i'},
	'o': {'ó', 'ò', 'ô', 'õ', 'ö', 'ő', 'ō', 'o'},
	'u': {'ú', 'ù', 'û', 'ü', 'ų', 'ū', 'u'},
	'y': {'ý', 'ÿ', 'y'},
	'A': {'Á', 'À', 'Â', 'Ã', 'Ä', 'Å', 'Ā', 'Ă', 'A'},
	'E': {'É', 'È', 'Ê', 'Ë', 'Ę', 'Ē', 'Ė', 'Ě', 'E'},
	'I': {'Í', 'Ì', 'Î', 'Ï', 'Į', 'Ī', 'I'},
	'O': {'Ó', 'Ò', 'Ô', 'Õ', 'Ö', 'Ő', 'Ō', 'O'},
	'U': {'Ú', 'Ù', 'Û', 'Ü', 'Ų', 'Ū', 'U'},
	'Y': {'Ý', 'Ÿ', 'Y'},
	'c': {'ç', 'ć', 'č', 'c'},
	'C': {'Ç', 'Ć', 'Č', 'C'},
	's': {'š', 'ś', 'ş', 's'},
	'z': {'ž', 'ź', 'ż', 'Z'},
	'S': {'Š', 'Ś', 'Ş', 'S'},
	'Z': {'Ž', 'Ź', 'Ż', 'Z'},
	'l': {'ł', 'ļ', 'l'},
	'L': {'Ł', 'Ļ', 'L'},
	'n': {'ñ', 'ń', 'ň', 'й', 'и', 'n'},
	'N': {'Ñ', 'Ń', 'Ň', 'N'},
	'd': {'đ', 'ď', 'd'},
	'D': {'Đ', 'Ď', 'D'},
	't': {'ť', 'ţ', 't'},
	'T': {'Ť', 'Ţ', 'T'},
	'r': {'ř', 'r'},
	'R': {'Ř', 'R'},
	'g': {'ğ', 'g'},
	'G': {'Ğ', 'G'},
	'!': {'¡', '!'},
	'?': {'¿', '?'},
	'α': {'ά', 'α'},
	'ε': {'έ', 'ε'},
	'е': {'ё', 'е'},
	'и': {'й', 'и'},
	'ا': {'أ', 'إ', 'ا'},
	'א': {'אַ', 'א'},
	'长': {'長', '长'},
	'あ': {'ア', 'あ'},
	'ㅏ': {'ㅐ', 'ㅏ'},
}

func applyAccents(input string, pos int, result []rune) {
	if pos == len(input) {
		fmt.Println(string(result))
		return
	}

	char := rune(input[pos])
	if replacements, ok := accentReplacements[char]; ok {
		for _, r := range replacements {
			result[pos] = r
			applyAccents(input, pos+1, result)
		}
	} else {
		result[pos] = char
		applyAccents(input, pos+1, result)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var inputs []string
	for scanner.Scan() {
		inputLine := strings.TrimSpace(scanner.Text())

		// check if word is HEX
		if strings.HasPrefix(inputLine, "$HEX[") && strings.HasSuffix(inputLine, "]") {
			hexStr := strings.TrimPrefix(inputLine, "$HEX[")
			hexStr = strings.TrimSuffix(hexStr, "]")
			decoded, err := hex.DecodeString(hexStr)
			if err == nil {
				inputLine = string(decoded)
			}
		}

		inputs = append(inputs, inputLine)
	}

	for _, input := range inputs {
		result := make([]rune, len(input))
		applyAccents(input, 0, result)
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", scanner.Err())
	}
}

// end code
