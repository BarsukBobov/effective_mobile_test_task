// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cars/": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "create",
                "parameters": [
                    {
                        "description": "create",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cars.createForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/baseApi.SuccessResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/baseApi.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/cars/get_all": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "getAll",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "mark",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "model",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "owner_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "owner_patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "owner_surname",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "reg_num",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "year",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cars.getAllResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/baseApi.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/cars/{id}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "edit",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "edit",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cars.editForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/baseApi.SuccessResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/baseApi.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/baseApi.SuccessResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/baseApi.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "baseApi.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "baseApi.SuccessResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "cars.createForm": {
            "type": "object",
            "required": [
                "regNums"
            ],
            "properties": {
                "regNums": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "cars.editForm": {
            "type": "object",
            "properties": {
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner_name": {
                    "type": "string"
                },
                "owner_patronymic": {
                    "type": "string"
                },
                "owner_surname": {
                    "type": "string"
                },
                "reg_num": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "cars.getAllResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sql.Car"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "sql.Car": {
            "type": "object",
            "properties": {
                "RegNum": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/sql.People"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "sql.People": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "EFFECTIVE MOBILE TEST TASK API",
	Description:      "API Server for effective mobile's test task",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
