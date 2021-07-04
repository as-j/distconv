package distconv

// MtoKM converts Miles to KM
func MtoKM(m Mile) KM { return KM(m) * KMPerMile }

// KMtoMI convers KM to Miles
func KMtoM(k KM) Mile { return Mile(k/KMPerMile) }
