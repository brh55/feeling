// Package feeling provides terminal output of feelings through text emoticons.
package feeling

import (
    "github.com/fatih/color"
)

// Uncertain appends (,,◕　⋏　◕,,) in blue
func Uncertain(text string) (string) {
    text += color.BlueString(" (,,◕　⋏　◕,,)")
    return text
}

// Whatever appends ¯\\_(ツ)_/¯ in green
func Whatever(text string) (string) {
     text += color.GreenString("¯\\_(ツ)_/¯")
     return text
}

// Scared appends ヽ(ﾟДﾟ)ﾉ in yellow
func Scared(text string) (string) {
    text += color.YellowString(" ヽ(ﾟДﾟ)ﾉ")
    return text
}

// Upset appends (╯°□°）╯︵ ┻━┻ in cyan
func Upset(text string) (string) {
    text += color.CyanString(" (╯°□°）╯︵ ┻━┻")
    return text;
}

// Cheerful appends (๑´▿｀๑)♫•*¨*•.¸¸♪✧ in cyan
func Cheerful(text string) (string) {
    text += color.CyanString(" (๑´▿｀๑)♫•*¨*•.¸¸♪✧")
    return text;
}

// Loved appends (*＾3＾）/～♡ in red
func Loved(text string) (string) {
    text += color.RedString(" (*＾3＾）/～♡")
    return text;
}

// Stumped appends (；¬＿¬) in blue
func Stumped(text string) (string) {
    text += color.BlueString("(；¬＿¬)")
    return text;
}
