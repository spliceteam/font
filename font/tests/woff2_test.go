package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/tdewolff/canvas/font"
	"github.com/tdewolff/test"
)

func TestWOFF2ValidationDecoder(t *testing.T) {
	filenames := []string{
		"validation-checksum-001",
		"validation-checksum-002",
		"validation-loca-format-001",
		"validation-loca-format-002",
	}
	for i := 1; i < 156; i++ {
		filenames = append(filenames, fmt.Sprintf("validation-off-%03d", i))
	}
	for _, filename := range filenames {
		t.Run(filename, func(t *testing.T) {
			b, err := ioutil.ReadFile("decoder/" + filename + ".woff2")
			test.Error(t, err)
			_, err = font.ParseWOFF2(b)
			test.Error(t, err)
		})
	}
}

func TestWOFF2ValidationDecoderRoundtrip(t *testing.T) {
	filenames := []string{
		"roundtrip-collection-dsig-001",
		"roundtrip-collection-order-001",
		"roundtrip-hmtx-lsb-001",
		"roundtrip-offset-tables-001",
	}
	for _, filename := range filenames {
		t.Run(filename, func(t *testing.T) {
			a, err := ioutil.ReadFile("decoder/" + filename + ".ttf")
			test.Error(t, err)
			b, err := ioutil.ReadFile("decoder/" + filename + ".woff2")
			test.Error(t, err)
			b, err = font.ParseWOFF2(b)
			test.Error(t, err)
			if !bytes.Equal(a, b) {
				test.Fail(t, "decoded WOFF2 unequal to TTF")
			}
		})
	}
}

