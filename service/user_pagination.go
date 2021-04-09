package user

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/furuTra/gin_train/entity"
)

// Pagination条件作成
func (s Service) GeneratePaginationFromRequest(c *gin.Context) (entity.Pagination) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	sort := c.DefaultQuery("sort", "asc")

	return entity.Pagination {
		Limit: limit,
		Offset: offset,
		Sort: sort,
	}
}