#Import alies

In the file **main.go** you will see that mysql driver import is thrown away.
This package only is needed for sql package to run with a query language driver.
When evertyhing is imported the package `go-sql-driver/mysql` will be setup implicitly by `database/sql`

```go
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)
```