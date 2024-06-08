package ginupload

import (
	"fmt"
	"net/http"

	"github.com/0xThomas3000/bookstore_api/component/appcontext"
	"github.com/0xThomas3000/bookstore_api/core"
	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(err)
		}

		// dst: tên của folder & file
		if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, core.SimpleSuccessResponse(core.Image{
			Id:        0,
			Url:       "http://localhost:8080/static/" + fileHeader.Filename,
			Width:     2119,
			Height:    1414,
			CloudName: "local",
			Extension: "jpeg",
		}))
	}
}
