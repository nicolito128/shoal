package shoal

import "fmt"

var (
	ErrReadOnlyCollection = fmt.Errorf("collection is read only")
	ErrIndexOutOfRaange   = fmt.Errorf("index for collection assign is out of range")
)
