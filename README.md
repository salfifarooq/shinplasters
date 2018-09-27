# shinplasters
Currency code to subunit mapping.

#### Installation

`go get github.com/salfifarooq/shinplasters`



### Usage

#### GetAll()

```go

import sp "github.com/salfifarooq/shinplasters"

subunits , err := sp.GetAll()
if err != nil {
      //....
}

```

#### GetSubunit

```go

import sp "github.com/salfifarooq/shinplasters"

value , err := sp.GetSubunit("USD")
if err != nil {
      //....
}
  
// returns 100 

```
#### License

MIT
