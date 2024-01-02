package soda

import (
	"fmt"
	"math"
)

type Size struct {
	Width, Height int
}

func (s Size) String() string {
	return fmt.Sprint(s.Width, "x", s.Height)
}

func (s Size) Area() int {
	return s.Width * s.Height
}

func (s Size) SplitVertical(topRatio float64) (Size, Size) {
	top, bottom := s.swapDimensions().SplitHorizotnal(topRatio)

	return top.swapDimensions(), bottom.swapDimensions()
}

func (s Size) SplitHorizotnal(leftRatio float64) (Size, Size) {
	width := float64(s.Width)
	left := width * leftRatio
	right := width - left

	return Size{
			Width:  int(math.Round(left)),
			Height: s.Height,
		}, Size{
			Width:  int(math.Round(right)),
			Height: s.Height,
		}
}

func (s Size) swapDimensions() Size {
	return Size{
		Width:  s.Height,
		Height: s.Width,
	}
}
