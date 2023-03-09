package entity

import (
	"reflect"
	"testing"
)

func Test_ParseOSArgs(t *testing.T) {
	type args struct {
		args []string
	}
	type want struct {
		value int
		error error
	}

	for _, testcase := range []struct {
		name string
		args args
		want *Params
	}{
		{
			name: "引数で渡された配列内の値が格納されたオブジェクトが返却されること",
			args: args{
				args: []string{"test", "-a=100", "-b=50"},
			},
			want: &Params{A: 100, B: 50},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			got := ParseOSArgs(testcase.args.args)

			if !reflect.DeepEqual(got, testcase.want) {
				t.Errorf("ParseOSArgs(): got %v, want %v", got, testcase.want)
			}
		})
	}
}

func Test_ParseJSArgs(t *testing.T) {
	type args struct {
		args []string
	}
	type want struct {
		value int
		error error
	}

	for _, testcase := range []struct {
		name string
		args args
		want *Params
	}{
		{
			name: "引数で渡された配列内の値が格納されたオブジェクトが返却されること",
			args: args{
				args: []string{"100", "50"},
			},
			want: &Params{A: 100, B: 50},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			got := ParseJSArgs(testcase.args.args)

			if !reflect.DeepEqual(got, testcase.want) {
				t.Errorf("ParseJSArgs(): got %v, want %v", got, testcase.want)
			}
		})
	}
}
