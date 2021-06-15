package user

import (
	"testing"
)

func Test_userService_Validate(t *testing.T) {
	type args struct {
		user User
	}
	tests := []struct {
		name        string
		userService *userService
		args        args
		wantErr     bool
	}{
		{
			name:        "Given user, When name is ok, Then no error",
			userService: &userService{},
			args: args{
				user: User{
					ID:   1,
					Name: "Carlos",
				},
			},
			wantErr: false,
		},
		{
			name:        "Given user, When name is short, Then return error",
			userService: &userService{},
			args: args{
				user: User{
					ID:   1,
					Name: "Af",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{}
			if err := us.Validate(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userService.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
