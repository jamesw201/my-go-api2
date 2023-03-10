// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is my new go api",
    "title": "new-go-api",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "sites"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list sites",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/site"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "sites"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/site"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/site"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "put": {
        "tags": [
          "sites"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/site"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/site"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "sites"
        ],
        "operationId": "destroyOne",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "domain": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "environment": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "site": {
      "type": "object",
      "required": [
        "domains",
        "environments"
      ],
      "properties": {
        "description": {
          "type": "string",
          "minLength": 1
        },
        "domains": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/domain"
          }
        },
        "environments": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/environment"
          }
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is my new go api",
    "title": "new-go-api",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "sites"
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list sites",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/site"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "sites"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/site"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/site"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "put": {
        "tags": [
          "sites"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/site"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/site"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "sites"
        ],
        "operationId": "destroyOne",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "domain": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "environment": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "site": {
      "type": "object",
      "required": [
        "domains",
        "environments"
      ],
      "properties": {
        "description": {
          "type": "string",
          "minLength": 1
        },
        "domains": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/domain"
          }
        },
        "environments": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/environment"
          }
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  }
}`))
}
