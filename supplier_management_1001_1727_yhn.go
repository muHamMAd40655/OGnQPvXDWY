// 代码生成时间: 2025-10-01 17:27:04
package main

import (
    "os"
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/generators/assets/es"
    "github.com/gobuffalo/buffalo/generators/assets/genny"
    "github.com/gobuffalo/envy"
    "github.com/pkg/errors"
)

// 供应商模型
type Supplier struct {
    ID    uint   `db:""`
    Name  string `db:"size:255"`
    City  string `db:"size:255"`
    Email string `db:"size:255"`
}

// SupplierResource 定义供应商资源
type SupplierResource struct {
    // 可以添加额外的字段和方法
}

// NewSupplierResource 创建一个新的供应商资源
func NewSupplierResource(c buffalo.Context) buffalo.Value {
    // 从数据库加载供应商数据
    return c.Value("suppliers", suppliers)
}

// List 列出所有供应商
func (v SupplierResource) List(c buffalo.Context) error {
    // 获取查询参数
    query := c.Param("filter")
    // 执行查询
    var suppliers []Supplier
    if err := v.db.Where(query).All(&suppliers); err != nil {
        return errors.WithStack(err)
    }
    // 将查询结果添加到上下文
    return c.Render(200, r.JSON(suppliers))
}

// Show 显示供应商详情
func (v SupplierResource) Show(c buffalo.Context) error {
    // 从URL获取供应商ID
    id := c.Param("id")
    // 从数据库加载供应商数据
    var supplier Supplier
    if err := v.db.FindByID(id, &supplier); err != nil {
        return errors.WithStack(err)
    }
    // 将供应商数据添加到上下文
    return c.Render(200, r.JSON(supplier))
}

// Add 添加新供应商
func (v SupplierResource) Add(c buffalo.Context) error {
    // 解析请求体中的供应商数据
    var supplier Supplier
    if err := c.Bind(&supplier); err != nil {
        return errors.WithStack(err)
    }
    // 将供应商数据保存到数据库
    if err := v.db.Create(&supplier); err != nil {
        return errors.WithStack(err)
    }
    // 返回新创建的供应商数据
    return c.Render(201, r.JSON(supplier))
}

// Update 更新供应商信息
func (v SupplierResource) Update(c buffalo.Context) error {
    // 从URL获取供应商ID
    id := c.Param("id\)
    // 从数据库加载供应商数据
    var supplier Supplier
    if err := v.db.FindByID(id, &supplier); err != nil {
        return errors.WithStack(err)
    }
    // 解析请求体中的供应商数据
    if err := c.Bind(&supplier); err != nil {
        return errors.WithStack(err)
    }
    // 将更新后的供应商数据保存到数据库
    if err := v.db.Update(&supplier); err != nil {
        return errors.WithStack(err)
    }
    // 返回更新后的供应商数据
    return c.Render(200, r.JSON(supplier))
}

// Delete 删除供应商
func (v SupplierResource) Delete(c buffalo.Context) error {
    // 从URL获取供应商ID
    id := c.Param("id\)
    // 从数据库加载供应商数据
    var supplier Supplier
    if err := v.db.FindByID(id, &supplier); err != nil {
        return errors.WithStack(err)
    }
    // 将供应商数据从数据库删除
    if err := v.db.Delete(&supplier); err != nil {
        return errors.WithStack(err)
    }
    // 返回删除成功的响应
    return c.Render(200, r.JSON(map[string]string{
        "message": "Supplier deleted successfully",
    }))
}

func main() {
    // 设置环境变量
    envy.Load()
    // 初始化Buffalo应用
    app := buffalo.Automatic()
    // 设置数据库
    app.Middleware.Clear()
    app.Middleware.Use(middleware.DefaultLogger(middleware.LoggerOptions{}))
    app.Middleware.Use(middleware.DefaultRecovery(middleware.RecoveryOptions{}))
    app.Middleware.Use(middleware.Static{})
    app.Middleware.Use(middleware.CSRF)
    app.Middleware.Use(middleware.Session MiddleWareSession)
    // 定义路由
    app.Resource("/suppliers", SupplierResource{})
    // 启动应用
    if err := app.Serve(); err != nil {
        app.Stop(err)
    }
}
