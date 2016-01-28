package feeling

import (
    "testing"
    "strings"
)

func TestFeeling(t *testing.T) {
    testCases := []struct {
        method, emoji string
    }{
        {"Scared", "ヽ(ﾟДﾟ)ﾉ"},
        {"Whatever", "¯\\_(ツ)_/¯"},
        {"Upset", "(╯°□°）╯︵ ┻━┻"},
        {"Uncertain", "(,,◕　⋏　◕,,)"},
        {"Cheerful", "(๑´▿｀๑)♫•*¨*•.¸¸♪✧"},
        {"Loved", "(*＾3＾）/～♡"},
        {"Stumped", "(；¬＿¬)"},
    }
    for _, c := range testCases {
        switch c.method {
            case "Scared":
                output := Scared("test");
                logHelper(output, c.emoji, t)

            case "Whatever":
                output := Whatever("test")
                logHelper(output, c.emoji, t)

            case "Upset":
                output := Upset("test")
                logHelper(output, c.emoji, t)

            case "Uncertain":
                output := Uncertain("test")
                logHelper(output, c.emoji, t)

            case "Cheerful":
                output := Cheerful("test")
                logHelper(output, c.emoji, t)

            case "Loved":
                output := Loved("test")
                logHelper(output, c.emoji, t)

            case "Stumped":
                output := Stumped("test")
                logHelper(output, c.emoji, t)
        }
    }
}

func logHelper(output string, emoji string, t *testing.T) int {
    if strings.Index(output, emoji) == -1 {
        t.Errorf("Expected " + emoji + " to exist")
    }
    return 0
}
