package handlers
import (
    "github.com/GIT_USER_ID/GIT_REPO_ID/models"
    "github.com/labstack/echo/v4"
    "net/http"
)

// DeleteSecureVariable - Delete a secure variable bag
func (c *Container) DeleteSecureVariable(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}


// GetSecureVariable - Fetch a secure variables bag
func (c *Container) GetSecureVariable(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}


// ListVars - List the variable bags
func (c *Container) ListVars(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}


// UpsertSecureVariable - Store a secure variable bag
func (c *Container) UpsertSecureVariable(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}
