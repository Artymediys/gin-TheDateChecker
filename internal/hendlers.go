package internal

import (
	"errors"
	"fmt"
	"gin_TheDateChecker/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetDays(context *gin.Context) {
	var yearSTR string = context.Param("year")

	year, err := strconv.Atoi(yearSTR)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusNotFound, "404 page not found")
		return
	} else if year < 0 {
		err = errors.New("only the our era needed")

		fmt.Println(err)
		context.String(http.StatusNotFound, err.Error())
		return
	}

	context.String(http.StatusOK, pkg.CalcDays(year))
}
