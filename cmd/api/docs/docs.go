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
        "/customer": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "get customer by document",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Document",
                        "name": "document",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Customer",
                        "schema": {
                            "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Customer"
                        }
                    },
                    "400": {
                        "description": "invalid document",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "customer not find",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "save customer in DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "create customer",
                "parameters": [
                    {
                        "description": "Customer",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Customer"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "invalid document, invalid email...",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "description": "Get Order from DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "get all order",
                "responses": {
                    "200": {
                        "description": "Order",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Order"
                            }
                        }
                    },
                    "400": {
                        "description": "invalid document",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "customer not find",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "save Order in DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "create order",
                "parameters": [
                    {
                        "description": "Order",
                        "name": "Order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Order"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders/sorted": {
            "get": {
                "description": "Get Order from DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "get all order",
                "responses": {
                    "200": {
                        "description": "Order",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Order"
                            }
                        }
                    },
                    "400": {
                        "description": "invalid document",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "customer not find",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/payment": {
            "post": {
                "description": "save Payment in DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payment"
                ],
                "summary": "create payment",
                "parameters": [
                    {
                        "description": "Payment",
                        "name": "Payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Payment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/payment/webhook": {
            "post": {
                "description": "update payment and set paid in DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payment"
                ],
                "summary": "create Payment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "get product by category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "invalid document",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "customer not find",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "update product in DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "update product",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "save product in DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "create product",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete product in DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "delete product",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "general error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Customer": {
            "type": "object",
            "properties": {
                "created_date": {
                    "type": "string"
                },
                "document": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_modified_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Order": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_date_db": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_modified_date_db": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.OrderProduct"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.OrderProduct": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Payment": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "notification_url": {
                    "type": "string"
                },
                "total_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "github_com_Pos-Tech-Challenge-48_delivery-api_internal_entities.Product": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "created_date_db": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "last_modified_date_db": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1/delivery",
	Schemes:          []string{},
	Title:            "Delivery API",
	Description:      "Aplicativo que gerencia atividades de um serviço de pedidos em um restaurante. Desde a base de clientes, catálogo de produtos, pedidos e fila de preparo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
