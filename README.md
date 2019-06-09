# go-smart-error
Golang smart error implementation

# Features

- Saving the position where the error was initiated
- Saving the error trace
- Short/Full error string
- The same interface with standard error 

# Using

```go
import se "github.com/panter-dsd/go-smart-error"

func foo() error {
  err := bar()
  return se.Wrapf(err, "foo %s", "some")
}
...........
err := foo()
if err != nil {
    fmt.Println(se.FullError(err))
}
```

# Config 
- se.IsFullErrorInsteadError (false by default) - return FullError when Error called
- se.IsShortError (false by default) - print only last error message
