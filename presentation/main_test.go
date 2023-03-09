package main

import (
	"context"
	"go-onion-sample/usecase"
	"reflect"
	"testing"
)

func Test_di(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
	}
	type want struct {
		value usecase.ICalcService
		error error
	}

	for _, testcase := range []struct {
		name string
		args args
		want want
	}{
		{
			name: "ICalcService が実装されたインスタンスを取得できること",
			args: args{
				ctx: ctx,
			},
			want: want{
				value: usecase.NewCalcService(),
				error: nil,
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			gotVal, gotErr := di(testcase.args.ctx)

			gotValue := reflect.ValueOf(gotVal).Elem()
			wantValue := reflect.ValueOf(testcase.want.value).Elem()

			for i := 0; i < gotValue.NumField(); i++ {
				gotType := gotValue.Field(i).Elem().Type().String()
				wantType := wantValue.Field(i).Elem().Type().String()
				if gotType != wantType {
					t.Errorf("di() = %v, want %v", gotType, wantType)
				}
			}
			if gotErr != testcase.want.error {
				t.Errorf("di(): error %v, want %v", gotErr, testcase.want.error)
			}
		})
	}
}
