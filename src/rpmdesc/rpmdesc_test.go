package local

import "testing"
import "strings"
import "github.com/PuerkitoBio/goquery"

func TestGetDescUrl(t *testing.T) {
	tt := []struct {
		name    string
		success bool
	}{
		{"librpm", true},
		{"libpython", true},
		{"libnss", true},
		{"vim", true},
		{"ssssssss", false},
	}
	parser := RpmFind{}
	for _, test := range tt {
		t.Run("GetDescUrl:"+test.name, func(t *testing.T) {
			searchURL := parser.GetSearchUrl(test.name, "", "")
			if (len(parser.GetDescUrl(searchURL)) > 0) != test.success {
				t.Error("Error")
			}
		})
	}
}

func TestGetLicense(t *testing.T) {
	tt := []struct {
		name    string
		license string
	}{
		{"librpm", "GPL"},
		{"vim", "Vim"},
		{"gimp", "GPLv2+"},
	}
	parser := RpmFind{}
	for _, test := range tt {
		t.Run("GetLicense:"+test.name, func(t *testing.T) {
			searchURL := parser.GetSearchUrl(test.name, "", "")
			doc, err := goquery.NewDocument(parser.GetPrefixURL() + parser.GetDescUrl(searchURL))
			if err != nil {
				t.Fatal("Can't obtain document")
			}
			license := strings.TrimRight(parser.GetLicenseFromRpm(doc), "\n")
			if strings.Compare(license, test.license) != 0 {
				t.Errorf("Invalid license %v != %v", license, test.license)
			}
		})
	}
}
