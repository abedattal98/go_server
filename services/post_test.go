package services

import (
	"reflect"
	"rgb/interfaces"
	"rgb/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_postService_GetPostByID(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockPostRepo := NewMockPostRepository(ctrl)

	// postService := NewPostsService
	type fields struct {
		postRepo interfaces.PostRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func()
		want    models.Post
		wantErr bool
	}{
		{name: "post",
			fields: fields{
				postRepo: mockPostRepo,
			},
			args: args{
				id: 1,
			},
			setup: func() {
				posts := models.Post{
					ID:      1,
					Title:   "Gotham cronicles",
					Content: "Joker is planning big hit tonight.",
					UserID:  1,
				}
				mockPostRepo.EXPECT().GetPostByID(1).Return(posts, nil)
			},
			want: models.Post{
				Title:   "Gotham cronicles",
				Content: "Joker is planning big hit tonight.",
				UserID:  1,
				ID:      1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			p := NewPostsService(tt.fields.postRepo)
			got, err := p.GetPostByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("postService.GetPostByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.GetPostByID() = %v, want %v", got, tt.want)
			}
			ctrl.Finish()
		})
	}
}

func Test_postService_AddPost(t *testing.T) {

	user := randomUser()
	ctrl := gomock.NewController(t)
	mockPostRepo := NewMockPostRepository(ctrl)

	type fields struct {
		postRepo interfaces.PostRepository
	}
	type args struct {
		userID int
		post   models.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func()
		want    models.Post
		wantErr bool
	}{
		{name: "post",
			fields: fields{
				postRepo: mockPostRepo,
			},
			args: args{
				userID: 1,
				post: models.Post{
					Title:   "Gotham cronicles",
					Content: "Joker is planning big hit tonight.",
				},
			},
			setup: func() {
				posts := models.Post{
					ID:      1,
					Title:   "Gotham cronicles",
					Content: "Joker is planning big hit tonight.",
					UserID:  user.ID,
				}
				mockPostRepo.EXPECT().AddPost(gomock.Any(), gomock.Any()).Return(posts, nil)
			},
			want: models.Post{
				Title:   "Gotham cronicles",
				Content: "Joker is planning big hit tonight.",
				UserID:  user.ID,
				ID:      1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			p := NewPostsService(tt.fields.postRepo)
			got, err := p.AddPost(tt.args.userID, tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("postService.AddPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.AddPost() = %v, want %v", got, tt.want)
			}
		})
	}
}
func randomUser() models.User {
	return models.User{
		ID:       int(time.Now().UTC().Unix()) % 100,
		Username: "joker",
		Password: "joker",
		Email:    "jokw",
	}
}

func Test_postService_DeletePost(t *testing.T) {
	// user := randomUser()
	ctrl := gomock.NewController(t)
	mockPostRepo := NewMockPostRepository(ctrl)

	type fields struct {
		postRepo interfaces.PostRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func()
		wantErr bool
	}{
		{name: "post",
			fields: fields{
				postRepo: mockPostRepo,
			},
			args: args{
				id: 1,
			},
			setup: func() {
				mockPostRepo.EXPECT().DeletePost(gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			p := NewPostsService(tt.fields.postRepo)
			if err := p.DeletePost(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("postService.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postService_UpdatePost(t *testing.T) {
	user := randomUser()
	ctrl := gomock.NewController(t)
	mockPostRepo := NewMockPostRepository(ctrl)

	type fields struct {
		postRepo interfaces.PostRepository
	}
	type args struct {
		id   int
		post models.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		setup   func()
		want    models.Post
		wantErr bool
	}{
		{name: "post",
			fields: fields{
				postRepo: mockPostRepo,
			},
			args: args{
				id: user.ID,
				post: models.Post{
					Title:   "Gotham cronicles",
					Content: "Joker is planning big hit tonight.",
				},
			},
			setup: func() {
				posts := models.Post{
					ID:      1,
					Title:   "Gotham cronicles Batman",
					Content: "Joker is planning big hit tonight.",
					UserID:  user.ID,
				}
				mockPostRepo.EXPECT().UpdatePost(gomock.Any(), gomock.Any()).Return(posts, nil)
			},
			want: models.Post{
				Title:   "Gotham cronicles Batman",
				Content: "Joker is planning big hit tonight.",
				UserID:  user.ID,
				ID:      1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			p := NewPostsService(tt.fields.postRepo)
			got, err := p.UpdatePost(tt.args.id, tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("postService.UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.UpdatePost() = %v, want %v", got, tt.want)
			}
		})
	}
}
