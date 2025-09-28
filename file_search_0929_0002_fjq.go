// 代码生成时间: 2025-09-29 00:02:58
 * Usage:
 *    file_search [directory]
 *
 * This program will start a web server that allows users to search and index files within the given directory.
 *
 * Note:
 *    Ensure that the directory provided is accessible and has the necessary permissions.
 */

package main

import (
    "os"
    "path/filepath"
    "strings"

    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/generators"
)

// Main is the entry point for the application
func main() {
    app := buffalo.Automatic()
    app.Serve()
}

// FileSearchHandler is a handler for searching and indexing files
type FileSearchHandler struct {
    //CTX is the context of the application
    CTX buffalo.Context
    //Log is the logger of the application
    Log buffalo.Logger
}

// IndexHandler is called when the user accesses the index page
func (f FileSearchHandler) IndexHandler(c buffalo.Context) error {
    // Get the directory path from the query parameter
    dirPath := c.Param("dir")
    if dirPath == "" {
        return c.Render(200, r.JSON(map[string]string{"error": "Directory path is required"}))
    }

    // Check if the directory exists
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        return c.Render(404, r.JSON(map[string]string{"error": "Directory not found"}))
    }

    // List files and directories within the given path
    files, err := filepath.Glob(filepath.Join(dirPath, "*.*"))
    if err != nil {
        return c.Render(500, r.JSON(map[string]string{"error": "Failed to list files"}))
    }

    // Return the list of files as JSON
    return c.Render(200, r.JSON(files))
}

// SearchHandler is called when the user submits a search query
func (f FileSearchHandler) SearchHandler(c buffalo.Context) error {
    // Get the directory path and search query from the query parameters
    dirPath := c.Param("dir")
    query := c.Param("query")
    if dirPath == "" || query == "" {
        return c.Render(400, r.JSON(map[string]string{"error": "Directory path and query are required"}))
    }

    // Check if the directory exists
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        return c.Render(404, r.JSON(map[string]string{"error": "Directory not found"}))
    }

    // Search for files that match the query
    files, err := filepath.Glob(filepath.Join(dirPath, "*"+query+"*.*"))
    if err != nil {
        return c.Render(500, r.JSON(map[string]string{"error": "Failed to search files"}))
    }

    // Return the list of matching files as JSON
    return c.Render(200, r.JSON(files))
}
