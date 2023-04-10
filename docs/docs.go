// Code generated by swaggo/swag. DO NOT EDIT
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
        "/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/files": {
            "get": {
                "description": "retrieve a file",
                "tags": [
                    "Files"
                ],
                "summary": "Files",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Blob"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "description": "retrieve a file",
                "tags": [
                    "Files"
                ],
                "summary": "Files",
                "responses": {
                    "200": {
                        "description": "video info",
                        "schema": {
                            "type": "Videos"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "post": {
                "description": "retrieve a file",
                "tags": [
                    "Files"
                ],
                "summary": "Files",
                "responses": {
                    "200": {
                        "description": "video info",
                        "schema": {
                            "type": "Videos"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "description": "retrieve a file",
                "tags": [
                    "Files"
                ],
                "summary": "Files",
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "patch": {
                "description": "retrieve a file",
                "tags": [
                    "Files"
                ],
                "summary": "Files",
                "responses": {
                    "200": {
                        "description": "video info",
                        "schema": {
                            "type": "Videos"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/files/:id": {
            "get": {
                "description": "retrieve a file",
                "tags": [
                    "Files"
                ],
                "summary": "Files",
                "responses": {
                    "200": {
                        "description": "video info",
                        "schema": {
                            "type": "Videos"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/files/detail": {
            "get": {
                "description": "retrieve a file",
                "tags": [
                    "Files"
                ],
                "summary": "Files",
                "responses": {
                    "200": {
                        "description": "video info",
                        "schema": {
                            "type": "Videos"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/files/files": {
            "get": {
                "description": "retrieve a file",
                "tags": [
                    "Files"
                ],
                "summary": "Files",
                "responses": {
                    "200": {
                        "description": "video info",
                        "schema": {
                            "type": "Videos"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/videos/:videoID": {
            "get": {
                "description": "get a video by id",
                "tags": [
                    "Videos"
                ],
                "summary": "Videos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Videos"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/videos/chann/:channId": {
            "get": {
                "description": "get all videos from a channel",
                "tags": [
                    "Videos"
                ],
                "summary": "Videos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Videos"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
