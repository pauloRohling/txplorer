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
            "name": "API Support",
            "url": "https://github.com/pauloRohling/txplorer"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/pauloRohling/txplorer/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "post": {
                "description": "Creates a new account and a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Create Account",
                "parameters": [
                    {
                        "description": "Account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateAccountInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/account.CreateAccountOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/operations/deposit": {
            "post": {
                "description": "Deposits funds to an account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Operation"
                ],
                "summary": "Deposit",
                "parameters": [
                    {
                        "description": "Account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.DepositInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/operation.DepositOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/operations/transfer": {
            "post": {
                "description": "Transfers funds from one account to another",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Operation"
                ],
                "summary": "Transfer",
                "parameters": [
                    {
                        "description": "Account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.TransferInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/operation.TransferOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/operations/withdraw": {
            "post": {
                "description": "Withdraws funds from an account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Operation"
                ],
                "summary": "Withdraw",
                "parameters": [
                    {
                        "description": "Account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.WithdrawInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/operation.WithdrawOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Generates a JWT token to authenticate the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Email",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.LoginOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "account.CreateAccountOutput": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/model.AccountStatus"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "model.AccountStatus": {
            "type": "string",
            "enum": [
                "ACTIVE",
                "INACTIVE"
            ],
            "x-enum-varnames": [
                "AccountStatusActive",
                "AccountStatusInactive"
            ]
        },
        "model.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/model.ErrorType"
                }
            }
        },
        "model.ErrorType": {
            "type": "string",
            "enum": [
                "ForbiddenError",
                "InternalError",
                "NotFoundError",
                "UnauthorizedError",
                "ValidationError"
            ],
            "x-enum-varnames": [
                "ForbiddenErrorType",
                "InternalErrorType",
                "NotFoundErrorType",
                "UnauthorizedErrorType",
                "ValidationErrorType"
            ]
        },
        "model.OperationStatus": {
            "type": "string",
            "enum": [
                "PENDING",
                "SUCCESS",
                "FAILED"
            ],
            "x-enum-varnames": [
                "OperationStatusPending",
                "OperationStatusSuccess",
                "OperationStatusFailed"
            ]
        },
        "operation.DepositOutput": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "fromAccountId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/model.OperationStatus"
                },
                "toAccountId": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "operation.TransferOutput": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "fromAccountId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/model.OperationStatus"
                },
                "toAccountId": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "operation.WithdrawOutput": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "fromAccountId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/model.OperationStatus"
                },
                "toAccountId": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "types.CreateAccountInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "types.DepositInput": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "string"
                },
                "amount": {
                    "type": "integer"
                }
            }
        },
        "types.LoginInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "types.TransferInput": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "fromAccountId": {
                    "type": "string"
                },
                "toAccountId": {
                    "type": "string"
                }
            }
        },
        "types.WithdrawInput": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "string"
                },
                "amount": {
                    "type": "integer"
                }
            }
        },
        "user.LoginOutput": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "TxPlorer API",
	Description:      "This is a transactional application that allows users to transfer funds between their accounts.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
