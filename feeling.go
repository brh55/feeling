/*
Package feeling provides terminal output of feelings through text emoticons.
*/

package feeling

import (
    "github.com/fatih/color"
)

// Appends (,,◕　⋏　◕,,) in blue
func Uncertain(text string) (string) {
    text += color.BlueString(" (,,◕　⋏　◕,,)")
    return text
}

// Appends ¯\\_(ツ)_/¯ in green
func Whatever(text string) (string) {
     text += color.GreenString("¯\\_(ツ)_/¯")
     return text
}

// Appends ヽ(ﾟДﾟ)ﾉ in yellow
func Scared(text string) (string) {
    text += color.YellowString(" ヽ(ﾟДﾟ)ﾉ")
    return text
}

// Appends (╯°□°）╯︵ ┻━┻ in cyan
func Upset(text string) (string) {
    text += color.CyanString(" (╯°□°）╯︵ ┻━┻")
    return text;
}

// Appends (๑´▿｀๑)♫•*¨*•.¸¸♪✧ in cyan
func Cheerful(text string) (string) {
    text += color.CyanString(" (๑´▿｀๑)♫•*¨*•.¸¸♪✧")
    return text;
}

// Appends (*＾3＾）/～♡ in red
func Loved(text string) (string) {
    text += color.RedString(" (*＾3＾）/～♡")
    return text;
}

// Appends (；¬＿¬) in blue
func Stumped(text string) (string) {
    text += color.BlueString("(；¬＿¬)")
    return text;
}
