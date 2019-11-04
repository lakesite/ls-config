# config #

⚙ generic helper library for configurations ⚙

## usage ##

Getenv("NAME", "DEFAULT") returns the string matching env `NAME`, or `DEFAULT`
if this environment variable is not set.

```
...

import (
  "github.com/lakesite/ls-config/pkg/config"  
)
...

func main() {
  ...
  address := config.Getenv("APP_HOST", "127.0.0.1") + ":" + config.Getenv("APP_PORT", "7999")
  ...
}
```

## license ##

MIT
