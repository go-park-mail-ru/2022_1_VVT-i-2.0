// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comment": {
            "post": {
                "description": "Add comments by restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Add comments by restaurant",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CommentDataDelivery"
                        }
                    }
                }
            }
        },
        "/comments/:id": {
            "get": {
                "description": "get comments by restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "List comments by restaurant",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CommentsDataDelivery"
                        }
                    }
                }
            }
        },
        "/restaurant/:slug": {
            "get": {
                "description": "get dishes by restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants"
                ],
                "summary": "List dishes by restaurant",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RestaurantsDishesJsonForKirill"
                        }
                    }
                }
            }
        },
        "/restaurants": {
            "get": {
                "description": "Get restaurants",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants"
                ],
                "summary": "List restaurants",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.RestaurantJsonForKirill"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CommentDataDelivery": {
            "type": "object",
            "properties": {
                "commentRating": {
                    "type": "integer"
                },
                "commentText": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "restaurants": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.CommentsDataDelivery": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CommentDataDelivery"
                    }
                }
            }
        },
        "models.DishJsonForKirill": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imgPath": {
                    "type": "string"
                },
                "info": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "productName": {
                    "type": "string"
                },
                "restaurany": {
                    "type": "integer"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "models.RestaurantJsonForKirill": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imgPath": {
                    "type": "string"
                },
                "min_price": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "restName": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "timeToDeliver": {
                    "type": "string"
                }
            }
        },
        "models.RestaurantsDishesJsonForKirill": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "avgPrice": {
                    "type": "integer"
                },
                "city": {
                    "type": "string"
                },
                "dishes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DishJsonForKirill"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "imgPath": {
                    "type": "string"
                },
                "minPrice": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "restName": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "timeToDeliver": {
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
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
