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
        "/alarms/{id}": {
            "put": {
                "description": "Met à jour les détails d'une alarme par ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Alarmes"
                ],
                "summary": "Met à jour les détails de l'alarme",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de l'alarme",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Détails de l'alarme",
                        "name": "alarm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Alarm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Alarm"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
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
        },
        "/raspberry": {
            "get": {
                "description": "Retourne une liste des alarmes mises à jour ou proches de sonner.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Raspberry"
                ],
                "summary": "Récupérer les alarmes mises à jour pour le Raspberry",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/raspberry.AlarmResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ringtones": {
            "get": {
                "description": "Récupère une liste de toutes les sonneries depuis la DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sonneries"
                ],
                "summary": "Récupère toutes les sonneries",
                "responses": {
                    "200": {
                        "description": "Ringtones send successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Ringtone"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/ringtones/name/{id}": {
            "put": {
                "description": "Modifie le nom d'une sonnerie en DB à partir de son ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sonneries"
                ],
                "summary": "Modifie le nom d'une sonnerie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Ringtone ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ringtone updated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Ringtone"
                        }
                    },
                    "404": {
                        "description": "Ringtone not found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/ringtones/upload": {
            "post": {
                "description": "Upload une nouvelle sonnerie et sauve son url dans la base de données",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sonneries"
                ],
                "summary": "Upload une sonnerie",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Fichier audio de la sonnerie à upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "File uploaded and transferred successfully"
                    },
                    "400": {
                        "description": "No file is received"
                    },
                    "409": {
                        "description": "Ringtone already exists"
                    },
                    "500": {
                        "description": "Unable to save the file"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Alarm": {
            "type": "object",
            "properties": {
                "calendar": {
                    "$ref": "#/definitions/models.Calendar"
                },
                "calendarID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "lastUpdate": {
                    "type": "string"
                },
                "locationEnd": {
                    "type": "string"
                },
                "locationStart": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ringDate": {
                    "type": "string"
                },
                "ringtone": {
                    "$ref": "#/definitions/models.Ringtone"
                },
                "ringtoneID": {
                    "type": "integer"
                },
                "update": {
                    "type": "string"
                }
            }
        },
        "models.Calendar": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "idgoogle": {
                    "description": "Url       string    ` + "`" + `gorm:\"size:255; not null; uniqueIndex:idx_calendar_url\"` + "`" + `",
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "userID": {
                    "description": "ID string ` + "`" + `gorm:\"primaryKey;size 255;\"` + "`" + ` test abandonné\nUserID uint   ` + "`" + `gorm:\"not null; uniqueIndex:idx_calendar_name; uniqueIndex:idx_calendar_idgoogle\"` + "`" + `",
                    "type": "integer"
                }
            }
        },
        "models.Ringtone": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "raspberry.AlarmResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "ringDate": {
                    "type": "string"
                },
                "ringtone": {
                    "$ref": "#/definitions/raspberry.RingtoneBrief"
                },
                "ringtoneID": {
                    "type": "integer"
                }
            }
        },
        "raspberry.RingtoneBrief": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "url": {
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
