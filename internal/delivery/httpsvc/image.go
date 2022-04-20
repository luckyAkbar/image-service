package httpsvc

import (
	"image-service/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) handleSave() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")

		if err != nil {
			logrus.Error(err)
			return ErrCustomMsgAndStatus(http.StatusBadRequest, err.Error())
		}

		imageHandler := usecase.NewImageHandler()
		res, err := imageHandler.HandleUpload(file)

		if err != nil {
			logrus.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}
