// 代码生成时间: 2025-10-06 20:40:41
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/v2/pop"
    "log"
)

// LicenseModel 定义许可证模型
type LicenseModel struct {
    ID        uint   `db:"id"`
    Name      string `db:"name"`
    Key       string `db:"key"`
    CreatedAt string `db:"created_at"`
    UpdatedAt string `db:"updated_at"`
}

// LicenseResource 定义许可证资源
type LicenseResource struct{
    // Standard Buffalo resource model
    // The `buffalo.BaseResource` struct will be embedded into this struct
    // by default, so you don't need to include it explicitly.
    // Add common methods here that can be used across all actions.
}

// NewLicenseResource 构造一个新的许可证资源
func NewLicenseResource(c buffalo.Context) *LicenseResource {
    return &LicenseResource{}
}

// List 列出所有许可证
func (l *LicenseResource) List(c buffalo.Context) error {
    var licenses []LicenseModel
    // 使用 Pop 进行数据库查询
    if err := pop.Q().All(&licenses); err != nil {
        return err
    }
    // 将查询结果设置到 context 中，以便在模板中使用
    c.Set("licenses", licenses)
    return c.Render(200, r.HTML("licenses/index.html"))
}

// Show 显示指定许可证的详细信息
func (l *LicenseResource) Show(c buffalo.Context) error {
    // 从 URL 中获取许可证 ID
    id := c.Param("id")
    // 通过 ID 查找许可证
    var license LicenseModel
    if err := pop.Find(&license, id); err != nil {
        return err
    }
    // 将许可证设置到 context 中，以便在模板中使用
    c.Set("license\, license)
    return c.Render(200, r.HTML("licenses/show.html"))
}

// Create 添加一个新的许可证
func (l *LicenseResource) Create(c buffalo.Context) error {
    // 解析请求中的表单数据
    if err := c.Bind(&license); err != nil {
        return err
    }
    // 保存许可证到数据库
    if err := pop.Create(&license); err != nil {
        return err
    }
    // 重定向到许可证列表页面
    return c.Redirect(302, "/licenses")
}

// Update 更新现有许可证
func (l *LicenseResource) Update(c buffalo.Context) error {
    // 从 URL 中获取许可证 ID
    id := c.Param("id")
    // 查找许可证
    var license LicenseModel
    if err := pop.Find(&license, id); err != nil {
        return err
    }
    // 解析请求中的表单数据
    if err := c.Bind(&license); err != nil {
        return err
    }
    // 更新许可证到数据库
    if err := pop.Update(&license); err != nil {
        return err
    }
    // 重定向到许可证列表页面
    return c.Redirect(302, "/licenses")
}

// Delete 删除许可证
func (l *LicenseResource) Delete(c buffalo.Context) error {
    // 从 URL 中获取许可证 ID
    id := c.Param("id")
    // 查找并删除许可证
    var license LicenseModel
    if err := pop.Find(&license, id); err != nil {
        return err
    }
    if err := pop.Destroy(&license); err != nil {
        return err
    }
    // 重定向到许可证列表页面
    return c.Redirect(302, "/licenses")
}

func main() {
    // 初始化 Buffalo
    app := buffalo.Automatic()
    
    // 注册许可证资源
    app.Resource("/licenses", NewLicenseResource(),
        buffalo.Handlers(
            buffalo.AssetsBoxer("/assets/"),
            buffalo.SessionReporter,
            buffalo.TemplateRenderer,
        ),
    )
    
    // 启动 Buffalo 应用
    app.Serve()
}