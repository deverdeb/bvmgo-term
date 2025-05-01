package ansi

import "testing"

func TestStringLen(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty string",
			args: args{str: ""},
			want: 0,
		}, {
			name: "String with numbers and letters",
			args: args{str: "Hello world 145"},
			want: 15,
		}, {
			name: "String with special characters",
			args: args{str: "Hello\nworld\t&é\"'(-è_çà$^*ù"},
			want: 26,
		}, {
			name: "String with ANSI sequences",
			args: args{str: StyleBold() + "Hello " + StyleUnderline() + "world" + StyleReset()},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringLen(tt.args.str); got != tt.want {
				t.Errorf("StringLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringRemoveAnsi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty string",
			args: args{str: ""},
			want: "",
		}, {
			name: "String without ANSI sequences",
			args: args{str: "Hello world 145"},
			want: "Hello world 145",
		}, {
			name: "String with special characters",
			args: args{str: "Hello\nworld\t&é\"'(-è_çà$^*ù"},
			want: "Hello\nworld\t&é\"'(-è_çà$^*ù",
		}, {
			name: "String with ANSI escape sequences",
			args: args{str: StyleBold() + "Hello " + StyleUnderline() + "world" + StyleReset()},
			want: "Hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringRemoveAnsi(tt.args.str); got != tt.want {
				t.Errorf("StringRemoveAnsi() = %v, want %v", got, tt.want)
			}
		})
	}
}
