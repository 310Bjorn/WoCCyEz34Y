// 代码生成时间: 2025-10-07 19:24:32
package main

import (
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/golang/freetype/truetype"
    "github.com/golang/freetype"
    "github.com/markbates/buffalo"
    "github.com/otiai10/gosseract/v2"
)

// OCRService 结构体封装OCR服务的方法
type OCRService struct{
    // 可以在此定义更多的字段，如配置参数等
}

// NewOCRService 创建一个新的OCRService实例
func NewOCRService() *OCRService {
    return &OCRService{}
}

// ProcessImage 处理上传的图片文件并返回识别的文本
func (s *OCRService) ProcessImage(c buffalo.Context) error {
    // 获取上传的文件
    file, err := c.File("image")
    if err != nil {
        return err
    }
    defer file.Close()

    // 保存文件到临时目录
    tempFilePath := filepath.Join(os.TempDir(), file.Filename)
    if err := c.SaveFile(tempFilePath, file); err != nil {
        return err
    }
    defer os.Remove(tempFilePath) // 确保文件最终被删除

    // 使用gosseract进行OCR处理
    text, err := gosseract.OCR(tempFilePath, nil)
    if err != nil {
        return err
    }

    // 返回识别的文本
    c.Set("text", text)
    return c.Render(200, r.RenderOptions{
        ContentType: "text/plain",
        Data: struct{
            Text string
        }{
            Text: text,
        },
    })
}

// main 函数设置Buffalo路由并启动服务器
func main() {
    // 初始化Buffalo应用
    app := buffalo.New(buffalo.Options{})

    // 创建OCR服务
    ocrService := NewOCRService()

    // 定义路由处理POST请求和上传图片文件
    app.POST("/ocr", ocrService.ProcessImage)

    // 启动Buffalo应用
    app.Serve()
}
