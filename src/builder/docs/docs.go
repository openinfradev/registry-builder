// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-02-17 15:04:39.168642198 +0900 KST m=+0.055646082

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
        "/docker/build/minio": {
            "post": {
                "description": "docker build by minio api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "docker build by minio",
                "parameters": [
                    {
                        "description": "Json Parameters ",
                        "name": "contents",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.DockerBuildByMinioParam"
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
        "/docker/build/minio-copy-as": {
            "post": {
                "description": "docker build by minio copy as api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "docker build by minio copy as",
                "parameters": [
                    {
                        "description": "Json Parameters ",
                        "name": "contents",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.DockerBuildByMinioCopyAsParam"
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
        "/listener": {
            "post": {
                "description": "registry hook listen",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "registry hook listen",
                "responses": {
                    "200": {}
                }
            }
        },
        "/minio": {
            "post": {
                "description": "create minio",
                "produces": [
                    "application/json"
                ],
                "summary": "create minio",
                "parameters": [
                    {
                        "description": "Json Parameters",
                        "name": "contents",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.MinioParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MinioResult"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete minio",
                "produces": [
                    "application/json"
                ],
                "summary": "delete minio",
                "parameters": [
                    {
                        "description": "Json Parameters",
                        "name": "contents",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.MinioParam"
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
        "/registry/catalog": {
            "get": {
                "description": "docker registry catalog api",
                "consumes": [
                    "application/json"
                ],
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
        "/registry/manifest-v1/{name}": {
            "get": {
                "description": "docker registry manifest v1 api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "docker registry manifest v1 api",
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
                            "$ref": "#/definitions/model.RegistryManifestV1"
                        }
                    }
                }
            }
        },
        "/registry/manifest-v2/{name}": {
            "get": {
                "description": "docker registry manifest v2 api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "docker registry manifest v2 api",
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
                            "$ref": "#/definitions/model.RegistryManifestV2"
                        }
                    }
                }
            }
        },
        "/registry/repositories/{name}": {
            "get": {
                "description": "docker registry repositories api",
                "consumes": [
                    "application/json"
                ],
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
                "consumes": [
                    "application/json"
                ],
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
                        "in": "query"
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
        "/sescan/layer/{name}": {
            "get": {
                "description": "security scanning layer api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "security scanning layer api",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "description": "Layer(history) ID",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SecurityScanLayer"
                        }
                    }
                }
            }
        },
        "/sescan/repository/{name}": {
            "get": {
                "description": "security scanning layer api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "security scanning layer api",
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
                            "$ref": "#/definitions/model.SecurityScanLayer"
                        }
                    }
                }
            },
            "patch": {
                "description": "security scan api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "security scan api",
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
        },
        "/test": {
            "get": {
                "description": "test",
                "produces": [
                    "application/json"
                ],
                "summary": "test api",
                "responses": {
                    "200": {}
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
                "build",
                "contents",
                "name",
                "push",
                "tag",
                "useCache"
            ],
            "properties": {
                "build": {
                    "type": "string"
                },
                "contents": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "push": {
                    "type": "boolean"
                },
                "tag": {
                    "type": "string"
                },
                "useCache": {
                    "type": "boolean"
                }
            }
        },
        "model.DockerBuildByGitParam": {
            "type": "object",
            "required": [
                "build",
                "gitRepo",
                "name",
                "push",
                "tag",
                "useCache",
                "userId",
                "userPw"
            ],
            "properties": {
                "build": {
                    "type": "string"
                },
                "gitRepo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "push": {
                    "type": "boolean"
                },
                "tag": {
                    "type": "string"
                },
                "useCache": {
                    "type": "boolean"
                },
                "userId": {
                    "type": "string"
                },
                "userPw": {
                    "type": "string"
                }
            }
        },
        "model.DockerBuildByMinioCopyAsParam": {
            "type": "object",
            "required": [
                "build",
                "name",
                "path",
                "push",
                "srcPath",
                "srcUserId",
                "tag",
                "useCache",
                "userId"
            ],
            "properties": {
                "build": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "push": {
                    "type": "boolean"
                },
                "srcPath": {
                    "type": "string"
                },
                "srcUserId": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "useCache": {
                    "type": "boolean"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "model.DockerBuildByMinioParam": {
            "type": "object",
            "required": [
                "build",
                "name",
                "path",
                "push",
                "tag",
                "useCache",
                "userId"
            ],
            "properties": {
                "build": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "push": {
                    "type": "boolean"
                },
                "tag": {
                    "type": "string"
                },
                "useCache": {
                    "type": "boolean"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "model.DockerPushParam": {
            "type": "object",
            "required": [
                "build",
                "name",
                "tag"
            ],
            "properties": {
                "build": {
                    "type": "string"
                },
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
                "build",
                "name",
                "newTag",
                "oldTag"
            ],
            "properties": {
                "build": {
                    "type": "string"
                },
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
        "model.MinioParam": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "string"
                },
                "userPw": {
                    "type": "string"
                }
            }
        },
        "model.MinioResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "domain": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "model.RegistryManifestV1": {
            "type": "object",
            "properties": {
                "architecture": {
                    "type": "string"
                },
                "fsLayers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RegistryManifestV1Layer"
                    }
                },
                "history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RegistryManifestV1History"
                    }
                },
                "name": {
                    "type": "string"
                },
                "schemaVersion": {
                    "type": "integer"
                },
                "signatures": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RegistryManifestV1Signature"
                    }
                },
                "tag": {
                    "type": "string"
                }
            }
        },
        "model.RegistryManifestV1History": {
            "type": "object",
            "properties": {
                "v1Compatibility": {
                    "type": "string"
                }
            }
        },
        "model.RegistryManifestV1Layer": {
            "type": "object",
            "properties": {
                "blobSum": {
                    "type": "string"
                }
            }
        },
        "model.RegistryManifestV1Signature": {
            "type": "object",
            "properties": {
                "protected": {
                    "type": "string"
                },
                "signature": {
                    "description": "Header    string ` + "`" + `json:\"header\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "model.RegistryManifestV2": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "object",
                    "$ref": "#/definitions/model.RegistryManifestV2Item"
                },
                "layers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RegistryManifestV2Item"
                    }
                },
                "mediaType": {
                    "type": "string"
                },
                "schemaVersion": {
                    "type": "integer"
                }
            }
        },
        "model.RegistryManifestV2Item": {
            "type": "object",
            "properties": {
                "digest": {
                    "type": "string"
                },
                "mediaType": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
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
                        "$ref": "#/definitions/model.TagResult"
                    }
                }
            }
        },
        "model.SecurityScanLayer": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "object"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.TagResult": {
            "type": "object",
            "properties": {
                "digest": {
                    "type": "string"
                },
                "name": {
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
