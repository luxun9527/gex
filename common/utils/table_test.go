package utils

import (
	"testing"
)

func TestWithShardingSuffix(t *testing.T) {
	type args struct {
		tableName string
		userId    int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				tableName: "t1",
				userId:    0,
			},
			want: "t1_00",
		}, {
			name: "t2",
			args: args{
				tableName: "t2",
				userId:    44,
			},
			want: "t2_04",
		}, {
			name: "t3",
			args: args{
				tableName: "t3",
				userId:    8,
			},
			want: "t3_08",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithShardingSuffix(tt.args.tableName, tt.args.userId); got != tt.want {
				t.Errorf("WithShardingSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
