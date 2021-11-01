package service

import "github.com/diegoclair/branas-clean-code-arch/domain/entity"

func FreightCalculator(item entity.Item) (freight float64) {
	freight = 1000 * item.GetVolume() * (item.GetDensity() / 100)
	if freight < 10 {
		return 10
	}
	return freight
}
