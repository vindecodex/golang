```go
import . "fmt"

Println("Hello")

import _ "fmt" // We are just using the init function of that package not the functions inside it

import f "fmt"

f.Println("Hello")
```
