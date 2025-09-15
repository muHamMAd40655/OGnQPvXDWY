// 代码生成时间: 2025-09-15 21:35:51
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/meta/ast"
    "log"
)

// dataModelGenerator 是一个用于生成数据模型的 generator
type dataModelGenerator struct {
    generators.DefaultGenerator
}

// 生成的数据模型文件路径
const dataModelPath = "models/model.go"

// NewDataModelGenerator 构造一个数据模型 generator
func NewDataModelGenerator() *dataModelGenerator {
    return &dataModelGenerator{}
}

// Generate 用于生成数据模型
func (g *dataModelGenerator) Generate(opts *ast.AppOptions) error {
    // 检查 appOptions 是否为空
    if opts == nil {
        return generators.ErrNoAppOpts
    }

    // 创建一个新的数据模型文件
    if err := g.CreateFile(dataModelPath, modelTemplate(opts)); err != nil {
        return err
    }

    return nil
}

// modelTemplate 用于生成数据模型的模板
func modelTemplate(opts *ast.AppOptions) []byte {
    // 定义数据模型的结构体
    model := `
// Model represents a data model
type Model struct {
    ID   uint   "db:""json:"id"""
    Name string "db:"name"json:"name""
}

// TableName 定义数据模型的表名
func (Model) TableName() string {
    return "models"
}
`
    return []byte(model)
}

// 主函数，用于初始化 Buffalo 应用并注册数据模型 generator
func main() {
    // 初始化 Buffalo 应用
    app := buffalo.NewApp(
        buffalo.Options{
            PreRunner: func(cmd *buffalo.Command) error {
                // 注册数据模型 generator
                cmd.RegisterCustomGenerator("model", NewDataModelGenerator())
                return nil
            },
        },
    )

    // 启动 Buffalo 应用
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }
}