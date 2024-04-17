package font

import (
	"io/ioutil"
	"testing"

	"github.com/tdewolff/test"
)

func TestSFNTDejaVuSerifTTF(t *testing.T) {
	b, err := ioutil.ReadFile("resources/DejaVuSerif.ttf")
	test.Error(t, err)

	sfnt, err := ParseSFNT(b, 0)
	test.Error(t, err)

	test.T(t, sfnt.Head.UnitsPerEm, uint16(2048))
	test.T(t, sfnt.Hhea.Ascender, int16(1901))
	test.T(t, sfnt.Hhea.Descender, int16(-483))
	test.T(t, sfnt.OS2.SCapHeight, int16(1493)) // height of H glyph
	test.T(t, sfnt.Head.XMin, int16(-1576))
	test.T(t, sfnt.Head.YMin, int16(-710))
	test.T(t, sfnt.Head.XMax, int16(4312))
	test.T(t, sfnt.Head.YMax, int16(2272))

	id := sfnt.GlyphIndex(' ')
	contour, err := sfnt.Glyf.Contour(id)
	test.Error(t, err)
	test.T(t, contour.GlyphID, id)
	test.T(t, len(contour.XCoordinates), 0)
}

func TestSFNTWrite(t *testing.T) {
	b, err := ioutil.ReadFile("resources/DejaVuSerif.ttf")
	test.Error(t, err)

	sfnt, err := ParseSFNT(b, 0)
	test.Error(t, err)

	b2 := sfnt.Write()
	sfnt2, err := ParseSFNT(b2, 0)
	test.Error(t, err)

	test.T(t, sfnt2.GlyphIndex('A'), sfnt.GlyphIndex('A'))
	test.T(t, sfnt2.GlyphIndex('B'), sfnt.GlyphIndex('B'))
	test.T(t, sfnt2.GlyphIndex('C'), sfnt.GlyphIndex('C'))

	//ioutil.WriteFile("out.otf", subset, 0644)
}

func TestSFNTSubset(t *testing.T) {
	b, err := ioutil.ReadFile("resources/DejaVuSerif.ttf")
	test.Error(t, err)

	sfnt, err := ParseSFNT(b, 0)
	test.Error(t, err)

	sfntSubset, err := sfnt.Subset([]uint16{0, 3, 6, 36, 37, 38, 55, 131}, SubsetOptions{Tables: KeepAllTables}) // .notdef, space, #, A, B, C, T, Á
	test.Error(t, err)

	test.T(t, sfntSubset.NumGlyphs(), uint16(9)) // Á is a composite glyph containing two simple glyphs: 36 and 3452

	test.T(t, sfntSubset.GlyphIndex('A'), uint16(3))
	test.T(t, sfntSubset.GlyphIndex('B'), uint16(4))
	test.T(t, sfntSubset.GlyphIndex('C'), uint16(5))

	//ioutil.WriteFile("out.otf", subset, 0644)
}
