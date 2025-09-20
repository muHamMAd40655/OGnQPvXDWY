// 代码生成时间: 2025-09-20 10:26:25
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop"
    "github.com/gobuffalo/envy"
    "log"
)

// Inventory 库存模型
type Inventory struct {
    ID       uint   "db:"id""
    Name     string "db:"name""
    Quantity int    "db:"quantity""
}

// InventoryResource 库存资源
type InventoryResource struct {
    *buffalo.Resource
    DB *pop.Connection
}

// NewInventoryResource 初始化库存资源
func NewInventoryResource() *InventoryResource {
    r := buffalo.Resource("inventory", 0)
    r.Validator = buffalo.DefaultValidator
    r允许使用.PopDB = pop.Connect(envy.Get("DATABASE_URL"))
    return &InventoryResource{Resource: r, DB: r允许使用.PopDB}
}

// List 列出所有库存项
func (r *InventoryResource) List(c buffalo.Context) error {
    var items []Inventory
    if err := r.DB.All(&items); err != nil {
        return err
    }
    return c.Render(200, buffalo.R.JSON(items))
}

// Show 显示单个库存项
func (r *InventoryResource) Show(c buffalo.Context) error {
    // 检索单个库存项
    obj := Inventory{}
    if err := r.PopulateObj(&obj, c.Request().URL.Query(), r.DB); err != nil {
        return err
    }
    return c.Render(200, buffalo.R.JSON(obj))
}

// Create 创建新的库存项
func (r *InventoryResource) Create(c buffalo.Context) error {
    var item Inventory
    if err := c.Bind(&item); err != nil {
        return err
    }
    if err := r.DB.Create(&item); err != nil {
        return err
    }
    return c.Render(201, buffalo.R.JSON(item))
}

// Update 更新库存项
func (r *InventoryResource) Update(c buffalo.Context) error {
    var item Inventory
    if err := r.PopulateObj(&item, c.Request().URL.Query(), r.DB); err != nil {
        return err
    }
    if err := c.Bind(&item); err != nil {
        return err
    }
    if err := r.DB.Update(&item); err != nil {
        return err
    }
    return c.Render(200, buffalo.R.JSON(item))
}

// Destroy 删除库存项
func (r *InventoryResource) Destroy(c buffalo.Context) error {
    var item Inventory
    if err := r.PopulateObj(&item, c.Request().URL.Query(), r.DB); err != nil {
        return err
    }
    if err := r.DB.Destroy(&item); err != nil {
        return err
    }
    return c.Render(200, buffalo.R.Empty())
}

// main 函数，Buffalo应用的入口点
func main() {
    app := buffalo.Automatic()

    // 注册库存资源
    app.Resource("/inventory", NewInventoryResource())

    // 启动应用
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }
}
