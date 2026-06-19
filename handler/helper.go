package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type Pagination struct {
	Limit  uint
	Offset uint
}

func GetUintContext(c fiber.Ctx, key string, def uint) uint {
	value, err := strconv.ParseUint(c.Query("limit", ""), 10, 64)
	if err != nil {
		return def
	}
	return uint(value)
}
func GetPaginationQuery(c fiber.Ctx) Pagination {
	limit := GetUintContext(c, "limit", 10)
	offset := GetUintContext(c, "offset", 0)

	return Pagination{
		Limit:  limit,
		Offset: offset,
	}
}
