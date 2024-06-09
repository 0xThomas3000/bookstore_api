package ginbook

import (
	"net/http"

	"github.com/0xThomas3000/bookstore_api/component/appcontext"
	"github.com/0xThomas3000/bookstore_api/core"
	bookBusiness "github.com/0xThomas3000/bookstore_api/features/book/business"
	bookStorage "github.com/0xThomas3000/bookstore_api/features/book/storage"
	"github.com/gin-gonic/gin"
)

func FindBook(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := core.FromBase58(c.Param("id"))

		if err != nil {
			panic(core.ErrBadRequest(err))
		}

		store := bookStorage.NewSQLStore(db)
		business := bookBusiness.NewFindBookBusiness(store)

		result, err := business.FindBook(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask()

		c.JSON(http.StatusOK, core.SimpleSuccessResponse(result))
	}
}
