package iphash

import (
	"net/url"
	"reflect"
	"testing"
)

func TestIPHash(t *testing.T) {
	tests := []struct {
		urls  []*url.URL
		ips   []*url.URL
		iserr bool
		want  []*url.URL
	}{
		{
			urls: []*url.URL{
				{Host: "192.168.1.10"},
				{Host: "192.168.1.11"},
				{Host: "192.168.1.12"},
				{Host: "192.168.1.13"},
			},
			ips: []*url.URL{
				{Host: "192.168.1.10"},
				{Host: "192.168.1.10"},
				{Host: "192.168.1.44"},
				{Host: "192.168.1.44"},
			},
			iserr: false,
			want: []*url.URL{
				{Host: "192.168.1.10"},
				{Host: "192.168.1.10"},
				{Host: "192.168.1.11"},
				{Host: "192.168.1.11"},
			},
		},
	}

	for _, test := range tests {
		iphash, err := New(test.urls)

		if got, want := !(err == nil), test.iserr; got != want {
			t.Errorf("IPHash err is wrong. want: %v, but got: %v", want, got)
		}

		gots := make([]*url.URL, 0, len(test.want))
		for _, ip := range test.ips {
			gots = append(gots, iphash.Next(ip))
		}

		if got, want := gots, test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("IPHash is wrong. want: %v, but got: %v", want, got)
		}
	}
}
