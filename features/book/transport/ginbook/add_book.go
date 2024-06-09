package ginbook

import (
	"net/http"

	"github.com/0xThomas3000/bookstore_api/component/appcontext"
	"github.com/0xThomas3000/bookstore_api/core"
	bookBusiness "github.com/0xThomas3000/bookstore_api/features/book/business"
	bookEntity "github.com/0xThomas3000/bookstore_api/features/book/entities"
	bookStorage "github.com/0xThomas3000/bookstore_api/features/book/storage"
	"github.com/gin-gonic/gin"
)

func AddBook(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data bookEntity.BookAdd

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := bookStorage.NewSQLStore(db)
		business := bookBusiness.NewAddBookBusiness(store)

		if err := business.AddBook(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusCreated, core.SimpleSuccessResponse(data.FakeId.String()))
	}
}
