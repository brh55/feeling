# Feeling
Feeling is a simple GoLang package to express your feelings through text emoticons

Get:

`go get gopkg.in/brh55/feeling.v1`

Import:

`import "gopkg.in/brh55/feeling.v1"`

## Example
````
package main

import (
    "fmt"
    "gopkg.in/brh55/feeling.v1"
)

func main() {
    fmt.Println(feeling.Scared("I give up if this won't work!"))
    fmt.Println(feeling.Uncertain("Balls deep, here we go..."))
    fmt.Println(feeling.Cheerful("Sing me a song if this works!"))
}
````

![output](https://cloud.githubusercontent.com/assets/6020066/12633977/cf95c134-c543-11e5-8ff4-39b5455f4033.png)
