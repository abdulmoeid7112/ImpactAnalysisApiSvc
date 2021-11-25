package mongo

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/abdulmoeid7112/impact-analysis-api-svc/db"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/models"
)

func Test_client_AddUser(t *testing.T) {
	os.Setenv("DB_HOST", "user-mongo-db")
	c, _ := NewMongoClient(db.Option{})
	tokenTime := time.Time{}
	user := &models.User{
		ID:            "123",
		Username:      "abdulmoeid",
		TokenIssuesAt: tokenTime,
	}

	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "success - add user in db",
			args: args{user: &models.User{
				ID:            "123",
				Username:      "abdulmoeid",
				TokenIssuesAt: tokenTime,
			}},
			want:    user,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.AddUser(context.TODO(), tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_GetUserByID(t *testing.T) {
	os.Setenv("DB_HOST", "user-mongo-db")
	c, _ := NewMongoClient(db.Option{})
	tokenTime := time.Time{}
	user := &models.User{
		ID:            "123",
		Username:      "abdulmoeid7112",
		TokenIssuesAt: tokenTime,
	}

	userAdded, _ := c.AddUser(context.TODO(), user)
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:    "success - get user by id",
			args:    args{name: userAdded.Username},
			want:    userAdded,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetUserByName(context.TODO(), tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_ListUser(t *testing.T) {
	os.Setenv("DB_HOST", "user-mongo-db")
	c, _ := NewMongoClient(db.Option{})
	tokenTime := time.Time{}
	user := &models.User{
		ID:            "123",
		Username:      "abdulmoeid7112",
		TokenIssuesAt: tokenTime,
	}
	tests := []struct {
		name    string
		want    []*models.User
		wantErr bool
	}{
		{
			name:    "success - get list of users",
			want:    []*models.User{user},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ListUser(context.TODO())
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
