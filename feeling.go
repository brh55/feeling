package feeling

import (
    "github.com/fatih/color"
)

/**
 * Express your feelings into the strings
 */

func Uncertain(text string) (string) {
    text += color.BlueString(" (,,◕　⋏　◕,,)")
    return text
}

func Whatever(text string) (string) {
     text += color.GreenString("¯\\_(ツ)_/¯")
     return text
}

func Scared(text string) (string) {
    text += color.YellowString("ヽ(ﾟДﾟ)ﾉ")
    return text
}

func Upset(text string) (string) {
    text += color.CyanString("(╯°□°）╯︵ ┻━┻")
    return text;
}

func Cheerful(text string) (string) {
    text += color.CyanString("(๑´▿｀๑)♫•*¨*•.¸¸♪✧")
    return text;
}

func Loved(text string) (string) {
    text += color.RedString("(*＾3＾）/～♡")
    return text;
}

func Stumped(text string) (string) {
    text += color.BlueString("(；¬＿¬)")
    return text;
}
