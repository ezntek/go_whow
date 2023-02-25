package category_test

import (
	"testing"

	color "github.com/ezntek/whow/internal/color"
	category "github.com/ezntek/whow/pkg/category"
)

func TestCategoryData(t *testing.T) {
	clr := color.New(color.Red)
	category := category.NewCategory("MyTestCategory", clr)

	t.Log("TESTTT")
	t.Log(category.SPrettyPrint())
}