func TestWOFF2ValidationFormat(t *testing.T) {
	var tts = []struct {
		filename string
		err      string
	}{
		{"valid-001", ""},
		{"valid-002", ""},
		{"valid-003", ""},
		{"valid-004", ""},
		{"valid-005", ""},
		{"valid-006", ""},
		{"valid-007", ""},
		{"valid-008", ""},
		{"header-signature-001", "err"},
		{"header-flavor-001", "err"},
		{"header-flavor-002", "err"},
		{"header-length-001", "err"},
		{"header-length-002", "err"},
		{"header-numTables-001", "err"},
		{"header-reserved-001", "err"},
		{"blocks-extraneous-data-001", "err"},
		{"blocks-extraneous-data-002", "err"},
		{"blocks-extraneous-data-003", "err"},
		{"blocks-extraneous-data-004", "err"},
		{"blocks-extraneous-data-005", "err"},
		{"blocks-extraneous-data-006", "err"},
		{"blocks-extraneous-data-007", "err"},
		{"blocks-metadata-absent-002", "err"},
		{"blocks-metadata-padding-001", "err"},
		{"blocks-metadata-padding-002", "err"},
		{"blocks-metadata-padding-003", "err"},
		{"blocks-metadata-padding-004", "err"},
		{"blocks-ordering-003", "err"},
		{"blocks-ordering-004", "err"},
		{"blocks-private-001", "err"},
		{"blocks-private-002", "err"},
		{"directory-table-order-001", ""},
		{"directory-table-order-002", "loca: must come after glyf table"},
		{"tabledata-extraneous-data-001", "err"},
		{"tabledata-brotli-001", "brotli: corrupted input"},
		{"tabledata-decompressed-length-001", "err"},
		{"tabledata-decompressed-length-002", "err"},
		{"tabledata-decompressed-length-003", "err"},
		{"tabledata-decompressed-length-004", "err"},
		{"tabledata-transform-length-001", "loca: transformLength must be zero"},
		{"tabledata-transform-length-002", "err"},
		{"tabledata-loca-size-001", ""},
		{"tabledata-loca-size-002", ""},
		{"tabledata-loca-size-003", ""},
		{"tabledata-hmtx-transform-001", ""},
		{"tabledata-hmtx-transform-002", ""},
		{"tabledata-hmtx-transform-003", ""},
		{"tabledata-transform-glyf-loca-001", "glyf and loca tables must be both present and either be both transformed or untransformed"},
		{"tabledata-transform-glyf-loca-002", "glyf and loca tables must be both present and either be both transformed or untransformed"},
		{"tabledata-glyf-composite-bbox-001", ""},
		//{"metadata-padding-001", "err"},
		//{"metadata-compression-001", "err"},
		//{"metadata-compression-002", "err"},
		//{"metadata-metaOrigLength-001", "err"},
		//{"metadata-metaOrigLength-002", "err"},
		//{"metadata-well-formed-001", "err"},
		//{"metadata-well-formed-002", "err"},
		//{"metadata-well-formed-003", "err"},
		//{"metadata-well-formed-004", "err"},
		//{"metadata-well-formed-005", "err"},
		//{"metadata-well-formed-006", "err"},
		//{"metadata-well-formed-007", "err"},
		//{"metadata-encoding-001", ""},
		//{"metadata-encoding-002", "err"},
		//{"metadata-encoding-003", "err"},
		//{"metadata-encoding-004", ""},
		//{"metadata-encoding-005", ""},
		//{"metadata-encoding-006", "err"},
		//{"metadata-schema-metadata-001", ""},
		//{"metadata-schema-metadata-002", "err"},
		//{"metadata-schema-metadata-003", "err"},
		//{"metadata-schema-metadata-004", "err"},
		//{"metadata-schema-metadata-005", "err"},
		//{"metadata-schema-metadata-006", "err"},
		//{"metadata-schema-uniqueid-001", ""},
		//{"metadata-schema-uniqueid-002", ""},
		//{"metadata-schema-uniqueid-003", "err"},
		//{"metadata-schema-uniqueid-004", "err"},
		//{"metadata-schema-uniqueid-005", "err"},
		//{"metadata-schema-uniqueid-006", "err"},
		//{"metadata-schema-uniqueid-007", "err"},
		//{"metadata-schema-vendor-001", ""},
		//{"metadata-schema-vendor-002", ""},
		//{"metadata-schema-vendor-003", ""},
		//{"metadata-schema-vendor-004", "err"},
		//{"metadata-schema-vendor-005", "err"},
		//{"metadata-schema-vendor-006", ""},
		//{"metadata-schema-vendor-007", ""},
		//{"metadata-schema-vendor-008", "err"},
		//{"metadata-schema-vendor-009", ""},
		//{"metadata-schema-vendor-010", "err"},
		//{"metadata-schema-vendor-011", "err"},
		//{"metadata-schema-vendor-012", "err"},
		//{"metadata-schema-credits-001", ""},
		//{"metadata-schema-credits-002", ""},
		//{"metadata-schema-credits-003", "err"},
		//{"metadata-schema-credits-004", "err"},
		//{"metadata-schema-credits-005", "err"},
		//{"metadata-schema-credits-006", "err"},
		//{"metadata-schema-credits-007", "err"},
		//{"metadata-schema-credit-001", ""},
		//{"metadata-schema-credit-002", ""},
		//{"metadata-schema-credit-003", ""},
		//{"metadata-schema-credit-004", "err"},
		//{"metadata-schema-credit-005", ""},
		//{"metadata-schema-credit-006", ""},
		//{"metadata-schema-credit-007", "err"},
		//{"metadata-schema-credit-008", ""},
		//{"metadata-schema-credit-009", "err"},
		//{"metadata-schema-credit-010", "err"},
		//{"metadata-schema-credit-011", "err"},
		//{"metadata-schema-description-001", ""},
		//{"metadata-schema-description-002", ""},
		//{"metadata-schema-description-003", ""},
		//{"metadata-schema-description-004", ""},
		//{"metadata-schema-description-005", ""},
		//{"metadata-schema-description-006", ""},
		//{"metadata-schema-description-007", ""},
		//{"metadata-schema-description-008", "err"},
		//{"metadata-schema-description-009", "err"},
		//{"metadata-schema-description-010", "err"},
		//{"metadata-schema-description-011", "err"},
		//{"metadata-schema-description-012", "err"},
		//{"metadata-schema-description-013", ""},
		//{"metadata-schema-description-014", ""},
		//{"metadata-schema-description-015", "err"},
		//{"metadata-schema-description-016", ""},
		//{"metadata-schema-description-017", "err"},
		//{"metadata-schema-description-018", "err"},
		//{"metadata-schema-description-019", ""},
		//{"metadata-schema-description-020", ""},
		//{"metadata-schema-description-021", ""},
		//{"metadata-schema-description-022", ""},
		//{"metadata-schema-description-023", ""},
		//{"metadata-schema-description-024", "err"},
		//{"metadata-schema-description-025", ""},
		//{"metadata-schema-description-026", ""},
		//{"metadata-schema-description-027", ""},
		//{"metadata-schema-description-028", ""},
		//{"metadata-schema-description-029", ""},
		//{"metadata-schema-description-030", ""},
		//{"metadata-schema-description-031", "err"},
		//{"metadata-schema-description-032", ""},
		//{"metadata-schema-license-001", ""},
		//{"metadata-schema-license-002", ""},
		//{"metadata-schema-license-003", ""},
		//{"metadata-schema-license-004", ""},
		//{"metadata-schema-license-005", ""},
		//{"metadata-schema-license-006", ""},
		//{"metadata-schema-license-007", ""},
		//{"metadata-schema-license-008", ""},
		//{"metadata-schema-license-009", "err"},
		//{"metadata-schema-license-010", ""},
		//{"metadata-schema-license-011", "err"},
		//{"metadata-schema-license-012", "err"},
		//{"metadata-schema-license-013", "err"},
		//{"metadata-schema-license-014", ""},
		//{"metadata-schema-license-015", ""},
		//{"metadata-schema-license-016", "err"},
		//{"metadata-schema-license-017", ""},
		//{"metadata-schema-license-018", "err"},
		//{"metadata-schema-license-019", "err"},
		//{"metadata-schema-license-020", ""},
		//{"metadata-schema-license-021", ""},
		//{"metadata-schema-license-022", ""},
		//{"metadata-schema-license-023", ""},
		//{"metadata-schema-license-024", ""},
		//{"metadata-schema-license-025", "err"},
		//{"metadata-schema-license-026", ""},
		//{"metadata-schema-license-027", ""},
		//{"metadata-schema-license-028", ""},
		//{"metadata-schema-license-029", ""},
		//{"metadata-schema-license-030", ""},
		//{"metadata-schema-license-031", ""},
		//{"metadata-schema-license-032", "err"},
		//{"metadata-schema-license-033", ""},
		//{"metadata-schema-copyright-001", ""},
		//{"metadata-schema-copyright-002", ""},
		//{"metadata-schema-copyright-003", ""},
		//{"metadata-schema-copyright-004", ""},
		//{"metadata-schema-copyright-005", ""},
		//{"metadata-schema-copyright-006", "err"},
		//{"metadata-schema-copyright-007", "err"},
		//{"metadata-schema-copyright-008", "err"},
		//{"metadata-schema-copyright-009", "err"},
		//{"metadata-schema-copyright-010", "err"},
		//{"metadata-schema-copyright-011", ""},
		//{"metadata-schema-copyright-012", ""},
		//{"metadata-schema-copyright-013", "err"},
		//{"metadata-schema-copyright-014", ""},
		//{"metadata-schema-copyright-015", "err"},
		//{"metadata-schema-copyright-016", "err"},
		//{"metadata-schema-copyright-017", ""},
		//{"metadata-schema-copyright-018", ""},
		//{"metadata-schema-copyright-019", ""},
		//{"metadata-schema-copyright-020", ""},
		//{"metadata-schema-copyright-021", ""},
		//{"metadata-schema-copyright-022", "err"},
		//{"metadata-schema-copyright-023", ""},
		//{"metadata-schema-copyright-024", ""},
		//{"metadata-schema-copyright-025", ""},
		//{"metadata-schema-copyright-026", ""},
		//{"metadata-schema-copyright-027", ""},
		//{"metadata-schema-copyright-028", ""},
		//{"metadata-schema-copyright-029", "err"},
		//{"metadata-schema-copyright-030", ""},
		//{"metadata-schema-trademark-001", ""},
		//{"metadata-schema-trademark-002", ""},
		//{"metadata-schema-trademark-003", ""},
		//{"metadata-schema-trademark-004", ""},
		//{"metadata-schema-trademark-005", ""},
		//{"metadata-schema-trademark-006", "err"},
		//{"metadata-schema-trademark-007", "err"},
		//{"metadata-schema-trademark-008", "err"},
		//{"metadata-schema-trademark-009", "err"},
		//{"metadata-schema-trademark-010", "err"},
		//{"metadata-schema-trademark-011", ""},
		//{"metadata-schema-trademark-012", ""},
		//{"metadata-schema-trademark-013", "err"},
		//{"metadata-schema-trademark-014", ""},
		//{"metadata-schema-trademark-015", "err"},
		//{"metadata-schema-trademark-016", "err"},
		//{"metadata-schema-trademark-017", ""},
		//{"metadata-schema-trademark-018", ""},
		//{"metadata-schema-trademark-019", ""},
		//{"metadata-schema-trademark-020", ""},
		//{"metadata-schema-trademark-021", ""},
		//{"metadata-schema-trademark-022", "err"},
		//{"metadata-schema-trademark-023", ""},
		//{"metadata-schema-trademark-024", ""},
		//{"metadata-schema-trademark-025", ""},
		//{"metadata-schema-trademark-026", ""},
		//{"metadata-schema-trademark-027", ""},
		//{"metadata-schema-trademark-028", ""},
		//{"metadata-schema-trademark-029", "err"},
		//{"metadata-schema-trademark-030", ""},
		//{"metadata-schema-licensee-001", ""},
		//{"metadata-schema-licensee-002", "err"},
		//{"metadata-schema-licensee-003", "err"},
		//{"metadata-schema-licensee-004", ""},
		//{"metadata-schema-licensee-005", ""},
		//{"metadata-schema-licensee-006", "err"},
		//{"metadata-schema-licensee-007", ""},
		//{"metadata-schema-licensee-008", "err"},
		//{"metadata-schema-licensee-009", "err"},
		//{"metadata-schema-licensee-010", "err"},
		//{"metadata-schema-extension-001", ""},
		//{"metadata-schema-extension-002", ""},
		//{"metadata-schema-extension-003", ""},
		//{"metadata-schema-extension-004", ""},
		//{"metadata-schema-extension-005", ""},
		//{"metadata-schema-extension-006", ""},
		//{"metadata-schema-extension-007", ""},
		//{"metadata-schema-extension-008", "err"},
		//{"metadata-schema-extension-009", "err"},
		//{"metadata-schema-extension-010", "err"},
		//{"metadata-schema-extension-011", "err"},
		//{"metadata-schema-extension-012", ""},
		//{"metadata-schema-extension-013", ""},
		//{"metadata-schema-extension-014", ""},
		//{"metadata-schema-extension-015", ""},
		//{"metadata-schema-extension-016", ""},
		//{"metadata-schema-extension-017", "err"},
		//{"metadata-schema-extension-018", ""},
		//{"metadata-schema-extension-019", "err"},
		//{"metadata-schema-extension-020", "err"},
		//{"metadata-schema-extension-021", ""},
		//{"metadata-schema-extension-022", ""},
		//{"metadata-schema-extension-023", ""},
		//{"metadata-schema-extension-024", ""},
		//{"metadata-schema-extension-025", ""},
		//{"metadata-schema-extension-026", ""},
		//{"metadata-schema-extension-027", ""},
		//{"metadata-schema-extension-028", "err"},
		//{"metadata-schema-extension-029", "err"},
		//{"metadata-schema-extension-030", "err"},
		//{"metadata-schema-extension-031", "err"},
		//{"metadata-schema-extension-032", "err"},
		//{"metadata-schema-extension-033", ""},
		//{"metadata-schema-extension-034", ""},
		//{"metadata-schema-extension-035", ""},
		//{"metadata-schema-extension-036", ""},
		//{"metadata-schema-extension-037", ""},
		//{"metadata-schema-extension-038", "err"},
		//{"metadata-schema-extension-039", ""},
		//{"metadata-schema-extension-040", "err"},
		//{"metadata-schema-extension-041", "err"},
		//{"metadata-schema-extension-042", ""},
		//{"metadata-schema-extension-043", ""},
		//{"metadata-schema-extension-044", ""},
		//{"metadata-schema-extension-045", ""},
		//{"metadata-schema-extension-046", ""},
		//{"metadata-schema-extension-047", "err"},
		//{"metadata-schema-extension-048", ""},
		//{"metadata-schema-extension-049", "err"},
		//{"metadata-schema-extension-050", "err"},
	}
	for _, tt := range tts {
		t.Run(tt.filename, func(t *testing.T) {
			b, err := ioutil.ReadFile("format/" + tt.filename + ".woff2")
			test.Error(t, err)
			_, err = font.ParseWOFF2(b)
			if tt.err == "" {
				test.Error(t, err)
			} else if err == nil {
				test.Fail(t, "must give error")
			} else {
				test.T(t, err.Error(), tt.err)
			}
		})
	}
}
