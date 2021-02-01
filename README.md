# GoCollection

Simple key:value based utility data structure for golang!

## Quick Example

First install the package!

```sh
go get github.com/scientific-guy/gocollection
```

In your golang file!

```go
package main

import "fmt"
import collection "github.com/scientific-guy/gocollection"

func main() {
    col := collection.Collection()
    col.Set("foo", "bar")
    fmt.Println(col.Get("foo")) // Will return "bar"
}
```

It is basically a better version of `map` in golang! GoCollection uses `map[string]interface{}` as a map typing!

### Storing structs

So incase if you are not aware how to store structs using this, here is quick tuto

```go
package main

import "fmt"
import collection "github.com/scientific-guy/gocollection"

type SimpleStruct struct{
    field string
}

func main() {
    col := collection.Collection()
    col.Set("foo", SimpleStruct{ field: "string" })
    
    fmt.Println(col.Get("foo").(SimpleStruct).field) // You can use the basic type conversion of golang for this!
}
```

## Support

Any kind of doubts on this package, you can make an issue in the github repo or join our discord server and ask us doubts!

**Discord Server:** [https://discord.gg/FrduEZd](https://discord.gg/FrduEZd)<br/>
**GitHub Repo:** [https://github.com/Scientific-Guy/gocollection/](https://github.com/Scientific-Guy/gocollection)<br/>
**Docs:** [https://github.com/Scientific-Guy/gocollection/wiki/Go-Collection](https://github.com/Scientific-Guy/gocollection/wiki/Go-Collection)<br/>
