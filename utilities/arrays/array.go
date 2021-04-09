package arrays

import (
	"paintbot-client/models"
)

func Contains(ns []int, n int) bool {
	for i := range ns {
		if ns[i] == n {
			return true
		}
	}
	return false
}

func ContainsCoordinates(xs []models.Coordinates, x models.Coordinates) bool {
	for i := range xs {
		if xs[i].Equal(&x) {
			return true
		}
	}
	return false
}
