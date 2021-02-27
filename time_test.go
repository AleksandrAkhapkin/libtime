package libtime

import (
	"database/sql/driver"
	"reflect"
	"testing"
	"time"
)

func MustParse(v string) time.Time {
	t, _ := time.Parse(ISO8601, v)
	return t
}

func TestTime_MarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Time: MustParse("15:10-03:00"),
			},
			want:    []byte(`"15:10-03:00"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t1 *testing.T) {
				t := Time{
					Time: tt.fields.Time,
				}
				got, err := t.MarshalJSON()
				if (err != nil) != tt.wantErr {
					t1.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t1.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				b: []byte(`"15:10-03:00"`),
			},
			wantErr: nil,
		},
		{
			name: "parse error",
			args: args{
				b: []byte(`"15:10-03"`),
			},
			wantErr: ErrTimeParse,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t1 *testing.T) {
				t := &Time{}
				if err := t.UnmarshalJSON(tt.args.b); err != tt.wantErr {
					t1.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestTime_Scan(t *testing.T) {
	type args struct {
		src interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				src: []byte(`15:10-03:00`),
			},
			wantErr: false,
		},
		{
			name: "parse error",
			args: args{
				src: []byte(`15:10-03`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t1 *testing.T) {
				t := &Time{}
				if err := t.Scan(tt.args.src); (err != nil) != tt.wantErr {
					t1.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestTime_Value(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Time: MustParse("15:10-03:00"),
			},
			want:    "15:10-03:00",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t1 *testing.T) {
				t := Time{
					Time: tt.fields.Time,
				}
				got, err := t.Value()
				if (err != nil) != tt.wantErr {
					t1.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t1.Errorf("Value() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
