package ginbook

import (
	"net/http"

	"github.com/0xThomas3000/bookstore_api/component/appcontext"
	"github.com/0xThomas3000/bookstore_api/core"
	bookBusiness "github.com/0xThomas3000/bookstore_api/features/book/business"
	bookEntity "github.com/0xThomas3000/bookstore_api/features/book/entities"
	bookStore "github.com/0xThomas3000/bookstore_api/features/book/storage"
	"github.com/gin-gonic/gin"
)

func DeleteBook(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := core.FromBase58(c.Param("id"))

		if err != nil {
			panic(core.ErrBadRequest(err))
		}

		var data bookEntity.Book

		if err := c.ShouldBind(&data); err != nil {
			panic(core.ErrBadRequest(err))
		}

		store := bookStore.NewSQLStore(db)
		business := bookBusiness.NewDeleteBookBusiness(store)

		if err := business.DeleteBook(
			c.Request.Context(),
			int(uid.GetLocalID()),
			&data,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusNoContent, core.SimpleSuccessResponse(nil))
	}
}
