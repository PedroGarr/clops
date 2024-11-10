package utils

import (
	"strconv"
	"time"
)

func ConvertEpochToHuman(convertEpoch string) string {
	epoch, err := strconv.ParseInt(convertEpoch, 10, 64)
	// Error treatment
	if err != nil {
		return " Erro in epoch conversion"
	}

	// Temps is just time in french, to not overlap with the package time
	temps := time.Unix(epoch, 0)
	return temps.Format("01-Jan-2006 15:04:05")
}
