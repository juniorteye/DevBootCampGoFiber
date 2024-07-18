// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bootcamp": {
            "post": {
                "description": "Create a new BootCamp",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BootCamps"
                ],
                "summary": "Create a new BootCamp",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Photo",
                        "name": "photo",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Bootcamp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        },
        "/bootcamp/{bootcampId}": {
            "get": {
                "description": "Get a BootCamp by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BootCamps"
                ],
                "summary": "Get a BootCamp by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "BootCamp ID",
                        "name": "bootcampId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Bootcamp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a BootCamp by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BootCamps"
                ],
                "summary": "Update a BootCamp by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "BootCamp ID",
                        "name": "bootcampId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "BootCamp",
                        "name": "bootcamp",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Bootcamp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Bootcamp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a BootCamp by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BootCamps"
                ],
                "summary": "Delete a BootCamp by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "BootCamp ID",
                        "name": "bootcampId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        },
        "/bootcamps": {
            "get": {
                "description": "Get all BootCamps",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BootCamps"
                ],
                "summary": "Get all BootCamps",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Bootcamp"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "fiber.Map": {
            "type": "object",
            "additionalProperties": true
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "model.Bootcamp": {
            "type": "object",
            "required": [
                "address",
                "careers",
                "description",
                "name",
                "type"
            ],
            "properties": {
                "acceptGi": {
                    "type": "boolean"
                },
                "address": {
                    "type": "string"
                },
                "averageCost": {
                    "type": "number"
                },
                "averageRating": {
                    "type": "number"
                },
                "careers": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "housing": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "jobAssistance": {
                    "type": "boolean"
                },
                "jobGuarantee": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                },
                "userID": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                },
                "zipcode": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/model.UserRole"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.UserRole": {
            "type": "string",
            "enum": [
                "admin",
                "user",
                "guest"
            ],
            "x-enum-varnames": [
                "AdminRole",
                "UsersRole",
                "GuestRole"
            ]
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "DevCamp API",
	Description:      "This is a sample server for DevCamp.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
