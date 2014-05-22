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
    return fmt.Sprintf("Whale error: '%s' @ %v", e.Wat, e.When)
}

func failWhale() error {
    return &WorseError{ time.Now(), "bad error was bad" }
}
