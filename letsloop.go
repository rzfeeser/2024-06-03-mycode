package main

import "fmt"

type Studentrecord struct {
    name string
    id int
}

func main() {

    // loop with a slice
    movies := []string{"Aliens", "Bob, What about", "Cars"}

    for position, value := range movies {
        fmt.Println("The movie", value, "is", position, "in the slice")
    }

    // range will yield just one result if you only capture one (the position)
    for position := range movies {
        fmt.Println("The movie position is", position)
    }

    // the value is stored in the second returned result, so if you don't want position, throw it out with _
    for _, value := range movies {
        fmt.Println("The movie was", value)
    }

    // it doesn't matter what is in the slice, we can still loop across it
    roster := []Studentrecord{Studentrecord{"larry", 24601}, Studentrecord{"alice", 10110}}

    for position, value := range roster {
        fmt.Println("The record", value, "is in position", position)
    }

    // suppose we don't care about the position but care about elements within the individual records
    for _, value := range roster {
        fmt.Println("The student", value.name, "has the identification", value.id, "assigned to them")
    }

    // we can also loop across a map
    piratetreasure := map[string]string{
        "dabloons": "under the palm tree shaped like a heart",
        "swords": "cave behind waterfall",
        "gems": "under the big X",
    }
    // maps do NOT have position! They have keys and values
          // keys          // values
    for left_of_colon, right_of_colon := range piratetreasure {
        fmt.Println("The treasure", left_of_colon, "is found at", right_of_colon)
    }

}
