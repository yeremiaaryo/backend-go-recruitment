package movie

import (
	"backend-go-recruitment/common"
	"github.com/labstack/echo"
	"strings"
)

func (mc *MovieController) FindMovieByKeyword(ec echo.Context) error {
	searchWord := ec.QueryParam("searchword")
	pagination := ec.QueryParam("pagination")
	if strings.EqualFold(pagination, "") {
		pagination = "1"
	}
	movieList, err := mc.MovieInterface.FindMovieByKeyword(searchWord, pagination)
	if err != nil {
		return err
	}
	return common.SystemResponse(ec, movieList, err)
}
