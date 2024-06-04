package main

import "fmt"

func poordesign(daysOff int, decision bool) (string, bool) {
    if daysOff > 10 {
        return "", false
    }
    if decision {
        return "hello world", decision
    } else {
        return "goodbye person", decision
    }

}

func main() {
    fmt.Println("call the poordesign function")
    fmt.Println(poordesign(100, true))
}
