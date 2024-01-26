Like in real life, a map is a way of containing data that defines something. Think of the map of New York City. It is a piece of paper that includes data about the city's layout– depending on the map type.
Maps are typically handy when you combine different data types because a slice will only permit one data type per-use.
We'll briefly cover Maps so the reader can be aware of them and see when they might be useful.

## Syntax
A map uses the 'map' keyword and [type] followed again by 'type.'

```
package main

func main() {
  nyc := map[string]int{
      "population": 8000000,
      "size": 302,
  }
  htx := map[string]int{
  	"population": 2000000,
    "size": 665,
  }
}
```

It says a map made of key-pair values of type string and int.
You can have any combination of types with maps, which is when you should consider using them over a slice.

## Use case
Although the data is small above it is enough to let us access quick facts about each city with a fast lookup, with syntax similar to a slice.

```
package main

import "fmt"
func main() {
  nyc := map[string]int{
      "population": 8000000,
      "size": 302,
  }
  htx := map[string]int{
  	"population": 2000000,
    "size": 665,
  }
  fmt.Println(htx["population"], nyc["population"])
}
```

The curious reader should take a look at Go Maps in Action by the GO dev team to get a better sense of when to use 'maps'.

## Conclusion
It is enough to know a map and how it generally works. Now, let's move on to part II of automating the boring stuff with go, which contains much more fun practical examples.

