package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackjack/webcam"
)

const (
	FMT_MJPEG = webcam.PixelFormat(0x47504A4D) // 'MJPG'
	FMT_YUYV  = webcam.PixelFormat(0x56595559) // 'YUYV'

	asciiChars   = " .,:;i1tfLCG08@"
	targetWidth  = 160
	targetHeight = 50
)

var asciiRamp = []rune(asciiChars)

func pixelToASCII(v uint8) rune {
	idx := int(float64(v) / 255.0 * float64(len(asciiRamp)-1))
	if idx >= len(asciiRamp) {
		idx = len(asciiRamp) - 1
	}
	return asciiRamp[idx]
}

func imageToGray(img image.Image) *image.Gray {
	b := img.Bounds()
	g := image.NewGray(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, gr, bl, _ := img.At(x, y).RGBA()
			lum := 0.299*float64(r>>8) + 0.587*float64(gr>>8) + 0.114*float64(bl>>8)
			g.SetGray(x, y, color.Gray{Y: uint8(lum)})
		}
	}
	return g
}

func resample(src *image.Gray, w, h int) *image.Gray {
	sb := src.Bounds()
	sw, sh := sb.Max.X-sb.Min.X, sb.Max.Y-sb.Min.Y
	dst := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			sx := sb.Min.X + int(math.Round(float64(x)*float64(sw)/float64(w)))
			sy := sb.Min.Y + int(math.Round(float64(y)*float64(sh)/float64(h)))
			if sx >= sb.Max.X {
				sx = sb.Max.X - 1
			}
			if sy >= sb.Max.Y {
				sy = sb.Max.Y - 1
			}
			dst.SetGray(x, y, src.GrayAt(sx, sy))
		}
	}
	return dst
}

func renderASCII(g *image.Gray) string {
	b := g.Bounds()
	buf := make([]byte, 0, (b.Max.X-b.Min.X)*2*(b.Max.Y-b.Min.Y+1))
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			ch := byte(pixelToASCII(g.GrayAt(x, y).Y))
			buf = append(buf, ch, ch)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

func clamp(v int) int {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return v
}

func yuyvToImage(data []byte, w, h int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for i := 0; i+3 < len(data); i += 4 {
		y0 := int(data[i])
		u := int(data[i+1]) - 128
		y1 := int(data[i+2])
		v := int(data[i+3]) - 128
		pixel := func(y int) color.RGBA {
			return color.RGBA{
				uint8(clamp(y + 1370705*v/1000000)),
				uint8(clamp(y - 337633*u/1000000 - 698001*v/1000000)),
				uint8(clamp(y + 1732446*u/1000000)),
				255,
			}
		}
		px := (i / 4) * 2
		x0, ypos := px%w, px/w
		if ypos < h {
			img.SetRGBA(x0, ypos, pixel(y0))
			if x0+1 < w {
				img.SetRGBA(x0+1, ypos, pixel(y1))
			}
		}
	}
	return img
}

var probeResolutions = [][2]uint32{
	{1280, 720}, {640, 480}, {960, 540}, {320, 240}, {1920, 1080},
}

func setupCamera(cam *webcam.Webcam) (format webcam.PixelFormat, w, h uint32, err error) {
	for _, fmt_ := range []webcam.PixelFormat{FMT_MJPEG, FMT_YUYV} {
		for _, res := range probeResolutions {
			var fw, fh uint32
			_, fw, fh, err = cam.SetImageFormat(fmt_, res[0], res[1])
			if err == nil {
				return fmt_, fw, fh, nil
			}
		}
	}
	return 0, 0, 0, fmt.Errorf("no working format/resolution found")
}

func main() {
	devicePath := "/dev/video0"
	if len(os.Args) > 1 {
		devicePath = os.Args[1]
	}

	cam, err := webcam.Open(devicePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open %s: %v\n", devicePath, err)
		fmt.Fprintln(os.Stderr, "Usage: ascii-camera [/dev/videoN]  (default: /dev/video0)")
		os.Exit(1)
	}
	defer cam.Close()

	format, w, h, err := setupCamera(cam)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Camera setup failed: %v\n", err)
		os.Exit(1)
	}

	fmtName := map[webcam.PixelFormat]string{FMT_MJPEG: "MJPEG", FMT_YUYV: "YUYV"}[format]
	fmt.Fprintf(os.Stderr, "Camera ready: %s %dx%d on %s\n", fmtName, w, h, devicePath)

	if err := cam.StartStreaming(); err != nil {
		fmt.Fprintf(os.Stderr, "StartStreaming: %v\n", err)
		os.Exit(1)
	}
	defer cam.StopStreaming()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Print("\033[?25h\n")
		fmt.Fprintln(os.Stderr, "Bye!")
		os.Exit(0)
	}()

	fmt.Print("\033[?25l") // hide cursor
	fmt.Print("\033[2J")   // clear screen once
	defer fmt.Print("\033[?25h")

	ticker := time.NewTicker(50 * time.Millisecond) // ~20 fps
	defer ticker.Stop()

	for range ticker.C {
		if err := cam.WaitForFrame(2); err != nil {
			continue
		}
		frameData, err := cam.ReadFrame()
		if err != nil || len(frameData) == 0 {
			continue
		}

		var img image.Image
		if format == FMT_MJPEG {
			img, _, err = image.Decode(bytes.NewReader(frameData))
			if err != nil {
				continue
			}
		} else {
			img = yuyvToImage(frameData, int(w), int(h))
		}

		gray := imageToGray(img)
		small := resample(gray, targetWidth, targetHeight)
		fmt.Print("\033[H")
		fmt.Print(renderASCII(small))
	}
}