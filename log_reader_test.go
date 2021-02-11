package main

import "testing"

func TestLogReader_FindInterconnection(t *testing.T) {
	type fields struct {
		usersCount	int
		file		string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"case 1",
			fields{
				usersCount: 13,
				file:       "testcases/find_interconnection_1.log",
			},
			"2021-02-12T12:45:27",
			false,
		},
		{
			"case 1 err",
			fields{
				usersCount: 13,
				file:       "testcases/find_interconnection_err.log",
			},
			"",
			true,
		},
		{
			"case  err",
			fields{
				usersCount: 13,
				file:       "testcases/find_interconnection_err_2.log",
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdr := InitLogReader(tt.fields.file, tt.fields.usersCount)
			got, err := rdr.FindInterconnection()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindInterconnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindInterconnection() got = %v, want %v", got, tt.want)
			}
		})
	}
}
