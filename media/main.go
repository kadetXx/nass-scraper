package media

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Cloud struct {
	Cld *cloudinary.Cloudinary
	Ctx context.Context
}

func (c *Cloud) Upload(url string) string {
	resp, err := c.Cld.Upload.Upload(c.Ctx, url, uploader.UploadParams{
		Folder: "nass-scraper",
	})

	if err != nil {
		fmt.Println(err)
	}

	return resp.SecureURL
}
