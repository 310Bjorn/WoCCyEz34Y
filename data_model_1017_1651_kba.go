// 代码生成时间: 2025-10-17 16:51:37
package main

import (
    "buffalo/buffalo"
    "github.com/buffalo/buffalo-pop/v2/pop/popmw"
    "github.com/gobuffalo/pop/v5"
    "github.com/gobuffalo/uuid"
    "github.com/markbates/validate"
)

// User represents the model for a user entity.
// It includes fields for ID, name, email, and password.
// It also includes methods for validating the user data.
type User struct {
    ID        uuid.UUID `db:"id"`
    Name      string   `db:"name"`
    Email     string   `db:"email"`
    Password  string   `db:"password"`
    CreatedAt string   `db:"created_at"`
    UpdatedAt string   `db:"updated_at"`
}

// Validate gets run every time you call .Save(), .Create(), or .Update() on a pop.Model.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
    var errors validate.Errors
    if u.Name == "" {
        errors["name"] = append(errors["name"], validate.Error{Field: "name", ActualTag: "required", Message: "Name is required"})
    }
    if u.Email == "" {
        errors["email"] = append(errors["email"], validate.Error{Field: "email", ActualTag: "required", Message: "Email is required"})
    }
    // Additional validation logic can be added here.
    return &errors, nil
}

// UsersResource is a resource for handling user operations.
type UsersResource struct{}

// List responds to a GET request for all users.
func (u UsersResource) List(c buffalo.Context) error {
    // Query the database for all users.
    users := []User{}
    err := c.Value("db").(*pop.Connection).Where("1 = 1").All(&users)
    if err != nil {
        return buffalo.NewError(err, 500)
    }
    return c.Render(200, r.JSON(users))
}

// Show responds to a GET request for a single user.
func (u UsersResource) Show(c buffalo.Context) error {
    // Get the user ID from the URL parameter.
    id := c.Param("id")
    user := &User{}
    err := c.Value("db").(*pop.Connection).Find(id, user)
    if err != nil {
        return buffalo.NewError(err, 500)
    }
    return c.Render(200, r.JSON(user))
}

// Create adds a new user to the database.
func (u UsersResource) Create(c buffalo.Context) error {
    // Decode the user data from the request body.
    user := &User{}
    if err := c.Bind(user); err != nil {
        return buffalo.NewError(err, 400)
    }
    // Validate the user data.
    if err := c.Validate(user); err != nil {
        return buffalo.NewError(err, 400)
    }
    // Save the user to the database.
    err := c.Value("db").(*pop.Connection).Create(user)
    if err != nil {
        return buffalo.NewError(err, 500)
    }
    return c.Render(201, r.JSON(user))
}
