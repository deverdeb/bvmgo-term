package ansi

import (
	"reflect"
	"testing"
)

func TestTextSplitToSize(t *testing.T) {
	type args struct {
		str      string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty text",
			args: args{
				str:      "",
				maxWidth: 10,
			},
			want: []string{""},
		},
		{
			name: "Text without ANSI sequence",
			args: args{
				str:      "toto tata titi",
				maxWidth: 10,
			},
			want: []string{"toto tata", "titi"},
		},
		{
			name: "Text without ANSI sequence",
			args: args{
				str:      StyleBold() + "hello world" + StyleReset(),
				maxWidth: 10,
			},
			want: []string{StyleBold() + "hello", "world" + StyleReset()},
		},
		{
			name: "Text with end line",
			args: args{
				str:      "hello\nworld everybody",
				maxWidth: 10,
			},
			want: []string{"hello", "world", "everybody"},
		},
		{
			name: "Text with too long word",
			args: args{
				str:      "thisWordIsTooLong",
				maxWidth: 10,
			},
			want: []string{"thisWordIs", "TooLong"},
		},
		{
			name: "Text with multiple space",
			args: args{
				str:      "  sp  ace ",
				maxWidth: 10,
			},
			want: []string{"  sp  ace "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TextSplitToSize(tt.args.str, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextSplitToSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
