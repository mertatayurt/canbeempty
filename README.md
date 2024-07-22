# canbeempty

`canbeempty` is a Go package that provides a custom JSON marshaling function, `MarshalWithCanBeEmpty`, which considers the custom `can_be_empty` tag. Fields tagged with `can_be_empty` will be included in the JSON output even if they have zero values.

## Features

- Custom JSON marshaling with `can_be_empty` tag support.
- Simple and intuitive to use in any Go project.

## Installation

To install the package, run:

```sh
go get github.com/yourusername/canbeempty
```

## Usage

### Step 1: Import the package

Import the `canbeempty` package in your Go file:

```go
import "github.com/yourusername/canbeempty"
```

### Step 2: Define your struct with the `can_be_empty` tag

Define your struct and use the `can_be_empty` tag for fields that should be included in the JSON output even if they have zero values:

```go
type MyStruct struct {
    Name  string `json:"name,can_be_empty"`
    Age   int    `json:"age"`
    Score int    `json:"score"`
}
```

### Step 3: Use `MarshalWithCanBeEmpty` to marshal your struct

Use the `MarshalWithCanBeEmpty` function to marshal your struct to JSON:

```go
package main

import (
    "fmt"
    "log"

    "github.com/yourusername/canbeempty"
)

type MyStruct struct {
    Name  string `json:"name,can_be_empty"`
    Age   int    `json:"age"`
    Score int    `json:"score"`
}

func main() {
    myStruct := MyStruct{
        Name:  "",
        Age:   0,
        Score: 0,
    }

    jsonData, err := canbeempty.MarshalWithCanBeEmpty(myStruct)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(jsonData))
}
```

### Example Output

Given the struct:

```go
myStruct := MyStruct{
    Name:  "",
    Age:   0,
    Score: 0,
}
```

The output will be:

```json
{
    "name": "",
    "age": 0,
    "score": 0
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
