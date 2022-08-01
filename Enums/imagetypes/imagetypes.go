//go:generate ..\go-enum -f=$GOFILE --noprefix --names --lower

package imagetypes

// ENUM(Unknown, Jpeg, Jpg, Png, Tiff, Gif)
type ImageType int
