package err

import (
    "fmt"
    "time"
)

type WorseError struct {
    When time.Time
    Wat string
}

func (e *WorseError) Error() string{
    return fmt.Sprintf("Worse error: '%s' @%v", e.Wat, e.When)
}
