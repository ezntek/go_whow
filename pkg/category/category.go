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
	color "github.com/fatih/color"
)

type Category struct {
	name  string      `toml:"name"`
	color color.Color `toml:"color"`
}

func (c Category) PrettyPrint() {
	fmt.Printf("%s", c.color.Sprintf(" %s ", c.name))
}

type CategoryTree struct {
	categories []Category `toml:"categories"`
}

func NewCategory(name string, color color.Color) *Category {
	return &Category{
		name,
		color,
	}
}

func (ct CategoryTree) GetCategoryFromName(name string) *Category {
	for _, category := range ct.categories {
		if strings.EqualFold(category.name, name) {
			return &category
		}
	}
	return NewCategory(name, *color.New(color.BgWhite))
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
