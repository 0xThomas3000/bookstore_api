package ginbook

import (
	"net/http"

	"github.com/0xThomas3000/bookstore_api/component/appcontext"
	"github.com/0xThomas3000/bookstore_api/core"
	bookBusiness "github.com/0xThomas3000/bookstore_api/features/book/business"
	bookStorage "github.com/0xThomas3000/bookstore_api/features/book/storage"
	"github.com/gin-gonic/gin"
)

func ListBook(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData core.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(core.ErrBadRequest(err))
		}

		pagingData.Fulfill()

		store := bookStorage.NewSQLStore(db)
		business := bookBusiness.NewListBookBusiness(store)

		result, err := business.ListBook(c.Request.Context(), &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, core.NewSuccessResponse(result, pagingData, ""))
	}
}
