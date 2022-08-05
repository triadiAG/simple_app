// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
    "github.com/swaggo/swag"
)


// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "SimpleAPIs",
	Description:      "BooksAPIs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}
func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

const docTemplate = `
{
  "swagger": "2.0",
  "info": {
    "version": "1.0",
    "title": "simpleapis",
    "contact": {}
  },
  "host": "localhost:8080",
  "basePath": "/",
  "securityDefinitions": {},
  "schemes": [
    "http"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/books": {
      "get": {
        "summary": "get books",
        "tags": [
          "buku"
        ],
        "operationId": "getbooks",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      },
      "post": {
        "summary": "post books",
        "tags": [
          "buku"
        ],
        "operationId": "postbooks",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "parameters": [
          {
            "name": "judul",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "pengarang",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "penerbit",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "tahun_terbit",
            "in": "formData",
            "required": true,
            "type": "integer",
            "format": "int32",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/books/1": {
      "get": {
        "summary": "get one book",
        "tags": [
          "buku"
        ],
        "operationId": "getonebook",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      },
      "put": {
        "summary": "udpate book",
        "tags": [
          "buku"
        ],
        "operationId": "udpatebook",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "parameters": [
          {
            "name": "judul",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "pengarang",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "penerbit",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "tahun_terbit",
            "in": "formData",
            "required": true,
            "type": "integer",
            "format": "int32",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/books/2": {
      "delete": {
        "summary": "delete",
        "tags": [
          "buku"
        ],
        "operationId": "delete",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/users": {
      "get": {
        "summary": "Get Users",
        "tags": [
          "user"
        ],
        "operationId": "GetUsers",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      },
      "post": {
        "summary": "post users",
        "tags": [
          "user"
        ],
        "operationId": "postusers",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "parameters": [
          {
            "name": "username",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "email",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "password",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/users/1": {
      "get": {
        "summary": "Get One User",
        "tags": [
          "user"
        ],
        "operationId": "GetOneUser",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/users/3": {
      "put": {
        "summary": "update user",
        "tags": [
          "user"
        ],
        "operationId": "updateuser",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "parameters": [
          {
            "name": "username",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "email",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "password",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/users/2": {
      "delete": {
        "summary": "delete user",
        "tags": [
          "user"
        ],
        "operationId": "deleteuser",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/auth/login": {
      "post": {
        "summary": "login",
        "tags": [
          "auth"
        ],
        "operationId": "login",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "parameters": [
          {
            "name": "email",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "password",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    }
  },
  "tags": [
    {
      "name": "buku"
    },
    {
      "name": "user"
    },
    {
      "name": "auth"
    }
  ]
}
`