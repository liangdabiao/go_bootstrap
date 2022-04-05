package category

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
	"gohub/pkg/types"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (category Category) {
	database.DB.Where("id", idstr).First(&category)
	return
}

// Get 通过 ID 获取用户
func GetOne(idstr string) (Category, error) {
	var category Category
	id := types.StringToUint64(idstr)
	if err := database.DB.First(&category, id).Error; err != nil {
		return category, err
	}
	return category, nil
}

func GetBy(field, value string) (category Category) {
	database.DB.Where("? = ?", field, value).First(&category)
	return
}

func All() (categories []Category) {
	database.DB.Find(&categories)
	return
}

// All 获取分类数据
func AllWeb() ([]Category, error) {
	var categories []Category
	if err := database.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Category{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Category{}),
		&categories,
		app.V1URL(database.TableName(&Category{})),
		perPage,
	)
	return
}
