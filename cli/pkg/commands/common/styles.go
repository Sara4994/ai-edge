/*
Copyright 2024. Open Data Hub Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import "github.com/charmbracelet/lipgloss"

var TableBaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#04B575"))

var MessageStyle = lipgloss.NewStyle().
	Bold(true)

var Success = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#04B575")).
	Bold(true)

var ErrorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF0000")).
	Bold(true).
	Height(4).
	Width(120)

var KeyStyle = lipgloss.NewStyle().Bold(true).Width(20)
var ParamKeyStyle = lipgloss.NewStyle().Width(40).MarginLeft(10)

var TitleStyle = lipgloss.NewStyle().Bold(true).Underline(true).PaddingTop(1)
