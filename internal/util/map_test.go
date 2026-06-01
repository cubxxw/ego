package util

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMergeStringMap(t *testing.T) {
	type args struct {
		dest map[string]any
		src  map[string]any
		tar  map[string]any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "二维测试",
			args: args{
				dest: map[string]any{
					"2w": map[string]any{
						"test":  "2wtd",
						"test1": "2wtd1",
					},
					"2wa": map[string]any{
						"test":  "2wtd",
						"test1": "2wtd1",
					},
					"2wi": map[any]any{
						"test":  "2wtd",
						"test1": "2wtd1",
					},
				},
				src: map[string]any{
					"2w": map[string]any{
						"test":  "2wtds",
						"test1": "2wtd1s",
					},
					"2wb": map[string]any{
						"test":  "2wtds",
						"test1": "2wtd1s",
					},
					"2wi": map[any]any{
						"test":  "2wtds",
						"test1": "2wtd1s",
					},
				},
				tar: map[string]any{
					"2w": map[string]any{
						"test":  "2wtds",
						"test1": "2wtd1s",
					},
					"2wb": map[string]any{
						"test":  "2wtds",
						"test1": "2wtd1s",
					},
					"2wa": map[string]any{
						"test":  "2wtd",
						"test1": "2wtd1",
					},
					"2wi": map[string]any{
						"test":  "2wtds",
						"test1": "2wtd1s",
					},
				},
			},
		},
		{
			name: "一维测试",
			args: args{
				dest: map[string]any{
					"1w":  "tt",
					"1wa": "mq",
				},
				src: map[string]any{
					"1w":  "tts",
					"1wb": "bq",
				},
				tar: map[string]any{
					"1w":  "tts",
					"1wa": "mq",
					"1wb": "bq",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeStringMap(tt.args.dest, tt.args.src)
			if !reflect.DeepEqual(tt.args.dest, tt.args.tar) {
				spew.Dump(tt.args.dest)
				t.FailNow()
			}
		})
	}
}

func TestDeepSearchInMap(t *testing.T) {
	type args struct {
		m     map[string]any
		paths []string
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "test",
			args: args{map[string]any{"key1": map[string]any{"subkey1": "subval1"}}, []string{"key1"}},
			want: map[string]any{"subkey1": "subval1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeepSearchInMap(tt.args.m, tt.args.paths...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeepSearchInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapStringInterface(t *testing.T) {
	type args struct {
		src map[any]any
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "test",
			args: args{map[any]any{1: 1}},
			want: map[string]any{"1": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapStringInterface(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapStringInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeepSearchInMap1(t *testing.T) {
	type args struct {
		m     map[string]any
		paths []string
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeepSearchInMap(tt.args.m, tt.args.paths...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeepSearchInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeStringMap1(t *testing.T) {
	type args struct {
		dest map[string]any
		src  map[string]any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeStringMap(tt.args.dest, tt.args.src)
		})
	}
}

func TestToMapStringInterface1(t *testing.T) {
	type args struct {
		src map[any]any
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapStringInterface(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapStringInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
