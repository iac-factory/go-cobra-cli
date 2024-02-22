package color

import (
	"testing"
)

func Test(t *testing.T) {
	ci = false

	t.Run("Colors", func(t *testing.T) {
		New().White("White").Black("Black").Cyan("Cyan").Red("Red").Yellow("Yellow").Green("Green").Gray("Gray").Purple("Purple").Blue("Blue").Default("Default").Write()
	})

	t.Run("CI", func(t *testing.T) {
		// override for uncolored output
		ci = true

		t.Run("Color", func(t *testing.T) {
			New().White("White").Black("Black").Cyan("Cyan").Red("Red").Yellow("Yellow").Green("Green").Gray("Gray").Purple("Purple").Blue("Blue").Default("Default").Write()
		})

		ci = false
	})

	t.Run("Bold", func(t *testing.T) {
		New().Bold("Bold").Write()

		t.Run("Red", func(t *testing.T) {
			New().Bold(New().Red("Bold & Red").String()).Write()
		})

		t.Run("Green", func(t *testing.T) {
			New().Bold(New().Green("Bold & Green").String()).Write()
		})

		t.Run("Blue", func(t *testing.T) {
			New().Bold(New().Blue("Bold & Blue").String()).Write()
		})
	})

	t.Run("Italic", func(t *testing.T) {
		New().Italic("Italic").Write()

		t.Run("Red", func(t *testing.T) {
			New().Italic(New().Red("Italic & Red").String()).Write()
		})

		t.Run("Green", func(t *testing.T) {
			New().Italic(New().Green("Italic & Green").String()).Write()
		})

		t.Run("Blue", func(t *testing.T) {
			New().Italic(New().Blue("Italic & Blue").String()).Write()
		})
	})

	t.Run("Special", func(t *testing.T) {
		New().Bold(New().Italic("Bold & Italic").String()).Write()

		t.Run("Red", func(t *testing.T) {
			New().Bold(New().Italic(New().Red("Bold & Italic & Red").String()).String()).Write()
		})

		t.Run("Green", func(t *testing.T) {
			New().Bold(New().Italic(New().Green("Bold & Italic & Green").String()).String()).Write()
		})

		t.Run("Blue", func(t *testing.T) {
			New().Bold(New().Italic(New().Blue("Bold & Italic & Blue").String()).String()).Write()
		})
	})
}
