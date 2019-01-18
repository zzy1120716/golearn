package unitconv

// length
func MToFt(m Meter) Foot  { return Foot(m * 3.28) }
func FtToM(ft Foot) Meter { return Meter(ft / 3.28) }

// weight
func KgToLb(kg Kilogram) Pound { return Pound(kg * 2.2) }
func LbToKg(lb Pound) Kilogram { return Kilogram(lb / 2.2) }
