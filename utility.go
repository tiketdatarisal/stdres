package stdres

import (
	"math"
	"time"
)

func normalizeParam(pageNumber, pageSize, total int64) (normPageNumber, normPageSize, normPageTotal, normTotal int64) {
	normTotal = total
	if normTotal < minTotal {
		normTotal = minTotal
	}

	normPageSize = pageSize
	if normPageSize < minPageSize {
		normPageSize = minPageSize
	}

	normPageTotal = int64(math.Ceil(float64(normTotal) / float64(normPageSize)))

	normPageNumber = pageNumber
	if normPageNumber < minPageNumber {
		normPageNumber = minPageNumber
	}

	if normPageNumber > normPageTotal {
		normPageNumber = normPageTotal
	}

	return
}

func getCurrentServerTime(serverTime ...int64) int64 {
	if len(serverTime) > 0 {
		return serverTime[0]
	}

	return time.Now().UTC().UnixNano() / oneMillion
}
