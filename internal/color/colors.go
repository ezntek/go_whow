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

package color

import (
	"fmt"
	"strings"

	clr "github.com/fatih/color"
)

// Interface Definitions
type IColor interface {
	Print(text string, mode string)
	Println(text string, mode string)
	Sprint(text string, mode string) string
	FromName(name string) interface{}
}

type colorFgBgPair struct {
	Fg *clr.Color
	Bg *clr.Color
}

// Structure Definitions
var (
	Black = colorFgBgPair{
		Fg: clr.New(clr.FgBlack),
		Bg: clr.New(clr.BgBlack),
	}

	Red = colorFgBgPair{
		Fg: clr.New(clr.FgRed),
		Bg: clr.New(clr.BgRed),
	}

	Green = colorFgBgPair{
		Fg: clr.New(clr.FgGreen),
		Bg: clr.New(clr.BgGreen),
	}

	Yellow = colorFgBgPair{
		Fg: clr.New(clr.FgYellow),
		Bg: clr.New(clr.BgYellow),
	}

	Blue = colorFgBgPair{
		Fg: clr.New(clr.FgBlue),
		Bg: clr.New(clr.BgBlue),
	}

	Magenta = colorFgBgPair{
		Fg: clr.New(clr.FgMagenta),
		Bg: clr.New(clr.BgMagenta),
	}

	Cyan = colorFgBgPair{
		Fg: clr.New(clr.FgCyan),
		Bg: clr.New(clr.BgCyan),
	}

	White = colorFgBgPair{
		Fg: clr.New(clr.FgWhite),
		Bg: clr.New(clr.BgWhite),
	}
)

type Color struct {
	Name      string
	ColorPair colorFgBgPair
}

func (color *Color) MakeBold() {
	color.ColorPair.Fg.Add(clr.Bold)
	color.ColorPair.Bg.Add(clr.Bold)
}

func (color Color) Sprint(text string, mode string) string {
	switch strings.ToLower(mode) {
	case "fore":
		return color.ColorPair.Fg.Sprint(text)
	case "back":
		return color.ColorPair.Bg.Sprint(text)
	default:
		return color.ColorPair.Fg.Sprint(text)
	}
}

func (color Color) Sprintf(format string, mode string, a ...interface{}) string {
	switch strings.ToLower(mode) {
	case "fore":
		return color.ColorPair.Fg.Sprintf(format, a...)
	case "back":
		return color.ColorPair.Bg.Sprintf(format, a...)
	default:
		return color.ColorPair.Fg.Sprintf(format, a...)
	}
}

func (color Color) Print(text string, mode string) {
	fmt.Print(color.Sprint(text, mode))
}

func (color Color) Println(text string, mode string) {
	fmt.Printf("%s\n", color.Sprint(text, mode))
}

func (color Color) Printf(text string, mode string, a ...interface{}) {
	fmt.Print(color.Sprintf(text, mode, a...))
}

func New(colorPair colorFgBgPair) Color {
	return Color{
		Name:      "black",
		ColorPair: colorPair,
	}
}
