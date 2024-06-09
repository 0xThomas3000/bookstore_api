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

func UpdateBook(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := core.FromBase58(c.Param("id"))

		if err != nil {
			panic(core.ErrBadRequest(err))
		}

		var data bookEntity.BookUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(core.ErrBadRequest(err))
		}

		store := bookStorage.NewSQLStore(db)
		business := bookBusiness.NewUpdateBookBusiness(store)

		if err := business.UpdateBook(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, core.SimpleSuccessResponse(true))
	}
}
