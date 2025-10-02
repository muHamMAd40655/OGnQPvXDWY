// 代码生成时间: 2025-10-03 03:49:28
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/v2/pop"
    "github.com/markbates/gorp"
)

// ContentModel represents the model for a content entity
type ContentModel struct {
    ID   uint   "db:id"
    Title string "db:title"
    Body  string "db:body"
}

// ContentService contains the business logic for the content management system
type ContentService struct {
    db *gorp.DbMap
}

// NewContentService creates a new instance of the content service
func NewContentService(db *gorp.DbMap) *ContentService {
    return &ContentService{db: db}
}

// Create adds a new content to the database
func (s *ContentService) Create(title, body string) (*ContentModel, error) {
    c := ContentModel{Title: title, Body: body}
    err := s.db.Insert(&c)
    if err != nil {
        return nil, err
    }
    return &c, nil
}

// Update modifies an existing content in the database
func (s *ContentService) Update(id uint, title, body string) (*ContentModel, error) {
    c := ContentModel{ID: id, Title: title, Body: body}
    err := s.db.Update(&c)
    if err != nil {
        return nil, err
    }
    return &c, nil
}

// Delete removes a content from the database
func (s *ContentService) Delete(id uint) error {
    c := ContentModel{ID: id}
    err := s.db.Delete(&c)
    if err != nil {
        return err
    }
    return nil
}

// FindByID retrieves a content by its ID from the database
func (s *ContentService) FindByID(id uint) (*ContentModel, error) {
    c := ContentModel{ID: id}
    err := s.db.SelectOne(&c, 'SELECT * FROM contents WHERE id = ?', id)
    if err != nil {
        return nil, err
    }
    return &c, nil
}

// Main function to start the Buffalo application
func main() {
    app := buffalo.Automatic()

    // Define routes for the content management system
    app.GET("/", HomeHandler)
    app.POST("/content", ContentCreateHandler)
    app.PUT("/content/{id}", ContentUpdateHandler)
    app.DELETE("/content/{id}", ContentDeleteHandler)
    app.GET("/content/{id}", ContentReadHandler)

    // Run the application
    app.Serve()
}

// HomeHandler serves the home page
func HomeHandler(c buffalo.Context) error {
    return c.Render(200, r.HTML("index.html"))
}

// ContentCreateHandler handles the creation of new content
func ContentCreateHandler(c buffalo.Context) error {
    title := c.Request().FormValue("title")
    body := c.Request().FormValue("body")
    s := NewContentService(pop.Connection(c))
    _, err := s.Create(title, body)
    if err != nil {
        return c.Error(500, err)
    }
    return c.Redirect("/")
}

// ContentUpdateHandler handles the update of existing content
func ContentUpdateHandler(c buffalo.Context) error {
    id, _ := c.ParamValues("id")
    title := c.Request().FormValue("title")
    body := c.Request().FormValue("body")
    s := NewContentService(pop.Connection(c))
    _, err := s.Update(uint(id), title, body)
    if err != nil {
        return c.Error(500, err)
    }
    return c.Redirect("/")
}

// ContentDeleteHandler handles the deletion of existing content
func ContentDeleteHandler(c buffalo.Context) error {
    id, _ := c.ParamValues("id")
    s := NewContentService(pop.Connection(c))
    err := s.Delete(uint(id))
    if err != nil {
        return c.Error(500, err)
    }
    return c.Redirect("/")
}

// ContentReadHandler retrieves and displays a content item
func ContentReadHandler(c buffalo.Context) error {
    id, _ := c.ParamValues("id")
    s := NewContentService(pop.Connection(c))
    content, err := s.FindByID(uint(id))
    if err != nil {
        return c.Error(404, err)
    }
    return c.Render(200, r.JSON(content))
}