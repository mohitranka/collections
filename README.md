# collections

A collections of various datastructures in go

## Usage

```lang=go

package main

import ( 
    "fmt"
    "github.com/mohitranka/collections/skiplist"
)

func main() {
    sl := skiplist.NewSkipList(3, 0.5)
    sl.Insert(10)
    sl.Insert(20)
    sl.Insert(5)

    sl.Search(5)    // returns true
    sl.Search(17)   // returns false

    sl.Delete(20)   // returns true
    sl.Search(20)   // returns false

    fmt.Print(sl.Display()) // Print the string representation on stdout
}
```

## Authors

* Mohit Ranka

## License

This project is licensed under the MIT License - See [LICENSE](LICENSE) for details.

