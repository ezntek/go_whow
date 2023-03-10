/*
 * Copyright 2023 ezntek (ezntek@xflymusic.com) and other contributors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package category

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	color "github.com/ezntek/whow/internal/color"
)

type Category struct {
	name  string      `toml:"name"`
	color color.Color `toml:"color"`
}

func (c Category) SPrettyPrint() string {
	return fmt.Sprintf("%s%s\n", c.color.Sprint("", "fore"), c.color.Sprintf(" %s ", "back", c.name))
}

func (c Category) PrettyPrint() {
	fmt.Print(c.SPrettyPrint())
}

type CategoryTree struct {
	categories []Category `toml:"categories"`
}

func NewCategory(name string, clr color.Color) *Category {
	clr.MakeBold()
	return &Category{
		name,
		clr,
	}
}

func (ct CategoryTree) GetCategoryFromName(name string) *Category {
	for _, category := range ct.categories {
		if strings.EqualFold(category.name, name) {
			return &category
		}
	}
	return NewCategory(name, color.New(color.White))
}

func NewCategoryTree() *CategoryTree {
	retval := CategoryTree{}

	tomlFileText, err := os.ReadFile(path.Join(os.Getenv("HOME"), "./.local/whow/categories.toml"))
	if err != nil {
		panic(err)
	}

	if _, err := toml.Decode(string(tomlFileText), retval); err != nil {
		panic(err)
	}

	return &retval
}
