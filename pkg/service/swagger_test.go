package service

import (
	"github.com/bbdshow/admin-rabc/pkg/model"
	"testing"
)

var swaggerTxt = `{
    "swagger": "2.0",
    "info": {
        "description": "admin rbac manage API",
        "title": "admin rbac",
        "contact": {},
        "version": "1.0.0"
    },
    "host": "API_HOST:49000",
    "basePath": "/",
    "paths": {
        "/rbac/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 登录/登出"
                ],
                "summary": "[RBAC 登录]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.LoginAccountResp"
                        }
                    }
                }
            }
        },
        "/rbac/loginOut": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 登录/登出"
                ],
                "summary": "[RBAC 登出]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginOutAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户配置创建]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户配置列表]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.ListAccount"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/pwd/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户配置密码更改]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateAccountPasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/role/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户配置角色更改]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateAccountRoleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/action/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 功能配置"
                ],
                "summary": "[功能配置列表]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListActionConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.ListActionConfig"
                        }
                    }
                }
            }
        },
        "/rbac/v1/action/upsert": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 功能配置"
                ],
                "summary": "[功能配置创建/更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpsertActionConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/app/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC APP配置"
                ],
                "summary": "[APP配置创建]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateAppConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/app/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC APP配置"
                ],
                "summary": "[APP配置列表]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListAppConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.ListAppConfig"
                        }
                    }
                }
            }
        },
        "/rbac/v1/app/select": {
            "get": {
                "description": "用于Select组件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC APP配置"
                ],
                "summary": "[APP配置筛选列表]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SelectAppConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.SelectAppConfig"
                        }
                    }
                }
            }
        },
        "/rbac/v1/app/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC APP配置"
                ],
                "summary": "[APP配置更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateAppConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/menu/action/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 菜单配置"
                ],
                "summary": "[菜单配置功能更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateMenuConfigActionReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/menu/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 菜单配置"
                ],
                "summary": "[菜单配置创建]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateMenuConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/menu/tree": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 菜单配置"
                ],
                "summary": "[菜单配置树]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetMenuTreeDirsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.GetMenuTreeDirsResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/menu/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 菜单配置"
                ],
                "summary": "[菜单配置更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateMenuConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/role/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 角色配置"
                ],
                "summary": "[角色配置创建]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateRoleConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/role/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 角色配置"
                ],
                "summary": "[角色配置列表]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListRoleConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.ListRoleConfig"
                        }
                    }
                }
            }
        },
        "/rbac/v1/role/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 角色配置"
                ],
                "summary": "[角色配置更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateRoleConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ginutil.BaseResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "traceId": {
                    "type": "string"
                }
            }
        },
        "model.CreateAccountReq": {
            "type": "object",
            "required": [
                "appId",
                "nickname",
                "pwd",
                "username"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "pwd": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.CreateAppConfigReq": {
            "type": "object",
            "required": [
                "memo",
                "name"
            ],
            "properties": {
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.CreateMenuConfigReq": {
            "type": "object",
            "required": [
                "appId",
                "name",
                "parentId",
                "path",
                "sequence"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                }
            }
        },
        "model.CreateRoleConfigReq": {
            "type": "object",
            "required": [
                "appId",
                "name"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "isRoot": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.GetMenuTreeDirsReq": {
            "type": "object",
            "required": [
                "appId"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                }
            }
        },
        "model.GetMenuTreeDirsResp": {
            "type": "object",
            "properties": {
                "dirs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MenuTreeDir"
                    }
                }
            }
        },
        "model.ListAccount": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "appName": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "integer"
                },
                "loginLock": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "pwdWrong": {
                    "type": "integer"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "integer"
                },
                "tokenExpired": {
                    "type": "integer"
                }
            }
        },
        "model.ListAccountReq": {
            "type": "object",
            "required": [
                "page",
                "size"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.ListActionConfig": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "integer"
                }
            }
        },
        "model.ListActionConfigReq": {
            "type": "object",
            "required": [
                "id",
                "page",
                "size"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "model.ListAppConfig": {
            "type": "object",
            "properties": {
                "accessKey": {
                    "description": "访问KEY",
                    "type": "string"
                },
                "appId": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "memo": {
                    "description": "备注",
                    "type": "string"
                },
                "name": {
                    "description": "APP名",
                    "type": "string"
                },
                "secretKey": {
                    "description": "加密KEY",
                    "type": "string"
                },
                "status": {
                    "description": "状态 1-正常 2-限制",
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "integer"
                }
            }
        },
        "model.ListAppConfigReq": {
            "type": "object",
            "required": [
                "page",
                "size"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.ListRoleConfig": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "appName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isRoot": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "integer"
                }
            }
        },
        "model.ListRoleConfigReq": {
            "type": "object",
            "required": [
                "page",
                "size"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "model.LoginAccountReq": {
            "type": "object",
            "required": [
                "appId",
                "password",
                "username"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.LoginAccountResp": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "tokenExpired": {
                    "type": "integer"
                }
            }
        },
        "model.LoginOutAccountReq": {
            "type": "object"
        },
        "model.MenuTreeDir": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "children": {
                    "$ref": "#/definitions/model.MenuTreeDirs"
                },
                "id": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.MenuTreeDirs": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/model.MenuTreeDir"
            }
        },
        "model.SelectAppConfig": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "memo": {
                    "description": "备注",
                    "type": "string"
                },
                "name": {
                    "description": "APP名",
                    "type": "string"
                }
            }
        },
        "model.SelectAppConfigReq": {
            "type": "object",
            "required": [
                "page",
                "size"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateAccountPasswordReq": {
            "type": "object",
            "required": [
                "id",
                "pwd"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "pwd": {
                    "type": "string"
                }
            }
        },
        "model.UpdateAccountRoleReq": {
            "type": "object",
            "required": [
                "id",
                "roles"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.UpdateAppConfigReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "secretKey": {
                    "description": "true = 重置加密KEY",
                    "type": "boolean"
                },
                "status": {
                    "description": "状态 1-正常 2-限制",
                    "type": "integer"
                }
            }
        },
        "model.UpdateMenuConfigActionReq": {
            "type": "object",
            "required": [
                "menuId"
            ],
            "properties": {
                "actionId": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "menuId": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateMenuConfigReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateRoleConfigReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isRoot": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.UpsertActionConfigReq": {
            "type": "object",
            "required": [
                "appId",
                "method",
                "name",
                "path"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        }
    }
}`

func TestService_SwaggerJSONToActions(t *testing.T) {
	in := &model.SwaggerJSONToActionsReq{
		AppId:      "000000",
		SwaggerTxt: swaggerTxt,
	}
	if err := svc.SwaggerJSONToActions(ctx, in); err != nil {
		t.Fatal(err)
	}
}
