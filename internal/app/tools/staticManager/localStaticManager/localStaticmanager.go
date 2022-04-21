package localStaticManager

import (
	"image"
	"os"

	"github.com/disintegration/imaging"
)

type LocalFileManager struct {
	staticUrl      string
	staticPath     string
	avatarPath     string
	restaurantPath string
	dishesPath     string
}

func NewLocalFileManager(staticUrl string, staticPath string) *LocalFileManager {
	return &LocalFileManager{
		staticUrl:      staticUrl,
		staticPath:     staticPath,
		avatarPath:     "avatar/",
		restaurantPath: "restaurants/",
		dishesPath:     "dishes/",
	}
}

func (f *LocalFileManager) SafeAvatar(img image.Image, imgName string) error {
	return imaging.Save(img, f.staticPath+f.avatarPath+imgName)
}

func (f *LocalFileManager) RemoveAvatar(avatarName string) error {
	return os.Remove(f.staticPath + f.avatarPath + avatarName)
}

func (f *LocalFileManager) IsNotSuchAvatarExist(avatarName string) bool {
	_, err := os.Stat(f.staticPath + f.avatarPath + avatarName)
	if err == nil {
		return false
	}
	return os.IsNotExist(err)
}

func (f *LocalFileManager) GetAvatarUrl(fileName string) string {
	return f.staticUrl + f.avatarPath + fileName
}

func (f *LocalFileManager) GetAvatarDirPath() string {
	return f.staticPath + f.avatarPath
}

func (f *LocalFileManager) GetRestaurantUrl(fileName string) string {
	return f.staticUrl + f.restaurantPath + fileName
}

func (f *LocalFileManager) GetDishesUrl(fileName string) string {
	return f.staticUrl + f.dishesPath + fileName
}
