// 代码生成时间: 2025-10-12 23:29:50
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop"
    "github.com/markbates/inflect"
    "log"
    "net/http"
)

// FirmwareUpdateModel 定义设备固件更新信息模型
type FirmwareUpdateModel struct {
    ID       uint   `db:"id"`
    DeviceID string `db:"device_id"`
    Version  string `db:"version"`
    CreatedAt time  `db:"created_at"`
    UpdatedAt time  `db:"updated_at"`
}

// FirmwareUpdateResource 定义固件更新资源
type FirmwareUpdateResource struct {
    Model FirmwareUpdateModel
    Trans  buffalo.Translation
    Validator
}

// NewFirmwareUpdateResource 创建一个新的固件更新资源
func NewFirmwareUpdateResource() *FirmwareUpdateResource {
    return &FirmwareUpdateResource{
        Validator: &validate.FM{}
    }
}

// CreateHandler 处理固件更新创建请求
func (res *FirmwareUpdateResource) CreateHandler(c buffalo.Context) error {
    var m FirmwareUpdateModel
    if err := c.Request().ParseForm(); err != nil {
        return handleError(err, http.StatusBadRequest)
    }
    if err := decodeJSON(c.Request().Body, &m); err != nil {
        return handleError(err, http.StatusBadRequest)
    }
    if err := res.ValidateAndCreate(c.Transaction(), &m); err != nil {
        return handleError(err, http.StatusInternalServerError)
    }
    return c.Render(http.StatusCreated, r.JSON(m))
}

// ValidateAndCreate 验证并创建固件更新记录
func (res *FirmwareUpdateResource) ValidateAndCreate(t *pop.Connection, m *FirmwareUpdateModel) error {
    if err := t.Where("device_id = ? AND version = ?", m.DeviceID, m.Version).Eager().First(m); err != nil {
        if err == sql.ErrNoRows {
            if err := t.Create(m); err != nil {
                return err
            }
        } else {
            return err
        }
    } else {
        return errors.New(res.Trans.Translate(c, "firmware_version_exists"))
    }
    return nil
}

// handleError 处理错误并返回错误响应
func handleError(err error, statusCode int) error {
    log.Printf("Error: %v", err)
    return c.Error(statusCode, err)
}

func main() {
    app := buffalo.Automatic()
    app.Resource("/firmware-updates", NewFirmwareUpdateResource())
    app.Serve()
}
