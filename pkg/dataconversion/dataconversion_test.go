package dataconversion

import (
	"testing"
)

func TestNewDataConverter(t *testing.T) {
	tests := []struct {
		name string
		want *DataConverter
	}{
		{
			name: "Test NewDataConverter",
			want: &DataConverter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDataConverter(); got == nil {
				t.Errorf("NewDataConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataConverter_StructToJSONString(t *testing.T) {
	type fields struct {
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test StructToJSONString",
			args: args{
				data: struct {
					YourName string `json:"your_name" yaml:"your_name"`
				}{YourName: "Brett"},
			},
			want:    "{\"your_name\":\"Brett\"}\n",
			wantErr: false,
		},
		{
			name: "Test StructToJSONString with error",
			args: args{
				data: make(chan int),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &DataConverter{}
			got, err := dc.StructToJSONString(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataConverter.StructToJSONString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DataConverter.StructToJSONString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataConverter_StructToYAMLString(t *testing.T) {
	type fields struct {
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test StructToYAMLString",
			args: args{
				data: struct {
					YourName string `json:"your_name" yaml:"your_name"`
				}{YourName: "Brett"},
			},
			want:    "---\nyour_name: Brett\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &DataConverter{}
			got, err := dc.StructToYAMLString(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataConverter.StructToYAMLString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DataConverter.StructToYAMLString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataConverter_JSONStringToStruct(t *testing.T) {
	type structData struct {
		YourName string `json:"your_name" yaml:"your_name"`
	}

	type fields struct {
	}
	type args struct {
		structData *structData
		data       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test JSONStringToStruct",
			args: args{
				structData: &structData{},
				data:       "{\"your_name\":\"Brett\"}",
			},
			want: struct {
				YourName string `json:"your_name" yaml:"your_name"`
			}{YourName: "Brett"},
			wantErr: false,
		},
		{
			name: "Test JSONStringToStruct with error",
			args: args{
				structData: &structData{},
				data:       "stuff",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &DataConverter{}
			err := dc.JSONStringToStruct(tt.args.structData, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataConverter.JSONStringToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataConverter_YAMLStringToStruct(t *testing.T) {
	type structData struct {
		YourName string `json:"your_name" yaml:"your_name"`
	}

	type fields struct {
	}
	type args struct {
		structData *structData
		data       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test YAMLStringToStruct",
			args: args{
				structData: &structData{},
				data:       "---\nyour_name: Brett\n",
			},
			want: struct {
				YourName string `json:"your_name" yaml:"your_name"`
			}{YourName: "Brett"},
			wantErr: false,
		},
		{
			name: "Test YAMLStringToStruct with error",
			args: args{
				structData: &structData{},
				data:       "stuff",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &DataConverter{}
			err := dc.YAMLStringToStruct(tt.args.structData, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataConverter.YAMLStringToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
