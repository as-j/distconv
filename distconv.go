// Package to convert distance between Miles and KM
package distconv

import "fmt"

type Mile float64
type KM float64

const (
	KMPerMile KM = 1.60934
)

func (m Mile) String() string { return fmt.Sprintf("%g mi", m) }
func (k KM) String() string   { return fmt.Sprintf("%g km", k) }
