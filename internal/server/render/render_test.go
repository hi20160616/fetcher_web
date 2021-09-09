package render

import (
	"testing"
)

func TestMarkdown(t *testing.T) {
	tcs := []struct {
		in  string
		out string
	}{
		{
			in:  "# test title h1",
			out: "<h1>test title h1</h1>",
		},
		{
			in:  "**临近32周年纪念，当年学生领袖之一的周锋锁这几天忙得不可开交，连续多日在网上举办马拉松式的纪念活动。**",
			out: "<p><strong>临近32周年纪念，当年学生领袖之一的周锋锁这几天忙得不可开交，连续多日在网上举办马拉松式的纪念活动。</strong></p>"},
	}
	for _, tc := range tcs {
		got, err := markdown(tc.in)
		if err != nil {
			t.Error(err)
		}
		if tc.out != got {
			t.Errorf("\nwant: %s\n got: %s", tc.out, got)
		}
	}
}
