package staticManager

import "image"

const (
	RestaurantType = iota
	DishType
	AvatarType
)

type FileManager interface {
	RemoveAvatar(avatarName string) error
	SafeAvatar(img image.Image, imgName string) error
	GetAvatarUrl(fileName string) string
	GetAvatarDirPath() string
	IsNotSuchAvatarExist(avatarName string) bool
	GetRestaurantUrl(fileName string) string
	GetDishesUrl(fileName string) string
	GetPromocodeUrl(fileName string) string
	GetPromocodeLogoUrl(fileName string) string
}
