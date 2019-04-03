package opencc

import "testing"

func TestConvert(t *testing.T) {
	cases := []struct {
		text, format, expected string
	}{
		{"開放中文轉換", "t2s", "开放中文转换"},
		{"开放中文转换", "s2t", "開放中文轉換"},
	}

	converter := Converter{}
	defer converter.Close()

	for _, c := range cases {
		if got, err := converter.Convert(c.text, c.format); err != nil {
			t.Error(err)
		} else if got != c.expected {
			t.Errorf("Convert(%q) == %q, want %q", c.text, got, c.expected)
		}
	}
}
