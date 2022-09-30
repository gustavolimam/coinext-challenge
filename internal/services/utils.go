package services

import "github.com/gustavolimam/coinext-challenge/internal/model"

func sumItemsPoints(item model.Inventory) int {
	sum := item.Water * 4
	sum += item.Food * 3
	sum += item.Drug * 2
	sum += item.Ammo * 1

	return sum
}
