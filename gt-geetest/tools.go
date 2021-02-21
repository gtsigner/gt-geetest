package geetest

import (
    "github.com/disintegration/imaging"
    "image"
    "image/color"
    "math"
)

var (
    geeOffsets = [][]int{{157, 145, 265, 277, 181, 169, 241, 253, 109, 97, 289, 301, 85, 73, 25, 37, 13, 1, 121, 133, 61, 49, 217, 229, 205, 193}, {145, 157, 277, 265, 169, 181, 253, 241, 97, 109, 301, 289, 73, 85, 37, 25, 1, 13, 133, 121, 49, 61, 229, 217, 193, 205}}
)

func recoverImgV2(src image.Image) (image.Image, error) {
    _, height := src.Bounds().Max.X, src.Bounds().Max.Y
    dst := imaging.New(260, height, color.NRGBA{})
    var imListUp []*image.NRGBA
    var imListDown []*image.NRGBA
    for y, loc := range geeOffsets {
        for _, x := range loc {
            if y == 0 {
                tmp := imaging.Crop(src, image.Rect(x, int(math.Floor(float64(height/2))), x+10, height))
                imListUp = append(imListUp, tmp)
            }
            if y == 1 {
                tmp := imaging.Crop(src, image.Rect(x, 0, x+10, int(math.Floor(float64(height/2)))))
                imListDown = append(imListDown, tmp)
            }
        }
    }
    off := 0
    for _, img := range imListUp {
        dst = imaging.Paste(dst, img, image.Pt(off, 0))
        off += img.Bounds().Max.X
    }
    off = 0
    for _, img := range imListDown {
        dst = imaging.Paste(dst, img, image.Pt(off, int(math.Floor(float64(height/2)))))
        off += img.Bounds().Max.X
    }
    return dst, nil
}

func geeOffsetV2(fgo image.Image, bgp image.Image) (int, error) {
    //1.先recover
    fg, err := recoverImgV2(fgo)
    if err != nil {
        return 0, err
    }

    bg, err := recoverImgV2(bgp)
    if err != nil {
        return 0, err
    }

    for i := 0; i < fg.Bounds().Max.X; i++ {
        for j := 0; j < fg.Bounds().Max.Y; j++ {
            if !pixelEqual(fg, bg, i, j) {
                return i, nil
            }
        }
    }
    return -1, ErrGetOffsetFail
}

func pixelEqual(img1 image.Image, img2 image.Image, x int, y int) bool {
    r1, g1, b1, _ := img1.At(x, y).RGBA()
    r2, g2, b2, _ := img2.At(x, y).RGBA()
    rAbs := math.Abs(float64(int(r1/257) - int(r2/257)))
    gAbs := math.Abs(float64(int(g1/257) - int(g2/257)))
    bAbs := math.Abs(float64(int(b1/257) - int(b2/257)))
    threshold := 60
    if int(rAbs) < threshold && int(gAbs) < threshold && int(bAbs) < threshold {
        return true
    }
    return false
}
