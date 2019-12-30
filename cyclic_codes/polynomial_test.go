package cyclic_codes

import (
	"reflect"
	"testing"
)

func TestPolynomial_Divide(t *testing.T) {
	type fields struct {
		Base         int
		Degree       int
		Coefficients []int
	}
	type args struct {
		polDiv *Polynomial
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Polynomial
		want1   Polynomial
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				Base:         5,
				Degree:       2,
				Coefficients: []int{1, 0, 4},
			},
			args: args{
				polDiv: &Polynomial{
					Base:         5,
					Degree:       1,
					Coefficients: []int{1, 4},
				},
			},
			want: Polynomial{
				Base:         5,
				Degree:       1,
				Coefficients: []int{1, 1},
			},
			want1: Polynomial{
				Base:         5,
				Degree:       0,
				Coefficients: []int{0},
			},
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				Base:         5,
				Degree:       2,
				Coefficients: []int{1, 0, 1},
			},
			args: args{
				polDiv: &Polynomial{
					Base:         5,
					Degree:       1,
					Coefficients: []int{1, 4},
				},
			},
			want: Polynomial{
				Base:         5,
				Degree:       1,
				Coefficients: []int{1, 1},
			},
			want1: Polynomial{
				Base:         5,
				Degree:       0,
				Coefficients: []int{2},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pol := &Polynomial{
				Base:         tt.fields.Base,
				Degree:       tt.fields.Degree,
				Coefficients: tt.fields.Coefficients,
			}
			got, got1, err := pol.Divide(tt.args.polDiv)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divide() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Divide() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPolynomial_Multiply(t *testing.T) {
	type fields struct {
		Base         int
		Degree       int
		Coefficients []int
	}
	type args struct {
		polMul *Polynomial
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Polynomial
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				Base:         5,
				Degree:       2,
				Coefficients: []int{1, 1, 1},
			},
			args: args{
				polMul: &Polynomial{
					Base:         5,
					Degree:       1,
					Coefficients: []int{2},
				},
			},
			want: Polynomial{
				Base:         5,
				Degree:       3,
				Coefficients: []int{2, 2, 2, 0},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pol := &Polynomial{
				Base:         tt.fields.Base,
				Degree:       tt.fields.Degree,
				Coefficients: tt.fields.Coefficients,
			}
			got, err := pol.Multiply(tt.args.polMul)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() got = %v, want %v", got, tt.want)
			}
		})
	}
}
