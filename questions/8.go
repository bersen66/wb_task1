package questions

import "fmt"

func SetBit(src *int64, bit int, value bool) error {
	if bit > 63 || bit < 0 {
		err := fmt.Errorf("Invalid bit idx must be [0, 64), got: %v", bit)
		return err
	}

	var mask int64 = 1 << bit
	bitVal := *src & mask

	if (bitVal != 0) && value == false {
		*src = *src & (^mask)
	} else if (bitVal == 0) && value == true {
		*src = *src | mask
	}

	return nil
}
