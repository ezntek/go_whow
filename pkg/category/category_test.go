package category_test

import (
	"testing"

	category "github.com/ezntek/whow/pkg/category"
	color "github.com/fatih/color"
)

func TestCategoryData(t *testing.T) {
	category := category.NewCategory("MyTestCategory", color.New(color.BgRed))
	t.Log("\033[1mTESTTT\033[0m")
	t.Log(category.SPrettyPrint())
}
