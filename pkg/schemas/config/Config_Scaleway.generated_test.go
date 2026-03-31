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
				in: config.Scaleway{},
			},
			want: _default,
		},
		{
			name: "AccessKey - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.AccessKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretKey - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.SecretKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProjectId - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.ProjectId = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Region - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.Region = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Zone - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.Zone = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Profile - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.Profile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KopsControllerImage - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.KopsControllerImage = ""
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

func TestFlattenConfigScaleway(t *testing.T) {
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
				in: config.Scaleway{},
			},
			want: _default,
		},
		{
			name: "AccessKey - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.AccessKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretKey - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.SecretKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ProjectId - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.ProjectId = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Region - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.Region = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Zone - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.Zone = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Profile - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.Profile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KopsControllerImage - default",
			args: args{
				in: func() config.Scaleway {
					subject := config.Scaleway{}
					subject.KopsControllerImage = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenConfigScaleway(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigScaleway() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
