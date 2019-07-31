package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortText(text string) {
	var hidup []string
	var mati []string
	splitString := strings.Split(text, "")

	for i := 0; i < len(splitString); i++ {
		if "a" == splitString[i] || "i" == splitString[i] || "u" == splitString[i] ||
			"e" == splitString[i] || "o" == splitString[i] {
			hidup = append(hidup, splitString[i])
		} else if " " != splitString[i] {
			mati = append(mati, splitString[i])
		}
	}

	sort.Strings(hidup)
	sort.Strings(mati)
	fmt.Println("Hasil : " + strings.Join(hidup, "") + strings.Join(mati, ""))

}
func main() {
	sortText("osama")
}
