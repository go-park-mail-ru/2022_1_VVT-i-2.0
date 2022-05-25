package localStaticManager

import (
	"image"
	"os"
	"reflect"
	"testing"
)

func TestNewLocalFileManager(t *testing.T) {
	response := NewLocalFileManager("staticUrl", "staticPath")
	expect := &LocalFileManager{
		staticUrl:         "staticUrl",
		staticPath:        "staticPath",
		avatarPath:        "avatar/",
		restaurantPath:    "restaurants/",
		dishesPath:        "dishes/",
		promocodePath:     "promocodes/main/",
		logoPromocodePath: "promocodes/logos/",
	}
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestLocalFileManager_RemoveAvatar(t *testing.T) {
	FileManager := NewLocalFileManager("staticUrl", "staticPath")
	response := FileManager.RemoveAvatar("")
	expect := os.Remove(FileManager.staticPath + FileManager.avatarPath + "")
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestLocalFileManager_IsNotSuchAvatarExist(t *testing.T) {
	FileManager := NewLocalFileManager("staticUrl", "staticPath")
	response := FileManager.IsNotSuchAvatarExist("")
	expect := true
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestLocalFileManager_GetAvatarDirPath(t *testing.T) {
	FileManager := NewLocalFileManager("staticUrl", "staticPath")
	response := FileManager.GetAvatarDirPath()
	expect := FileManager.staticPath + FileManager.avatarPath
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestLocalFileManager_GetAvatarUrl(t *testing.T) {
	FileManager := NewLocalFileManager("staticUrl", "staticPath")
	response := FileManager.GetAvatarUrl("fileName")
	expect := FileManager.staticUrl + FileManager.avatarPath + "fileName"
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestLocalFileManager_GetAvatarUrl_Err(t *testing.T) {
	FileManager := NewLocalFileManager("staticUrl", "staticPath")
	response := FileManager.GetAvatarUrl("")
	expect := ""
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestLocalFileManager_GetDishesUrl(t *testing.T) {
	FileManager := NewLocalFileManager("staticUrl", "staticPath")
	response := FileManager.GetDishesUrl("")
	expect := FileManager.staticUrl + FileManager.dishesPath + ""
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestLocalFileManager_GetPromocodeLogoUrl(t *testing.T) {
	FileManager := NewLocalFileManager("staticUrl", "staticPath")
	response := FileManager.GetPromocodeLogoUrl("")
	expect := FileManager.staticUrl + FileManager.logoPromocodePath + ""
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestLocalFileManager_GetPromocodeUrl(t *testing.T) {
	FileManager := NewLocalFileManager("staticUrl", "staticPath")
	response := FileManager.GetPromocodeUrl("")
	expect := FileManager.staticUrl + FileManager.promocodePath + ""
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestLocalFileManager_GetRestaurantUrl(t *testing.T) {
	FileManager := NewLocalFileManager("staticUrl", "staticPath")
	response := FileManager.GetRestaurantUrl("")
	expect := FileManager.staticUrl + FileManager.restaurantPath + ""
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
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
