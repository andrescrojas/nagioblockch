/*******************************************************************************
Copyright (c) 2017 IBM Corporation and other Contributors.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.


Contributors:

Andres Rojas  - Assist Consultores - acrojas@assist.com.co
******************************************************************************/


package main

var schemas = `{
	"API": {
		"createAsset": {
			"description": "Create an asset. One argument, a JSON encoded event. name is required with zero or more writable properties. Establishes an initial asset state.",
			"properties": {
				"args": {
					"description": "args are JSON encoded strings",
					"items": {
						"description": "A set of fields that constitute the writable fields in an asset's state. name is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
						"properties": {
							"txID": {
								"description": "The ID of a managed asset. The resource focal point for a smart contract.",
								"type": "string"
							},
							"timestamp": {
								"description": "Datetime of the transaction",
								"type": "string"
							},
							"issuing": {
								"description": "Issuing of the transaction",
								"type": "number"
							},
							"receiver": {
								"description": "Receiver of the transaction",
								"type": "string"
							},
							"type": {
								"description": "Type of the transaction",
								"type": "string"
							},
							"acquirer": {
								"description": "Acquirer of the transaction",
								"type": "string"
							},
							"acquirerBank": {
								"description": "AcquirerBank of the transaction",
								"type": "string"
							},
							"issuer": {
								"description": "Issuer of the transaction",
								"type": "string"
							},
							"amount": {
								"description": "Amount involved in the transaction",
								"type": "number"
							},
							"currency": {
								"description": "Currency of the amount specified",
								"type": "number"
							},
							"origin": {
								"description": "Origin of the transaction",
								"type": "string"
							}
						},
						"required": ["txID",
							"timestamp",
							"issuing",
							"receiver",
							"type",
							"acquirer",
							"acquirerBank",
							"issuer",
							"amount",
							"currency",
							"origin"
						],
						"type": "object"
					},
					"maxItems": 1,
					"minItems": 1,
					"type": "array"
				},
				"function": {
					"description": "createAsset function",
					"enum": ["createAsset"],
					"type": "string"
				},
				"method": "invoke"
			},
			"type": "object"
		},
		"deleteAsset": {
			"description": "Delete an asset. Argument is a JSON encoded string containing only an name.",
			"properties": {
				"args": {
					"description": "args are JSON encoded strings",
					"items": {
						"description": "An object containing only an name for use as an argument to read or delete.",
						"properties": {
							"txID": {
								"description": "The ID of a managed asset. The resource focal point for a smart contract.",
								"type": "string"
							}
						},
						"type": "object"
					},
					"maxItems": 1,
					"minItems": 1,
					"type": "array"
				},
				"function": {
					"description": "deleteAsset function",
					"enum": ["deleteAsset"],
					"type": "string"
				},
				"method": "invoke"
			},
			"type": "object"
		},
		"init": {
			"description": "Initializes the contract when started, either by deployment or by peer restart.",
			"properties": {
				"args": {
					"description": "args are JSON encoded strings",
					"items": {
						"description": "event sent to init on deployment",
						"properties": {
							"transaction": {
								"default": "InitChainCode",
								"description": "Init of the current contract",
								"type": "string"
							},
							"version": {
								"description": "The ID of a managed asset. The resource focal point for a smart contract.",
								"type": "string"
							}
						},
						"required": ["version"],
						"type": "object"
					},
					"maxItems": 1,
					"minItems": 1,
					"type": "array"
				},
				"function": {
					"description": "init function",
					"enum": ["init"],
					"type": "string"
				},
				"method": "deploy"
			},
			"type": "object"
		},
		"readAsset": {
			"description": "Returns the state an asset. Argument is a JSON encoded string. name is the only accepted property.",
			"properties": {
				"args": {
					"description": "args are JSON encoded strings",
					"items": {
						"description": "An object containing only an name for use as an argument to read or delete.",
						"properties": {
							"txID": {
								"description": "The ID of a managed asset. The resource focal point for a smart contract.",
								"type": "string"
							}
						},
						"type": "object"
					},
					"maxItems": 1,
					"minItems": 1,
					"type": "array"
				},
				"function": {
					"description": "readAsset function",
					"enum": ["readAsset"],
					"type": "string"
				},
				"method": "query",
				"result": {
					"description": "A set of fields that constitute the complete asset state.",
					"properties": {
						"txID": {
							"description": "The ID of a managed asset. The resource focal point for a smart contract.",
							"type": "string"
						},
						"timestamp": {
							"description": "Datetime of the transaction",
							"type": "string"
						},
						"issuing": {
							"description": "Issuing of the transaction",
							"type": "number"
						},
						"receiver": {
							"description": "Receiver of the transaction",
							"type": "string"
						},
						"type": {
							"description": "Type of the transaction",
							"type": "string"
						},
						"acquirer": {
							"description": "Acquirer of the transaction",
							"type": "string"
						},
						"acquirerBank": {
							"description": "AcquirerBank of the transaction",
							"type": "string"
						},
						"issuer": {
							"description": "Issuer of the transaction",
							"type": "string"
						},
						"amount": {
							"description": "Amount involved in the transaction",
							"type": "number"
						},
						"currency": {
							"description": "Currency of the amount specified",
							"type": "number"
						},
						"origin": {
							"description": "Origin of the transaction",
							"type": "string"
						}
					},
					"type": "object"
				}
			},
			"type": "object"
		},
		"readAssetSamples": {
			"description": "Returns a string generated from the schema containing sample Objects as specified in generate.json in the scripts folder.",
			"properties": {
				"args": {
					"description": "accepts no arguments",
					"items": {

					},
					"maxItems": 0,
					"minItems": 0,
					"type": "array"
				},
				"function": {
					"description": "readAssetSamples function",
					"enum": ["readAssetSamples"],
					"type": "string"
				},
				"method": "query",
				"result": {
					"description": "JSON encoded object containing selected sample data",
					"type": "string"
				}
			},
			"type": "object"
		},
		"readAssetSchemas": {
			"description": "Returns a string generated from the schema containing APIs and Objects as specified in generate.json in the scripts folder.",
			"properties": {
				"args": {
					"description": "accepts no arguments",
					"items": {

					},
					"maxItems": 0,
					"minItems": 0,
					"type": "array"
				},
				"function": {
					"description": "readAssetSchemas function",
					"enum": ["readAssetSchemas"],
					"type": "string"
				},
				"method": "query",
				"result": {
					"description": "JSON encoded object containing selected schemas",
					"type": "string"
				}
			},
			"type": "object"
		},
		"updateAsset": {
			"description": "Update the state of an asset. The one argument is a JSON encoded event. name is required along with one or more writable properties. Establishes the next asset state. ",
			"properties": {
				"args": {
					"description": "args are JSON encoded strings",
					"items": {
						"description": "A set of fields that constitute the writable fields in an asset's state. name is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
						"properties": {
							"txID": {
								"description": "The ID of a managed asset. The resource focal point for a smart contract.",
								"type": "string"
							},
							"timestamp": {
								"description": "Datetime of the transaction",
								"type": "string"
							},
							"issuing": {
								"description": "Issuing of the transaction",
								"type": "number"
							},
							"receiver": {
								"description": "Receiver of the transaction",
								"type": "string"
							},
							"type": {
								"description": "Type of the transaction",
								"type": "string"
							},
							"acquirer": {
								"description": "Acquirer of the transaction",
								"type": "string"
							},
							"acquirerBank": {
								"description": "AcquirerBank of the transaction",
								"type": "string"
							},
							"issuer": {
								"description": "Issuer of the transaction",
								"type": "string"
							},
							"amount": {
								"description": "Amount involved in the transaction",
								"type": "number"
							},
							"currency": {
								"description": "Currency of the amount specified",
								"type": "number"
							},
							"origin": {
								"description": "Origin of the transaction",
								"type": "string"
							}
						},
						"required": ["txID",
							"timestamp",
							"issuing",
							"receiver",
							"type",
							"acquirer",
							"acquirerBank",
							"issuer",
							"amount",
							"currency",
							"origin"
						],
						"type": "object"
					},
					"maxItems": 1,
					"minItems": 1,
					"type": "array"
				},
				"function": {
					"description": "updateAsset function",
					"enum": ["updateAsset"],
					"type": "string"
				},
				"method": "invoke"
			},
			"type": "object"
		}
	},
	"objectModelSchemas": {
		"deviceIDKey": {
			"description": "An object containing only an name for use as an argument to read or delete.",
			"properties": {
				"txID": {
					"description": "The ID of a managed asset. The resource focal point for a smart contract.",
					"type": "string"
				}
			},
			"type": "object"
		},
		"event": {
			"description": "A set of fields that constitute the writable fields in an asset's state. name is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
			"properties": {
				"txID": {
					"description": "The ID of a managed asset. The resource focal point for a smart contract.",
					"type": "string"
				},
				"timestamp": {
					"description": "Datetime of the transaction",
					"type": "string"
				},
				"issuing": {
					"description": "Issuing of the transaction",
					"type": "number"
				},
				"receiver": {
					"description": "Receiver of the transaction",
					"type": "string"
				},
				"type": {
					"description": "Type of the transaction",
					"type": "string"
				},
				"acquirer": {
					"description": "Acquirer of the transaction",
					"type": "string"
				},
				"acquirerBank": {
					"description": "AcquirerBank of the transaction",
					"type": "string"
				},
				"issuer": {
					"description": "Issuer of the transaction",
					"type": "string"
				},
				"amount": {
					"description": "Amount involved in the transaction",
					"type": "number"
				},
				"currency": {
					"description": "Currency of the amount specified",
					"type": "number"
				},
				"origin": {
					"description": "Origin of the transaction",
					"type": "string"
				}
			},
			"required": ["txID",
				"timestamp",
				"issuing",
				"receiver",
				"type",
				"acquirer",
				"acquirerBank",
				"issuer",
				"amount",
				"currency",
				"origin"
			],
			"type": "object"
		},
		"initEvent": {
			"description": "event sent to init on deployment",
			"properties": {
				"transaction": {
					"default": "InitChainCode",
					"description": "Init of the current contract",
					"type": "string"
				},
				"version": {
					"description": "The ID of a managed asset. The resource focal point for a smart contract.",
					"type": "string"
				}
			},
			"required": ["version"],
			"type": "object"
		},
		"state": {
			"description": "A set of fields that constitute the complete asset state.",
			"properties": {
				"txID": {
					"description": "The ID of a managed asset. The resource focal point for a smart contract.",
					"type": "string"
				},
				"timestamp": {
					"description": "Datetime of the transaction",
					"type": "string"
				},
				"issuing": {
					"description": "Issuing of the transaction",
					"type": "number"
				},
				"receiver": {
					"description": "Receiver of the transaction",
					"type": "string"
				},
				"type": {
					"description": "Type of the transaction",
					"type": "string"
				},
				"acquirer": {
					"description": "Acquirer of the transaction",
					"type": "string"
				},
				"acquirerBank": {
					"description": "AcquirerBank of the transaction",
					"type": "string"
				},
				"issuer": {
					"description": "Issuer of the transaction",
					"type": "string"
				},
				"amount": {
					"description": "Amount involved in the transaction",
					"type": "number"
				},
				"currency": {
					"description": "Currency of the amount specified",
					"type": "number"
				},
				"origin": {
					"description": "Origin of the transaction",
					"type": "string"
				}
			},
			"type": "object"
		}
	}
}`