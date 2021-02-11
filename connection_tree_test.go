package main

import (
	"reflect"
	"testing"
)

func TestConnectionTree_Union(t *testing.T) {
	type fields struct {
		nodes      []int
		treeWeight []int
		maxWeight  int
	}
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name        	string
		fields      	fields
		args        	args
		expected    	bool
		wantErr     	bool
		expectedNodes	[]int
		expectedWeights []int
		expMaxWeight	int
	}{
		{
			"success 1",
			fields{
				nodes:      []int{0, 0, 0, 6, 1, 1, 6, 3, 6, 6, 8, 9, 9},
				treeWeight: []int{4, 2, 0, 1, 0, 0, 7, 0, 1, 2, 0, 0, 0},
				maxWeight:  7,
			},
			args{
				left:  0,
				right: 3,
			},
			true,
			false,
			[]int{6, 0, 0, 6, 1, 1, 6, 3, 6, 6, 8, 9, 9},
			[]int{4, 2, 0, 1, 0, 0, 12, 0, 1, 2, 0, 0, 0},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctr := &ConnectionTree{
				nodes:      tt.fields.nodes,
				treeWeight: tt.fields.treeWeight,
				maxWeight:  tt.fields.maxWeight,
			}
			got, err := ctr.Union(tt.args.left, tt.args.right)
			if (err != nil) != tt.wantErr {
				t.Errorf("Union() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("Union() got = %v, want %v", got, tt.expected)
				return
			}
			if !reflect.DeepEqual(tt.expectedNodes, ctr.nodes) {
				t.Errorf("ctr.nodes has wrong \nvalue = %v, \nwant  = %v", ctr.nodes, tt.expectedNodes)
				return
			}
			if !reflect.DeepEqual(tt.expectedWeights, ctr.treeWeight) {
				t.Errorf("ctr.treeWeight has wrong \nvalue = %v, \nwant  = %v", ctr.treeWeight, tt.expectedWeights)
				return
			}
			if tt.expMaxWeight != ctr.maxWeight {
				t.Errorf("ctr.maxWeight has wrong \nvalue = %v, \nwant  = %v", ctr.maxWeight, tt.expMaxWeight)
				return
			}
		})
	}
}
