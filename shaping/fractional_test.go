// SPDX-License-Identifier: Unlicense OR BSD-3-Clause

package shaping

import (
	"testing"

	"github.com/go-text/typesetting/di"
	"github.com/go-text/typesetting/language"
	"golang.org/x/image/math/fixed"
)

// Both entry points must retain the caller's fractional ppem. An integer
// reference avoids encoding font-specific advances in this regression.
func TestFractionalPpem(t *testing.T) {
	for _, tc := range []struct {
		name  string
		shape func(*HarfbuzzShaper, Input) Output
	}{
		{"Shape", (*HarfbuzzShaper).Shape},
		{"ShapeNoExtents", (*HarfbuzzShaper).ShapeNoExtents},
	} {
		t.Run(tc.name, func(t *testing.T) {
			shaper := HarfbuzzShaper{}
			text := []rune("Guardar cambios")
			input := Input{
				Text: text, RunEnd: len(text), Face: benchEnFace,
				Direction: di.DirectionLTR, Script: language.Latin,
				Size: fixed.I(64),
			}
			reference := tc.shape(&shaper, input).Advance
			for _, size := range []fixed.Int26_6{9*64 + 40, 11*64 + 24, 12*64 + 16} {
				input.Size = size
				got := tc.shape(&shaper, input).Advance
				want := float64(reference) * float64(size) / float64(fixed.I(64))
				if deviation := float64(got)/want - 1; deviation < -0.005 || deviation > 0.005 {
					t.Errorf("size %v: advance %v differs from proportional advance %.4f by %+.3f%%", size, got, want/64, deviation*100)
				}
			}
		})
	}
}
