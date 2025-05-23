// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag/v2"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "components": {"schemas":{"models.Message":{"properties":{"message":{"type":"string"}},"type":"object"},"models.Todo":{"properties":{"id":{"type":"string"},"task":{"type":"string"}},"type":"object"}}},
    "info": {"contact":{"email":"support@swagger.io","name":"API Support","url":"http://www.swagger.io/support"},"description":"{{escape .Description}}","license":{"name":"MIT","url":"https://opensource.org/licenses/MIT"},"title":"{{.Title}}","version":"{{.Version}}"},
    "externalDocs": {"description":"","url":""},
    "paths": {"/todo":{"get":{"description":"this is a description of sorts, so to speak, as it were","operationId":"get-all-todos","responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/models.Todo"}}},"description":"OK"}},"summary":"get all items in the todo list"}},"/todo/{id}":{"get":{"description":"this is a pretty verbose description that really dives in depth\nand tells the user\neverything they could possibly need to know and in such an effective manner that\nthe only emails they will send us are to thank us for changing their lives for the\nbetter so profoundly.","operationId":"get-todo-by-id","parameters":[{"description":"todo ID","in":"path","name":"id","required":true,"schema":{"type":"string"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/models.Todo"}}},"description":"OK"},"404":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/models.Message"}}},"description":"Not Found"}},"summary":"get a todo item by ID"}}},
    "openapi": "3.1.0",
    "servers": [
        {"url":"/"}
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Title:            "Go + Gin Todo API",
	Description:      "This is a sample server todo server. You can visit the GitHub repository at https://github.com/kwinwithak/swag-gin-demo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
