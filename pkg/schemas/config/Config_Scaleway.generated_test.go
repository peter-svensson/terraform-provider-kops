package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/terraform-kops/terraform-provider-kops/pkg/api/config"
)

func TestExpandConfigScaleway(t *testing.T) {
	_default := config.Scaleway{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want config.Scaleway
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"access_key":            "",
					"secret_key":            "",
					"project_id":            "",
					"region":                "",
					"zone":                  "",
					"profile":               "",
					"kops_controller_image": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandConfigScaleway(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandConfigScaleway() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenConfigScalewayInto(t *testing.T) {
	_default := map[string]interface{}{
		"access_key":            "",
		"secret_key":            "",
		"project_id":            "",
		"region":                "",
		"zone":                  "",
		"profile":               "",
		"kops_controller_image": "",
	}
	type args struct {
		in config.Scaleway
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenConfigScalewayInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigScaleway() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
