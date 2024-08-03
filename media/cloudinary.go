package media

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
)

func Config() (*cloudinary.Cloudinary, context.Context) {
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx
}
