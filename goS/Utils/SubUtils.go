package Utils

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

func Sub(a, b int) int {
	fmt.Println("Utils moduls=> Subutils.go", a, b)
	log.Info().Msg("============ Sub ==========")
	return a - b
}
