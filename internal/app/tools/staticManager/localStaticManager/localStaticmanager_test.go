package localStaticManager

import (
	"image"
	"reflect"
	"testing"
)

func TestLocalFileManager_GetAvatarDirPath(t *testing.T) {
	type fields struct {
		staticUrl         string
		staticPath        string
		avatarPath        string
		restaurantPath    string
		promocodePath     string
		logoPromocodePath string
		dishesPath        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LocalFileManager{
				staticUrl:         tt.fields.staticUrl,
				staticPath:        tt.fields.staticPath,
				avatarPath:        tt.fields.avatarPath,
				restaurantPath:    tt.fields.restaurantPath,
				promocodePath:     tt.fields.promocodePath,
				logoPromocodePath: tt.fields.logoPromocodePath,
				dishesPath:        tt.fields.dishesPath,
			}
			if got := f.GetAvatarDirPath(); got != tt.want {
				t.Errorf("GetAvatarDirPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalFileManager_GetAvatarUrl(t *testing.T) {
	type fields struct {
		staticUrl         string
		staticPath        string
		avatarPath        string
		restaurantPath    string
		promocodePath     string
		logoPromocodePath string
		dishesPath        string
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LocalFileManager{
				staticUrl:         tt.fields.staticUrl,
				staticPath:        tt.fields.staticPath,
				avatarPath:        tt.fields.avatarPath,
				restaurantPath:    tt.fields.restaurantPath,
				promocodePath:     tt.fields.promocodePath,
				logoPromocodePath: tt.fields.logoPromocodePath,
				dishesPath:        tt.fields.dishesPath,
			}
			if got := f.GetAvatarUrl(tt.args.fileName); got != tt.want {
				t.Errorf("GetAvatarUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalFileManager_GetDishesUrl(t *testing.T) {
	type fields struct {
		staticUrl         string
		staticPath        string
		avatarPath        string
		restaurantPath    string
		promocodePath     string
		logoPromocodePath string
		dishesPath        string
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LocalFileManager{
				staticUrl:         tt.fields.staticUrl,
				staticPath:        tt.fields.staticPath,
				avatarPath:        tt.fields.avatarPath,
				restaurantPath:    tt.fields.restaurantPath,
				promocodePath:     tt.fields.promocodePath,
				logoPromocodePath: tt.fields.logoPromocodePath,
				dishesPath:        tt.fields.dishesPath,
			}
			if got := f.GetDishesUrl(tt.args.fileName); got != tt.want {
				t.Errorf("GetDishesUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalFileManager_GetPromocodeLogoUrl(t *testing.T) {
	type fields struct {
		staticUrl         string
		staticPath        string
		avatarPath        string
		restaurantPath    string
		promocodePath     string
		logoPromocodePath string
		dishesPath        string
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LocalFileManager{
				staticUrl:         tt.fields.staticUrl,
				staticPath:        tt.fields.staticPath,
				avatarPath:        tt.fields.avatarPath,
				restaurantPath:    tt.fields.restaurantPath,
				promocodePath:     tt.fields.promocodePath,
				logoPromocodePath: tt.fields.logoPromocodePath,
				dishesPath:        tt.fields.dishesPath,
			}
			if got := f.GetPromocodeLogoUrl(tt.args.fileName); got != tt.want {
				t.Errorf("GetPromocodeLogoUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalFileManager_GetPromocodeUrl(t *testing.T) {
	type fields struct {
		staticUrl         string
		staticPath        string
		avatarPath        string
		restaurantPath    string
		promocodePath     string
		logoPromocodePath string
		dishesPath        string
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LocalFileManager{
				staticUrl:         tt.fields.staticUrl,
				staticPath:        tt.fields.staticPath,
				avatarPath:        tt.fields.avatarPath,
				restaurantPath:    tt.fields.restaurantPath,
				promocodePath:     tt.fields.promocodePath,
				logoPromocodePath: tt.fields.logoPromocodePath,
				dishesPath:        tt.fields.dishesPath,
			}
			if got := f.GetPromocodeUrl(tt.args.fileName); got != tt.want {
				t.Errorf("GetPromocodeUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalFileManager_GetRestaurantUrl(t *testing.T) {
	type fields struct {
		staticUrl         string
		staticPath        string
		avatarPath        string
		restaurantPath    string
		promocodePath     string
		logoPromocodePath string
		dishesPath        string
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LocalFileManager{
				staticUrl:         tt.fields.staticUrl,
				staticPath:        tt.fields.staticPath,
				avatarPath:        tt.fields.avatarPath,
				restaurantPath:    tt.fields.restaurantPath,
				promocodePath:     tt.fields.promocodePath,
				logoPromocodePath: tt.fields.logoPromocodePath,
				dishesPath:        tt.fields.dishesPath,
			}
			if got := f.GetRestaurantUrl(tt.args.fileName); got != tt.want {
				t.Errorf("GetRestaurantUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalFileManager_IsNotSuchAvatarExist(t *testing.T) {
	type fields struct {
		staticUrl         string
		staticPath        string
		avatarPath        string
		restaurantPath    string
		promocodePath     string
		logoPromocodePath string
		dishesPath        string
	}
	type args struct {
		avatarName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LocalFileManager{
				staticUrl:         tt.fields.staticUrl,
				staticPath:        tt.fields.staticPath,
				avatarPath:        tt.fields.avatarPath,
				restaurantPath:    tt.fields.restaurantPath,
				promocodePath:     tt.fields.promocodePath,
				logoPromocodePath: tt.fields.logoPromocodePath,
				dishesPath:        tt.fields.dishesPath,
			}
			if got := f.IsNotSuchAvatarExist(tt.args.avatarName); got != tt.want {
				t.Errorf("IsNotSuchAvatarExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalFileManager_RemoveAvatar(t *testing.T) {
	type fields struct {
		staticUrl         string
		staticPath        string
		avatarPath        string
		restaurantPath    string
		promocodePath     string
		logoPromocodePath string
		dishesPath        string
	}
	type args struct {
		avatarName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LocalFileManager{
				staticUrl:         tt.fields.staticUrl,
				staticPath:        tt.fields.staticPath,
				avatarPath:        tt.fields.avatarPath,
				restaurantPath:    tt.fields.restaurantPath,
				promocodePath:     tt.fields.promocodePath,
				logoPromocodePath: tt.fields.logoPromocodePath,
				dishesPath:        tt.fields.dishesPath,
			}
			if err := f.RemoveAvatar(tt.args.avatarName); (err != nil) != tt.wantErr {
				t.Errorf("RemoveAvatar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalFileManager_SafeAvatar(t *testing.T) {
	type fields struct {
		staticUrl         string
		staticPath        string
		avatarPath        string
		restaurantPath    string
		promocodePath     string
		logoPromocodePath string
		dishesPath        string
	}
	type args struct {
		img     image.Image
		imgName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &LocalFileManager{
				staticUrl:         tt.fields.staticUrl,
				staticPath:        tt.fields.staticPath,
				avatarPath:        tt.fields.avatarPath,
				restaurantPath:    tt.fields.restaurantPath,
				promocodePath:     tt.fields.promocodePath,
				logoPromocodePath: tt.fields.logoPromocodePath,
				dishesPath:        tt.fields.dishesPath,
			}
			if err := f.SafeAvatar(tt.args.img, tt.args.imgName); (err != nil) != tt.wantErr {
				t.Errorf("SafeAvatar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewLocalFileManager(t *testing.T) {
	type args struct {
		staticUrl  string
		staticPath string
	}
	tests := []struct {
		name string
		args args
		want *LocalFileManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLocalFileManager(tt.args.staticUrl, tt.args.staticPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLocalFileManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
