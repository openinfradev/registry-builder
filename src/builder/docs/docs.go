// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-08-13 17:46:09.77778 +0900 KST m=+0.027545183

package docs

import (
	"bytes"
	"encoding/json"

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/docker/build/file": {
            "post": {
                "description": "docker build by dockerfile api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "docker build by dockerfile",
                "parameters": [
                    {
                        "description": "Json Parameters (contents is base64 encoded)",
                        "name": "contents",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.DockerBuildByFileParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResult"
                        }
                    }
                }
            }
        },
        "/docker/build/git": {
            "post": {
                "description": "docker build by git api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "docker build by git",
                "parameters": [
                    {
                        "description": "Json Parameters (userPW is base64 encoded)",
                        "name": "contents",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.DockerBuildByGitParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResult"
                        }
                    }
                }
            }
        },
        "/docker/push": {
            "put": {
                "description": "docker image push",
                "produces": [
                    "application/json"
                ],
                "summary": "docker image push",
                "parameters": [
                    {
                        "description": "Json Parameters",
                        "name": "contents",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.DockerPushParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResult"
                        }
                    }
                }
            }
        },
        "/docker/tag": {
            "patch": {
                "description": "docker image tag",
                "produces": [
                    "application/json"
                ],
                "summary": "docker image tag",
                "parameters": [
                    {
                        "description": "Json Parameters",
                        "name": "contents",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.DockerTagParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResult"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "builder의 health를 체크할 목적의 api",
                "produces": [
                    "application/json"
                ],
                "summary": "health check api",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResult"
                        }
                    }
                }
            }
        },
        "/registry/catalog": {
            "get": {
                "description": "docker registry catalog api",
                "produces": [
                    "application/json"
                ],
                "summary": "docker registry catalog api",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CatalogResult"
                        }
                    }
                }
            }
        },
        "/registry/repositories/{name}": {
            "get": {
                "description": "docker registry repositories api",
                "produces": [
                    "application/json"
                ],
                "summary": "docker registry repositories api",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "description": "Repository Name",
                        "name": "name",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.RepositoryResult"
                        }
                    }
                }
            },
            "delete": {
                "description": "docker registry repository delete api",
                "produces": [
                    "application/json"
                ],
                "summary": "docker registry repository delete api",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "description": "Repository Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Tag Name",
                        "name": "tag",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.BasicResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.CatalogResult": {
            "type": "object",
            "properties": {
                "repositories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.DockerBuildByFileParam": {
            "type": "object",
            "required": [
                "contents",
                "name"
            ],
            "properties": {
                "contents": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.DockerBuildByGitParam": {
            "type": "object",
            "required": [
                "gitRepo",
                "name",
                "userId",
                "userPw"
            ],
            "properties": {
                "gitRepo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                },
                "userPw": {
                    "type": "string"
                }
            }
        },
        "model.DockerPushParam": {
            "type": "object",
            "required": [
                "name",
                "tag"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                }
            }
        },
        "model.DockerTagParam": {
            "type": "object",
            "required": [
                "name",
                "newTag",
                "oldTag"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "newTag": {
                    "type": "string"
                },
                "oldTag": {
                    "type": "string"
                }
            }
        },
        "model.RepositoriesResult": {
            "type": "object",
            "properties": {
                "repositories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RepositoryResult"
                    }
                }
            }
        },
        "model.RepositoryResult": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
	Version:     "{{.Version}}",
	Host:        "{{.Host}}",
	BasePath:    "{{.BasePath}}",
	Schemes:     []string{},
	Title:       "{{.Title}}",
	Description: "{{.Description}}",
}

type s struct{}

func (s *s) ReadDoc() string {
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
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
