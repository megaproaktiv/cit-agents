package starter

import (
	"testing"
	"github.com/aws/aws-sdk-go-v2/aws"

)

func TestGetLocalIp(t *testing.T) {
	tests := []struct {
		name string
		want *string
	}{
		{
			name: "Get local ip - change ip for test",
			want: aws.String("94.134.102.95"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetLocalIp(); *got != *tt.want {
				t.Errorf("GetLocalIp() = <%v>, want <%v>", *got, *tt.want)
			}
		})
	}
}
