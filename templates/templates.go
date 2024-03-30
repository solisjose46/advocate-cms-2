package templates

const (
	updateImageTitle = "Update an Image"
	updateImageEndpoint = "/update-image"
)

type imageGalleryItem struct {
	ImgID string
	ImgSrc string
	ImgAlt string
}

type crudImage struct {
	Title string
	Endpoint string
	ImageGalleryItems []imageGalleryItem
}