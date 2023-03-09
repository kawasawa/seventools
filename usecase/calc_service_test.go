package usecase

import (
	"context"
	"go-onion-sample/domain/entity"
	"reflect"
	"testing"
)

func Test_NewCalcService(t *testing.T) {
	for _, testcase := range []struct {
		name string
		want *CalcService
	}{
		{
			name: "インスタンスが生成されること",
			want: &CalcService{},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			got := NewCalcService()

			gotValue := reflect.ValueOf(got).Elem()
			wantValue := reflect.ValueOf(testcase.want).Elem()

			for i := 0; i < gotValue.NumField(); i++ {
				gotType := gotValue.Field(i).Elem().Type().String()
				wantType := wantValue.Field(i).Elem().Type().String()
				if gotType != wantType {
					t.Errorf("NewCalcService() = %v, want %v", gotType, wantType)
				}
			}
		})
	}
}

func Test_CalcService_Add(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx    context.Context
		params *entity.Params
	}
	type want struct {
		value int
		error error
	}

	for _, testcase := range []struct {
		name string
		args args
		want want
	}{
		{
			name: "引数で渡された数値を加算した値が返却されること",
			args: args{
				ctx:    ctx,
				params: &entity.Params{A: 100, B: 50},
			},
			want: want{
				value: 150,
				error: nil,
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			calcService := NewCalcService()
			gotVal, gotErr := calcService.Add(testcase.args.ctx, testcase.args.params)

			if !reflect.DeepEqual(gotVal, testcase.want.value) {
				t.Errorf("CalcService.Add(): got %v, want %v", gotVal, testcase.want.value)
			}
			if gotErr != testcase.want.error {
				t.Errorf("CalcService.Add(): error %v, want %v", gotErr, testcase.want.error)
			}
		})
	}
}

func Test_CalcService_Subtract(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx    context.Context
		params *entity.Params
	}
	type want struct {
		value int
		error error
	}

	for _, testcase := range []struct {
		name string
		args args
		want want
	}{
		{
			name: "引数で渡された数値を減算した値が返却されること",
			args: args{
				ctx:    ctx,
				params: &entity.Params{A: 100, B: 50},
			},
			want: want{
				value: 50,
				error: nil,
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			calcService := NewCalcService()
			gotVal, gotErr := calcService.Subtract(testcase.args.ctx, testcase.args.params)

			if !reflect.DeepEqual(gotVal, testcase.want.value) {
				t.Errorf("CalcService.Subtract(): got %v, want %v", gotVal, testcase.want.value)
			}
			if gotErr != testcase.want.error {
				t.Errorf("CalcService.Subtract(): error %v, want %v", gotErr, testcase.want.error)
			}
		})
	}
}
