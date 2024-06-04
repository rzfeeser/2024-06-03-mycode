/* Alta3 Research | RZFeeser
   Variables - Shadowing */


package main

import (
    "fmt"
)

func main() {
    var shadow string = "world"
    fmt.Println(shadow)     // world

    shadow = "56"

    {
        shadow := 55   // outer shadow is inaccessible from this point
        fmt.Println(shadow) // hello
    }

    fmt.Println(shadow)     // world
}

