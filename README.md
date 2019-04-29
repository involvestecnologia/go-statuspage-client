## go-statuspage-client

This is a client to easily communicate with the [Metronome API](https://github.com/involvestecnologia/statuspage)


### Current Features:
 - Create a Component
 - Find a Component by Name
 - Find Components by one or more labels
 - Create a Client

## Usage:

```golang
package main

import (
	"fmt"

	metro "github.com/involvestecnologia/go-statuspage-client"
)

func main() {
    m := metro.DefaultClient("https://metronome-api.com")

    prodComps, err := m.GetComponentsWithLabels("production")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    
    fmt.Printf("%+v", prodComps)
    
    myComp, err := m.FindComponent("myComponent")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Printf("Access %s in %s\n",myComp.Name,myComp.Address)
}
```