# libtime

## Usage
```go
import (
    "github.com/Delisa-sama/libtime"
)

type Shop struct {
    Open  libtime.Time
    Close libtime.Time
}

func main() {
    var s Shop
    data := []byte(`{"open":"09:00-03:00","close":"09:00-03:00"}`)
    if err := json.Unmarshal(data, &s); err != nil {
    	// handle unmarshaling error
    }
    now := time.Now()
    if s.Open.Before(now) && s.Close.After(now) {
    	// shop is open now
    }
}
```