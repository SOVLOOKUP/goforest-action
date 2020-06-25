package hello

import (
    "bytes"
    "encoding/base64"
    "github.com/gogf/gf/net/ghttp"
)

func init()  {

}

type ImageMsg struct {
    ID string
    Image string
    Msg string
}

type Image struct {
    ID string
    Image *bytes.Buffer
    Msg string
}

func (i *ImageMsg)base64ToImage() *Image {
    ddd, _ := base64.StdEncoding.DecodeString(i.Image) //成图片文件并把文件写入到buffer
    return &Image{
        i.ID,
        bytes.NewBuffer(ddd),
        i.Msg,
    }
}


// Hello is a demonstration route handler for output "Hello World!".
func Hello(r *ghttp.Request) {
    r.Response.Writeln("Hello World!")
}

func SendImage(r *ghttp.Request) *Image {
    var imagemsg *ImageMsg
    _ = r.Parse(&imagemsg)
    image := imagemsg.base64ToImage()
    return image
}