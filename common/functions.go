package common

import (
	"github.com/labstack/echo"
	"net/http"
)

func SystemResponse(ec echo.Context, data interface{}, err error) error {
	if err == nil {
		return ec.JSON(http.StatusOK, &BaseResponseDto{
			Code: http.StatusOK,
			Data: data,
		})
	}
	return ec.JSON(http.StatusInternalServerError, &BaseResponseDto{
		Code:  http.StatusInternalServerError,
		Data:  data,
		Error: err.Error(),
	})
}
