// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "HiRaise",
            "url": "https://hiraise.net/",
            "email": "musaev.ae@hiraise.net"
        },
        "license": {
            "name": "MIT License",
            "url": "https://mit-license.org/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/v1/auth"
                ],
                "summary": "login user",
                "parameters": [
                    {
                        "description": "user email and password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "invalid request body",
                        "schema": {
                            "$ref": "#/definitions/customerrors.Err"
                        }
                    },
                    "401": {
                        "description": "invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/customerrors.Err"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "$ref": "#/definitions/customerrors.Err"
                        }
                    }
                }
            }
        },
        "/v1/auth/logout": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/v1/auth"
                ],
                "summary": "logout user",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "authentication required",
                        "schema": {
                            "$ref": "#/definitions/customerrors.Err"
                        }
                    }
                }
            }
        },
        "/v1/auth/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/v1/auth"
                ],
                "summary": "refresh tokens pair",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "refresh token is invalid",
                        "schema": {
                            "$ref": "#/definitions/customerrors.Err"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "$ref": "#/definitions/customerrors.Err"
                        }
                    }
                }
            }
        },
        "/v1/auth/register": {
            "post": {
                "description": "endpoint for register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/v1/auth"
                ],
                "summary": "register new user",
                "parameters": [
                    {
                        "description": "user email and password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "invalid request body",
                        "schema": {
                            "$ref": "#/definitions/customerrors.Err"
                        }
                    },
                    "409": {
                        "description": "user already exists",
                        "schema": {
                            "$ref": "#/definitions/customerrors.Err"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "$ref": "#/definitions/customerrors.Err"
                        }
                    }
                }
            }
        },
        "/v1/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "...",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/v1/users"
                ],
                "summary": "return user by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "customerrors.Err": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {}
                },
                "msg": {
                    "type": "string"
                },
                "responseData": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "source": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "request.User": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "at",
            "in": "cookie"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Task Trail API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
