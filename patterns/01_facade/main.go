package main

import "log"

type imgType byte

const (
	ImgProfile imgType = iota
	ImgCover
)

func main() {
	facade := ImageFacade{}
	if err := facade.UploadImage("profile.jpg", ImgProfile); err != nil {
		log.Fatal(err)
	}
	if err := facade.UploadImage("cover.png", ImgCover); err != nil {
		log.Fatal(err)
	}
}

type ImageFacade struct{}

func (f *ImageFacade) UploadImage(filename string, itype imgType) error {
	imgs := new(images)
	srvc := new(service)
	img, err := imgs.Upload(filename)
	if err != nil {
		return err
	}
	webp, err := imgs.ConvertWEBP(img)
	if err != nil {
		return err
	}
	switch itype {
	case ImgProfile:
		imgS, err := imgs.Resize(webp, 64, 64)
		if err != nil {
			return err
		}
		imgL, err := imgs.Resize(webp, 256, 256)
		if err != nil {
			return err
		}
		srvc.SaveImage(imgS)
		srvc.SaveImage(imgL)
		log.Printf("Profile image %s succesfully uploaded\n", filename)
	case ImgCover:
		cover, err := imgs.Resize(webp, 1280, 320)
		if err != nil {
			return err
		}
		srvc.SaveImage(cover)
		log.Printf("Cover image %s succesfully uploaded\n", filename)
	}
	return nil
}

type images struct{}

func (i *images) Upload(filename string) ([]byte, error) {
	return []byte{'I', 'm', 'a', 'g', 'e'}, nil
}

func (i *images) Resize(image []byte, w, h int) ([]byte, error) {
	return []byte{'I', 'm', 'a', 'g', 'e'}, nil
}

func (i *images) ConvertWEBP(image []byte) ([]byte, error) {
	return []byte{'I', 'm', 'a', 'g', 'e'}, nil
}

type service struct{}

func (s *service) SaveImage(image []byte) {
}
