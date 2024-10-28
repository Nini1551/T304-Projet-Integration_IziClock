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
        "/alarms/state/{id}": {
            "put": {
                "responses": {}
            }
        },
        "/calendar": {
            "get": {
                "description": "Récupère une liste des événements à venir depuis le calendrier Google de l'utilisateur.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calendrier"
                ],
                "summary": "Récupère les événements du calendrier",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Code d'autorisation pour l'API Google",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Liste des événements",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Requête incorrecte",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Erreur interne du serveur",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/calendars": {
            "get": {
                "description": "Récupère une liste de toutes les alarmes depuis la DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calendriers"
                ],
                "summary": "Récupère tous les calendriers",
                "responses": {
                    "200": {
                        "description": "Calendars send successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Calendar"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/calendars/state/{id}": {
            "put": {
                "description": "Change l'attribut d'activité d'un calendrier via l'identifiant de celui-ci",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calendriers"
                ],
                "summary": "Change l'attribut d'activité d'un calendrier",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Calendar ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Calendar updated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Calendar"
                        }
                    },
                    "404": {
                        "description": "Calendar not found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/calendars/{id}": {
            "delete": {
                "description": "Efface un calendrier via son ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calendriers"
                ],
                "summary": "Efface un calendrier",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Calendar ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Calendar deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Calendar not found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/loginCalendar": {
            "get": {
                "description": "Fournit un lien permettant à l'utilisateur de se connecter et d'autoriser l'accès en lecture seule à son calendrier Google.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Authentification"
                ],
                "summary": "Génère un lien d'autorisation Google",
                "responses": {
                    "200": {
                        "description": "URL d'autorisation Google pour l'utilisateur",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erreur interne du serveur",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Renvoie un message 'pong' pour vérifier que le serveur est en ligne",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ping"
                ],
                "summary": "Envoie un ping",
                "responses": {
                    "200": {
                        "description": "Ping Pong"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Calendar": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "IziClock API Documentation",
	Description:      "Il s'agit de la documentation de l'API IziClock.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,

}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
