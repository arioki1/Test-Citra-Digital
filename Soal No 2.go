package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
)

func main() {
	router := gin.Default()
	router.GET("/:string", sortTextApi)
	router.Run(":3000")
}

func sortTextApi(c *gin.Context) {
	var hidup []string
	var mati []string
	inputParam := c.Param("string")
	splitString := strings.Split(inputParam, "")

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

	result := gin.H{
		"input": inputParam,
		"hasil": strings.Join(hidup, "") + strings.Join(mati, ""),
	}
	c.JSON(http.StatusOK, result)

}
