// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/consumerinformationmodule/api/consumer": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "consumerinformationmodule"
                ],
                "summary": "ViewAllConsumers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller_echo.ListConsumerResponseViewAllConsumers"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "consumerinformationmodule"
                ],
                "summary": "EditConsumer",
                "parameters": [
                    {
                        "description": "model.Consumer body data",
                        "name": "body-param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Consumer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Consumer"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "consumerinformationmodule"
                ],
                "summary": "CreateConsumer",
                "parameters": [
                    {
                        "description": "model.Consumer body data",
                        "name": "body-param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Consumer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Consumer"
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
                    "consumerinformationmodule"
                ],
                "summary": "DeleteConsumers",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/consumerinformationmodule/api/consumer/profile": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "consumerinformationmodule"
                ],
                "summary": "FilterCompanyById",
                "parameters": [
                    {
                        "type": "string",
                        "description": " string consumerid query ",
                        "name": "consumerid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller_echo.ListConsumerResponseFilterCompanyById"
                        }
                    }
                }
            }
        },
        "/consumerinformationmodule/api/consumerbyid": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "consumerinformationmodule"
                ],
                "summary": "FindConsumer",
                "parameters": [
                    {
                        "type": "string",
                        "description": " string consumerid query ",
                        "name": "consumerid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller_echo.ListConsumerResponseFindConsumer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller_echo.ListConsumerResponseFilterCompanyById": {
            "type": "object",
            "properties": {
                "consumer": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Consumer"
                    }
                }
            }
        },
        "controller_echo.ListConsumerResponseFindConsumer": {
            "type": "object",
            "properties": {
                "consumer": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Consumer"
                    }
                }
            }
        },
        "controller_echo.ListConsumerResponseViewAllConsumers": {
            "type": "object",
            "properties": {
                "consumer": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Consumer"
                    }
                }
            }
        },
        "model.Consumer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "clientid": {
                    "type": "string"
                },
                "clientname": {
                    "type": "string"
                },
                "clientrepresentative": {
                    "type": "string"
                },
                "consumerid": {
                    "type": "string"
                },
                "consumername": {
                    "type": "string"
                },
                "cspid": {
                    "type": "string"
                },
                "dateofbirth": {
                    "type": "string"
                },
                "fedexid": {
                    "type": "string"
                },
                "objectuuid": {
                    "type": "string"
                },
                "othernames": {
                    "type": "string"
                },
                "ssn": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
