package jsonschema

import (
	"golang.org/x/crypto/bcrypt"
)

// this is version
var TableVersion = map[string]uint16{
	"versions":               1,
	"auth_token":             1,
	"applications":           1,
	"agent_location_session": 1,
	"alias":                  1,
	"global_settings":        1,
	"hepsub_mapping_schema":  1,
	"mapping_schema":         1,
	"users":                  1,
	"user_settings":          1,
}

var MinimumClickHouse = 10

var DashboardHome = `{"id":"home","name":"Home","alias":"home","selectedItem":"","title":"Home","weight":10,"widgets":[{"x":0,"y":0,"cols":2,"rows":1,"name":"clock","title":"clock","id":"clock214","output":{},"config":{"id":"clock214","datePattern":"YYYY-MM-DD","location":{"value":-60,"offset":"+1","name":"Europe/Amsterdam","desc":"Central European Time"},"showseconds":false,"timePattern":"HH:mm:ss","title":"Home Clock"}},{"x":0,"y":1,"cols":2,"rows":3,"name":"display-results","title":"display-results","id":"display-results370","output":{},"config":{"id":"display-results370","title":"CALL SIP SEARCH","group":"Search","name":"protosearch","description":"Display Search Form component","refresh":false,"sizeX":2,"sizeY":2,"config":{"title":"CALL SIP SEARCH","searchbutton":true,"protocol_id":{"name":"SIP","value":1},"protocol_profile":{"name":"call","value":"call"}},"uuid":"ed426bd0-ff21-40f7-8852-58700abc3762","fields":[{"field_name":"data_header.from_user","hepid":1,"name":"1:call:data_header.from_user","selection":"SIP From user","type":"string"},{"field_name":"data_header.to_user","hepid":1,"name":"1:call:data_header.to_user","selection":"SIP To user","type":"string"},{"field_name":"data_header.method","hepid":1,"name":"1:call:data_header.method","selection":"SIP Method","type":"string"},{"field_name":"data_header.callid","hepid":1,"name":"1:call:data_header.callid","selection":"SIP Callid","type":"string"},{"field_name":"limit","hepid":1,"name":"1:call:limit","selection":"Query Limit","type":"string"},{"field_name":"targetResultsContainer","hepid":1,"name":"1:call:targetResultsContainer","selection":"Results Container","type":"string"}],"row":0,"col":1,"cols":2,"rows":2,"x":0,"y":1,"protocol_id":{"name":"SIP","value":100}}},{"x":2,"y":0,"cols":4,"rows":4,"name":"result","title":"result","id":"result560","output":{}}],"config":{"margins":[10,10],"columns":"6","pushing":true,"draggable":{"handle":".box-header"},"resizable":{"enabled":true,"handles":["n","e","s","w","ne","se","sw","nw"]}}}`

var CorrelationMappingdefault = `{"source_field": "data_header.callid",
      "lookup_id": 0,
      "lookup_type": "pubsub",
      "lookup_profile": "cdr",
      "lookup_field": "{\"data\":$source_field,\"fromts\":$fromts,\"tots\":$tots}",
	  "lookup_range": [-300, 200]}`

var CorrelationScriptdefault = `//this is javascript default`

var EmptyJson = `{}`

var DefaultAdminPassword, _ = bcrypt.GenerateFromPassword([]byte("sipcapture"), bcrypt.DefaultCost)
var DefaultSupportPassword, _ = bcrypt.GenerateFromPassword([]byte("sipcapture"), bcrypt.DefaultCost)

var GrafanaConfig = `{"host": "http://grafana:3000","user": "admin","password":"admin","token": "ABCDEFGHKLMN"}`
var PrometheusConfig = `{"host":"http://prometheus:9090/api/v1/"}`
var LokiConfig = `{"host":"http://loki:3100"}`

var AuthTypesConfig = `{"internal": {"name": "Internal","type": "internal","enable": true}}`

var ExportConfig = `{"openwindow": false,"tabpositon": "flow"}`
var TransactionConfig = `{"tabpositon": "flow"}`

var AgentObjectforAuthToken = `{
   "username": "test",
   "firstname": "Tester",
   "lastname": "Tester",
   "email": "tester@test.com",
   "usergroup": "user",
   "id": 1000,
   "partid": 10
  }`

var FieldsMapping1callold = `[
    {
        "id": "uuid",
        "type": "string",
        "name": "UUID",
        "index": "secondary",
        "form_type": "input",
        "alias": "uuid",
        "position": 1,
 	    "skip": false,
        "hide": false
    },{
        "id": "create_ts",
        "type": "integer",
        "name": "Timestamp",
        "form_type": "input",
        "index": "secondary",
        "alias": "micro_ts",
        "position": 2,
	    "skip": false,
        "hide": false
    },{
        "id": "callid",
        "type": "string",
        "form_type": "input",
        "name": "CallID",       
        "index": "wildcard",
        "alias": "sid",
        "sid_type": true,
        "position": 3,
	    "skip": false,
        "hide": false
    },{
        "id": "event",
        "type": "string",
        "name": "Method",        
        "index": "none",
        "alias": "method",
	    "form_type": "input",
        "form_default": [
            "INVITE",
            "BYE",
            "100",
            "200",
            "183",
            "CANCEL"
        ],
	    "skip": false,
        "hide": false,
        "method_type": true,
        "position": 4
    },{
     	"id": "source_ip",
        "type": "string",
        "name": "Source IP",        
        "index": "secondary",
        "alias": "srcIp",
	    "skip": false,
        "hide": false,
        "position": 5
    },{
        "id": "source_port",
        "type": "integer",
        "name": "Source Port",
        "alias": "srcPort", 
        "form_type": "input",       
        "index": "secondary",
	    "skip": false,
        "hide": false,
        "position": 6  	
    },{
        "id": "destination_ip",
        "type": "string",
        "name": "Destination IP",  
        "alias": "dstIp",      
        "index": "secondary",
        "form_type": "input",       
    	"skip": false,
        "hide": false,
        "position": 7
    },{
        "id": "destination_port",
        "type": "integer",
        "name": "Destination Port", 
        "alias": "dstPort",       
        "index": "secondary",
        "form_type": "input",       
	    "skip": false,
        "hide": false,
        "position": 8
    },{
        "id": "vlan",
        "type": "integer",
        "name": "Vlan",
        "alias": "vlan",       
        "form_type": "input",        
        "index": "secondary",
	    "skip": false,
        "hide": true,
        "position": 9
    },{
        "id": "capture_ip",
        "type": "string",
        "name": "Capture IP",
        "alias": "capture_ip",        
        "form_type": "input",       
        "index": "wildcard",
	    "skip": false,
        "hide": true,
        "position": 10
    },{
        "id": "proto",
        "type": "integer",
        "name": "Proto",
        "alias": "protocol",   
        "form_type": "input",            
        "index": "none",
	    "skip": false,
        "hide": false,
        "position": 11
    },{
        "id": "node",
        "name": "Node",
        "type": "string",
        "index": "none",
        "form_type": "multiselect",
        "form_default": [
            {
                "value": "localnode",
                "name": "Local node"
            }
        ],
        "_form_api": "/database/node/list",
        "system_param": true,
        "mapping": "param.location.node",
        "skip": true,
        "hide": true,
        "position": 12
    },{
        "id": "captid",
        "type": "integer",
        "name": "Capture ID",
        "alias": "captureId",  
        "form_type": "input",      
        "index": "none",
    	"skip": false,
        "hide": true,        
        "position": 13
    },{
        "id": "region_id",
        "type": "string",
        "name": "Region ID",
        "alias": "regionid",        
        "index": "none",
        "form_type": "input",
	    "skip": false,
        "hide": true,        
        "position": 14
    },{
        "id": "data",
        "type": "string",
        "name": "Data",
        "alias": "data",        
        "index": "wildcard",
        "form_type": "input",
        "skip": false,
        "hide": true,        
        "position": 15
    },{
        "id": "message",
        "type": "string",
        "name": "Message",
        "form_type": "input",
        "index": "wildcard",
        "alias": "raw",
	    "skip": false,
        "hide": true,
        "position": 16  
    },{
        "id": "data.from_user",
        "type": "string",
        "virtual": true,
        "name": "From User",
        "alias": "from_user",   
        "form_type": "input",     
        "index": "wildcard",
	    "skip": false,
        "hide": false,        
        "position": 17
    },{
        "id": "data.to_user",
        "type": "string",
        "virtual": true,
        "name": "To User",
        "alias": "to_user",  
        "form_type": "input",      
        "index": "wildcard",
        "skip": false,
        "hide": false,        
        "position": 17  
    },{
        "id": "data.ruri_user",
        "type": "string",
        "virtual": true,
        "name": "RURI User",
        "form_type": "input",
        "alias": "ruri_user",        
        "index": "wildcard",
	    "skip": false,
        "hide": true,        
        "position": 17  
    },{
        "id": "data.cseq",
        "type": "string",
        "virtual": true,
        "name": "Cseq",
        "form_type": "input",
        "alias": "cseq",        
        "index": "wildcard",
    	"skip": false,
        "hide": true,        
        "position": 17  
    }
]`

var FieldsMapping1registrationold = `[
	{
	  "id": "sid",
	  "type": "string",
	  "index": "secondary",
	  "name": "Session ID",
	  "form_type": "input",
	  "position": 1,
	  "sid_type": true,
	  "hide": false
	},
	{
	  "id": "data_header.method",
	  "name": "SIP Method",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "form_default": [
		"INVITE",
		"BYE",
		"100",
		"200",
		"REGISTER",
    "CANCEL",
    "OPTIONS",
    "NOTIFY"
	  ],
	  "position": 2,
	  "skip": false,
	  "hide": false,
	  "method_type": true
	},
	{
	  "id": "protocol_header.correlation_id",
	  "name": "Correlation ID",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 3,
	  "skip": false,
	  "hide": true,
	  "sid_type": true
	},
	{
	  "id": "data_header.callid",
	  "name": "SIP Callid",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 4,
	  "skip": false,
	  "hide": true,
	  "sid_type": true
  },
	{
	  "id": "data_header.from_user",
	  "name": "SIP From user",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 6,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "data_header.to_user",
	  "name": "SIP To user",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 7,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.srcIp",
	  "name": "Source IP",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 8,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.srcPort",
	  "name": "Src Port",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 9,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.dstIp",
	  "name": "Destination IP",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 10,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.dstPort",
	  "name": "Dst Port",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 11,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.timeSeconds",
	  "name": "Timeseconds",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 12,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.timeUseconds",
	  "name": "Usecond time",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 13,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.payloadType",
	  "name": "Payload type",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 14,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.protocolFamily",
	  "name": "Proto Family",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 15,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.protocol",
	  "name": "Protocol Type",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 16,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.captureId",
	  "name": "Capture ID",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 17,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.capturePass",
	  "name": "Capture Pass",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 18,
	  "skip": true,
	  "hide": true
	},
	{
	  "id": "data_header.cseq",
	  "name": "SIP Cseq",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 19,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "data_header.from_tag",
	  "name": "SIP From tag",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 20,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "data_header.protocol",
	  "name": "SIP Protocol",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 21,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "raw",
	  "name": "SIP RAW",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 22,
	  "skip": true,
	  "hide": true
  },
  {
    "id": "node",
    "name": "Node",
    "type": "string",
    "index": "none",
    "form_type": "multiselect",
    "form_default": [
        {"value":"localnode","name":"Local node"}
    ],
    "_form_api": "/database/node/list",  
    "system_param": true,
    "mapping": "param.location.node",
    "position": 23,
    "skip": true,
    "hide": true
  }
  ]
`

var FieldsMapping1defaultold = `[
	{
	  "id": "sid",
	  "type": "string",
	  "index": "secondary",
	  "name": "Session ID",
	  "form_type": "input",
	  "position": 1,
	  "sid_type": true,
	  "hide": false
	},
	{
	  "id": "data_header.method",
	  "name": "SIP Method",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "form_default": [
		"INVITE",
		"BYE",
		"100",
		"200",
		"REGISTER",
    "CANCEL",
    "OPTIONS",
    "NOTIFY"
	  ],
	  "position": 2,
	  "skip": false,
	  "hide": false,
	  "method_type": true
	},
	{
	  "id": "protocol_header.correlation_id",
	  "name": "Correlation ID",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 3,
	  "skip": false,
	  "hide": true,
	  "sid_type": true
	},
	{
	  "id": "data_header.callid",
	  "name": "SIP Callid",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 4,
	  "skip": false,
	  "hide": true,
	  "sid_type": true
  },
	{
	  "id": "data_header.from_user",
	  "name": "SIP From user",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 6,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "data_header.to_user",
	  "name": "SIP To user",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 7,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.srcIp",
	  "name": "Source IP",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 8,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.srcPort",
	  "name": "Src Port",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 9,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.dstIp",
	  "name": "Destination IP",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 10,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.dstPort",
	  "name": "Dst Port",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 11,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.timeSeconds",
	  "name": "Timeseconds",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 12,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.timeUseconds",
	  "name": "Usecond time",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 13,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.payloadType",
	  "name": "Payload type",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 14,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.protocolFamily",
	  "name": "Proto Family",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 15,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.protocol",
	  "name": "Protocol Type",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 16,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.captureId",
	  "name": "Capture ID",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 17,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.capturePass",
	  "name": "Capture Pass",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 18,
	  "skip": true,
	  "hide": true
	},
	{
	  "id": "data_header.cseq",
	  "name": "SIP Cseq",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 19,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "data_header.from_tag",
	  "name": "SIP From tag",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 20,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "data_header.protocol",
	  "name": "SIP Protocol",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 21,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "raw",
	  "name": "SIP RAW",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 22,
	  "skip": true,
	  "hide": true
  },
  {
    "id": "node",
    "name": "Node",
    "type": "string",
    "index": "none",
    "form_type": "multiselect",
    "form_default": [
        {"value":"localnode","name":"Local node"}
    ],
    "_form_api": "/database/node/list",  
    "system_param": true,
    "mapping": "param.location.node",
    "position": 23,
    "skip": true,
    "hide": true
  }
  ]
`

/* NEW */
var FieldsMapping1call = `[
	{
	  "id": "sid",
	  "type": "string",
	  "index": "secondary",
	  "name": "Session ID",
	  "form_type": "input",
	  "position": 1,
	  "sid_type": true,
	  "hide": false
	},
	{
	  "id": "data_header.method",
	  "name": "SIP Method",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "form_default": [
		"INVITE",
		"BYE",
		"100",
		"200",
		"183",
		"CANCEL"
	  ],
	  "position": 2,
	  "skip": false,
	  "hide": false,
	  "method_type": true
	},
	{
	  "id": "protocol_header.correlation_id",
	  "name": "Correlation ID",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 3,
	  "skip": false,
	  "hide": true,
	  "sid_type": true
	},
	{
	  "id": "data_header.callid",
	  "name": "CallID",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 4,
	  "skip": false,
	  "hide": true,
	  "sid_type": true
  },
  {
    "id": "data_header.ruri_user",
    "name": "RURI user",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 5,
    "skip": false,
    "hide": true
  },
	{
	  "id": "data_header.from_user",
	  "name": "SIP From user",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 6,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "data_header.to_user",
	  "name": "SIP To user",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 7,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.srcIp",
	  "name": "Source IP",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 8,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.srcPort",
	  "name": "Src Port",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 9,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.dstIp",
	  "name": "Destination IP",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 10,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.dstPort",
	  "name": "Dst Port",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 11,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.timeSeconds",
	  "name": "Timeseconds",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 12,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.timeUseconds",
	  "name": "Usecond time",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 13,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.payloadType",
	  "name": "Payload type",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 14,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.protocolFamily",
	  "name": "Proto Family",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 15,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.protocol",
	  "name": "Protocol Type",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 16,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.captureId",
	  "name": "Capture ID",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 17,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.capturePass",
	  "name": "Capture Pass",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 18,
	  "skip": true,
	  "hide": true
	},
	{
	  "id": "data_header.cseq",
	  "name": "SIP Cseq",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 19,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "data_header.from_tag",
	  "name": "SIP From tag",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 20,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "data_header.protocol",
	  "name": "SIP Protocol",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 21,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "raw",
	  "name": "SIP RAW",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 22,
	  "skip": true,
	  "hide": true
  },
  {
    "id": "node",
    "name": "Node",
    "type": "string",
    "index": "none",
    "form_type": "multiselect",
    "form_default": [
        {"value":"localnode","name":"Local node"}
    ],
    "_form_api": "/database/node/list",
    "system_param": true,
    "mapping": "param.location.node",
    "position": 23,
    "skip": true,
    "hide": true
  }
  ]
`

var FieldsMapping1registration = `[
	{
	  "id": "sid",
	  "type": "string",
	  "index": "secondary",
	  "name": "Session ID",
	  "form_type": "input",
	  "position": 1,
	  "sid_type": true,
	  "hide": false
	},
	{
	  "id": "data_header.method",
	  "name": "SIP Method",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "form_default": [
		"INVITE",
		"BYE",
		"100",
		"200",
		"REGISTER",
    "CANCEL",
    "OPTIONS",
    "NOTIFY"
	  ],
	  "position": 2,
	  "skip": false,
	  "hide": false,
	  "method_type": true
	},
	{
	  "id": "protocol_header.correlation_id",
	  "name": "Correlation ID",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 3,
	  "skip": false,
	  "hide": true,
	  "sid_type": true
	},
	{
	  "id": "data_header.callid",
	  "name": "SIP Callid",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 4,
	  "skip": false,
	  "hide": true,
	  "sid_type": true
  },
	{
	  "id": "data_header.from_user",
	  "name": "SIP From user",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 6,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "data_header.to_user",
	  "name": "SIP To user",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 7,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.srcIp",
	  "name": "Source IP",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 8,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.srcPort",
	  "name": "Src Port",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 9,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.dstIp",
	  "name": "Destination IP",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 10,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.dstPort",
	  "name": "Dst Port",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 11,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.timeSeconds",
	  "name": "Timeseconds",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 12,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.timeUseconds",
	  "name": "Usecond time",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 13,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.payloadType",
	  "name": "Payload type",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 14,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.protocolFamily",
	  "name": "Proto Family",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 15,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.protocol",
	  "name": "Protocol Type",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 16,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.captureId",
	  "name": "Capture ID",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 17,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.capturePass",
	  "name": "Capture Pass",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 18,
	  "skip": true,
	  "hide": true
	},
	{
	  "id": "data_header.cseq",
	  "name": "SIP Cseq",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 19,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "data_header.from_tag",
	  "name": "SIP From tag",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 20,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "data_header.protocol",
	  "name": "SIP Protocol",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 21,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "raw",
	  "name": "SIP RAW",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 22,
	  "skip": true,
	  "hide": true
  },
  {
    "id": "node",
    "name": "Node",
    "type": "string",
    "index": "none",
    "form_type": "multiselect",
    "form_default": [
        {"value":"localnode","name":"Local node"}
    ],
    "_form_api": "/database/node/list",  
    "system_param": true,
    "mapping": "param.location.node",
    "position": 23,
    "skip": true,
    "hide": true
  }
  ]
`

var FieldsMapping1default = `[
	{
	  "id": "sid",
	  "type": "string",
	  "index": "secondary",
	  "name": "Session ID",
	  "form_type": "input",
	  "position": 1,
	  "sid_type": true,
	  "hide": false
	},
	{
	  "id": "data_header.method",
	  "name": "SIP Method",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "form_default": [
		"INVITE",
		"BYE",
		"100",
		"200",
		"REGISTER",
    "CANCEL",
    "OPTIONS",
    "NOTIFY"
	  ],
	  "position": 2,
	  "skip": false,
	  "hide": false,
	  "method_type": true
	},
	{
	  "id": "protocol_header.correlation_id",
	  "name": "Correlation ID",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 3,
	  "skip": false,
	  "hide": true,
	  "sid_type": true
	},
	{
	  "id": "data_header.callid",
	  "name": "SIP Callid",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 4,
	  "skip": false,
	  "hide": true,
	  "sid_type": true
  },
	{
	  "id": "data_header.from_user",
	  "name": "SIP From user",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 6,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "data_header.to_user",
	  "name": "SIP To user",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 7,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.srcIp",
	  "name": "Source IP",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 8,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.srcPort",
	  "name": "Src Port",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 9,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.dstIp",
	  "name": "Destination IP",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 10,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.dstPort",
	  "name": "Dst Port",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 11,
	  "skip": false,
	  "hide": false
	},
	{
	  "id": "protocol_header.timeSeconds",
	  "name": "Timeseconds",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 12,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.timeUseconds",
	  "name": "Usecond time",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 13,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.payloadType",
	  "name": "Payload type",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 14,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.protocolFamily",
	  "name": "Proto Family",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 15,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.protocol",
	  "name": "Protocol Type",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 16,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.captureId",
	  "name": "Capture ID",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 17,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "protocol_header.capturePass",
	  "name": "Capture Pass",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 18,
	  "skip": true,
	  "hide": true
	},
	{
	  "id": "data_header.cseq",
	  "name": "SIP Cseq",
	  "type": "integer",
	  "index": "none",
	  "form_type": "input",
	  "position": 19,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "data_header.from_tag",
	  "name": "SIP From tag",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 20,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "data_header.protocol",
	  "name": "SIP Protocol",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 21,
	  "skip": false,
	  "hide": true
	},
	{
	  "id": "raw",
	  "name": "SIP RAW",
	  "type": "string",
	  "index": "none",
	  "form_type": "input",
	  "position": 22,
	  "skip": true,
	  "hide": true
  },
  {
    "id": "node",
    "name": "Node",
    "type": "string",
    "index": "none",
    "form_type": "multiselect",
    "form_default": [
        {"value":"localnode","name":"Local node"}
    ],
    "_form_api": "/database/node/list",  
    "system_param": true,
    "mapping": "param.location.node",
    "position": 23,
    "skip": true,
    "hide": true
  }
  ]
`

var FieldsMapping34default = `[
  {
    "id": "sid",
    "type": "string",
    "index": "secondary",
    "name": "Session ID",
    "form_type": "input",
    "position": 1,
    "skip": false,
    "hide": false,
    "sid_type": true
  },
  {
    "id": "protocol_header.correlation_id",
    "name": "Correlation ID",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 2,
    "skip": false,
    "hide": true,
    "sid_type": true
  },
  {
    "id": "protocol_header.srcIp",
    "name": "Source IP",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 3,
    "skip": false,
    "hide": false
  },
  {
    "id": "protocol_header.srcPort",
    "name": "Src Port",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 4,
    "skip": false,
    "hide": false
  },
  {
    "id": "protocol_header.dstIp",
    "name": "Destination IP",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 5,
    "skip": false,
    "hide": false
  },
  {
    "id": "protocol_header.dstPort",
    "name": "Dst Port",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 6,
    "skip": false,
    "hide": false
  },
  {
    "id": "protocol_header.timeSeconds",
    "name": "Timeseconds",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 7,
    "skip": false,
    "hide": true
  },
  {
    "id": "protocol_header.timeUseconds",
    "name": "Usecond time",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 8,
    "skip": false,
    "hide": true
  },
  {
    "id": "protocol_header.protocolFamily",
    "name": "Proto Family",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 9,
    "skip": false,
    "hide": true
  },
  {
    "id": "protocol_header.protocol",
    "name": "Protocol Type",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 10,
    "skip": false,
    "hide": true
  },
  {
    "id": "protocol_header.payloadType",
    "name": "Payload type",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 11,
    "skip": false,
    "hide": true
  },
  {
    "id": "protocol_header.captureId",
    "name": "Capture ID",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 12,
    "skip": false,
    "hide": true
  },
  {
    "id": "protocol_header.capturePass",
    "name": "Capture Pass",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 13,
    "skip": true,
    "hide": true
  },
  {
    "id": "raw",
    "name": "RAW",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 14,
    "skip": true,
    "hide": true
  }
]`

/* CDR Call */
var FieldsMapping60callOld = `[    
      {
          "id": "uuid",
          "type": "string",
          "index": "secondary",
          "name": "Session ID",
          "form_type": "input",
          "position": 1,
          "sid_type": true,
          "hide": false
      },
      {
        "id": "cdr_start",
        "type": "integer",        
        "name": "cdr_start",
        "form_type": "input",
        "system_param": true,
        "mapping": "param.range.time",
        "position": 1,
        "skip": true,
        "hide": true
      },
      {
          "id": "callid",
          "name": "Call-ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 2,
          "sid_type": true,
          "skip": false,
          "hide": false
      },
      {
          "id": "captid",
          "name": "Capt ID",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 3,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_connect",
          "name": "CDR connect",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 4,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_progress",
          "name": "CDR progress",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 5,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_ringing",
          "name": "CDR ringing",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 6,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_start",
          "name": "CDR start",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 7,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_stop",
          "name": "CDR stop",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 8,
          "skip": false,
          "hide": true
      },
      {
          "id": "contact_user",
          "name": "Contact user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 9,
          "skip": false,
          "hide": true
      },
      {
          "id": "correlation_id",
          "name": "Correlation ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 10,
          "skip": false,
          "hide": true
      },
      {
          "id": "Correlations",
          "name": "Contact user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 11,
          "skip": false,
          "hide": true
      },
      {
          "id": "create_date",
          "name": "Create date",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 12,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_1",
          "name": "custom_1",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 13,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_2",
          "name": "custom_2",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 14,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_3",
          "name": "custom_3",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 15,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_4",
          "name": "custom_4",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 16,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_5",
          "name": "custom_5",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 17,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_6",
          "name": "custom_6",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 18,
          "skip": false,
          "hide": true
      },
      {
          "id": "data",
          "name": "Data",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 19,
          "skip": false,
          "hide": true
      },
      {
          "id": "dest_cc",
          "name": "Dest. Country",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 20,
          "skip": false,
          "hide": true
      },
      {
          "id": "destination_ip",
          "name": "Dest IP",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 21,
          "skip": false,
          "hide": false
      },
      {
          "id": "destination_port",
          "name": "Dest Port",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 22,
          "skip": false,
          "hide": false
      },
      {
          "id": "duration",
          "name": "Duration",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 23,
          "skip": false,
          "hide": false
      },
      {
          "id": "event",
          "name": "Event",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 24,
          "skip": false,
          "hide": true
      },
      {
          "id": "family",
          "name": "Family",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 25,
          "skip": false,
          "hide": true
      },
      {
          "id": "from_user",
          "name": "From User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 26,
          "skip": false,
          "hide": false
      },
      {
          "id": "geo_cc",
          "name": "Geo Country",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 27,
          "skip": false,
          "hide": true
      },
      {
          "id": "ipgroup_in",
          "name": "IP Group In",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 28,
          "skip": false,
          "hide": true
      },
      {
          "id": "ipgroup_out",
          "name": "IP Group Out",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 29,
          "skip": false,
          "hide": true
      },
      {
          "id": "jitter",
          "name": "Jitter",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 30,
          "skip": false,
          "hide": true
      },
      {
          "id": "mos",
          "name": "Mos",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 31,
          "skip": false,
          "hide": false
      },
      {
          "id": "msg_info",
          "name": "Msg Info",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 32,
          "skip": false,
          "hide": true
      },
      {
          "id": "pdd",
          "name": "PDD",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 33,
          "skip": false,
          "hide": true
      },
      {
          "id": "pid_user",
          "name": "PID user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 34,
          "skip": false,
          "hide": true
      },
      {
          "id": "pl",
          "name": "Packet Loss",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 35,
          "skip": false,
          "hide": true
      },
      {
          "id": "region_id",
          "name": "Regiond ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 36,
          "skip": false,
          "hide": true
      },
      {
          "id": "rtp_stat_a",
          "name": "RTP Stats A",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 37,
          "skip": false,
          "hide": true
      },
      {
          "id": "rtp_stat_b",
          "name": "RTP Stats B",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 38,
          "skip": false,
          "hide": true
      },
      {
          "id": "ruri_user",
          "name": "RURI User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 39,
          "skip": false,
          "hide": false
      },
      {
          "id": "server_type_in",
          "name": "Server Type IN",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 40,
          "skip": false,
          "hide": true
      },
      {
          "id": "server_type_out",
          "name": "Server Type OUT",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 41,
          "skip": false,
          "hide": true
      },
      {
          "id": "source_ip",
          "name": "Source IP",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 42,
          "skip": false,
          "hide": false
      },
      {
          "id": "source_port",
          "name": "Source Port",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 43,
          "skip": false,
          "hide": false
      },
      {
          "id": "srd",
          "name": "SRD",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 44,
          "skip": false,
          "hide": true
      },
      {
          "id": "sss",
          "name": "SSS",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 45,
          "skip": false,
          "hide": true
      },
      {
          "id": "status",
          "name": "Status",
          "type": "integer",
          "index": "none",
          "form_type": "select",
          "form_default": [
              {
                  "name": "Init",
                  "value": 1
              },
              {
                  "name": "Unauthorized",
                  "value": 2
              },
              {
                  "name": "Progress",
                  "value": 3
              },
              {
                  "name": "Ringing",
                  "value": 4
              },
              {
                  "name": "Connected",
                  "value": 5
              },
              {
                  "name": "Moved",
                  "value": 6
              },
              {
                  "name": "Busy",
                  "value": 7
              },
              {
                  "name": "User Failure",
                  "value": 8
              },
              {
                  "name": "Hard Failure",
                  "value": 9
              },
              {
                  "name": "Finished",
                  "value": 10
              },
              {
                  "name": "Canceled",
                  "value": 11
              },
              {
                  "name": "Timed Out",
                  "value": 12
              },
              {
                  "name": "Bad Termination",
                  "value": 13
              },
              {
                  "name": "Declined",
                  "value": 14
              },
              {
                  "name": "Unknown",
                  "value": 15
              }
          ],
          "position": 46,
          "skip": false,
          "hide": false
      },
      {
          "id": "table",
          "name": "Table",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 47,
          "skip": false,
          "hide": true
      },
      {
          "id": "termcode",
          "name": "Term. Code",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 48,
          "skip": false,
          "hide": true
      },
      {
          "id": "to_user",
          "name": "To User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 49,
          "skip": false,
          "hide": true
      },
      {
          "id": "uas",
          "name": "UAS",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 50,
          "skip": false,
          "hide": true
      },
      {
          "id": "update_ts",
          "name": "Update TS",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 51,
          "skip": false,
          "hide": true
      },
      {
          "id": "usergroup",
          "name": "User Group",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 52,
          "skip": false,
          "hide": true
      },
      {
          "id": "vlan",
          "name": "Vlan",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 53,
          "skip": false,
          "hide": true
      },
      {
          "id": "vqr_a",
          "name": "VQR A",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 54,
          "skip": false,
          "hide": true
      },
      {
          "id": "vqr_b",
          "name": "VQR B",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 55,
          "skip": false,
          "hide": true
      },
      {
          "id": "vst",
          "name": "Vst",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 56,
          "skip": false,
          "hide": true
      },
      {
          "id": "xgroup",
          "name": "XGroup",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 57,
          "skip": false,
          "hide": true
      }
]`

/* CDR Call */
var FieldsMapping60registrationOld = `[    
      {
          "id": "uuid",
          "type": "string",
          "index": "secondary",
          "name": "Session ID",
          "form_type": "input",
          "position": 1,
          "sid_type": true,
          "hide": false
      },
      {
        "id": "cdr_start",
        "type": "integer",        
        "name": "cdr_start",
        "form_type": "input",
        "system_param": true,
        "mapping": "param.range.time",
        "position": 1,
        "skip": true,
        "hide": true
      },
      {
          "id": "callid",
          "name": "Call-ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 2,
          "sid_type": true,
          "skip": false,
          "hide": false
      },    
      {
          "id": "captid",
          "name": "Capt ID",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 3,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_connect",
          "name": "CDR connect",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 4,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_progress",
          "name": "CDR progress",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 5,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_ringing",
          "name": "CDR ringing",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 6,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_start",
          "name": "CDR start",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 7,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_stop",
          "name": "CDR stop",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 8,
          "skip": false,
          "hide": true
      },
      {
          "id": "contact_user",
          "name": "Contact user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 9,
          "skip": false,
          "hide": true
      },
      {
          "id": "correlation_id",
          "name": "Correlation ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 10,
          "skip": false,
          "hide": true
      },
      {
          "id": "Correlations",
          "name": "Contact user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 11,
          "skip": false,
          "hide": true
      },
      {
          "id": "create_date",
          "name": "Create date",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 12,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_1",
          "name": "custom_1",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 13,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_2",
          "name": "custom_2",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 14,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_3",
          "name": "custom_3",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 15,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_4",
          "name": "custom_4",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 16,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_5",
          "name": "custom_5",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 17,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_6",
          "name": "custom_6",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 18,
          "skip": false,
          "hide": true
      },
      {
          "id": "data",
          "name": "Data",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 19,
          "skip": false,
          "hide": true
      },
      {
          "id": "dest_cc",
          "name": "Dest. Country",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 20,
          "skip": false,
          "hide": true
      },
      {
          "id": "destination_ip",
          "name": "Dest IP",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 21,
          "skip": false,
          "hide": false
      },
      {
          "id": "destination_port",
          "name": "Dest Port",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 22,
          "skip": false,
          "hide": false
      },
      {
          "id": "duration",
          "name": "Duration",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 23,
          "skip": false,
          "hide": false
      },
      {
          "id": "event",
          "name": "Event",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 24,
          "skip": false,
          "hide": true
      },
      {
          "id": "family",
          "name": "Family",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 25,
          "skip": false,
          "hide": true
      },
      {
          "id": "from_user",
          "name": "From User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 26,
          "skip": false,
          "hide": false
      },
      {
          "id": "geo_cc",
          "name": "Geo Country",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 27,
          "skip": false,
          "hide": true
      },
      {
          "id": "ipgroup_in",
          "name": "IP Group In",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 28,
          "skip": false,
          "hide": true
      },
      {
          "id": "ipgroup_out",
          "name": "IP Group Out",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 29,
          "skip": false,
          "hide": true
      },
      {
          "id": "jitter",
          "name": "Jitter",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 30,
          "skip": false,
          "hide": true
      },
      {
          "id": "mos",
          "name": "Mos",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 31,
          "skip": false,
          "hide": false
      },
      {
          "id": "msg_info",
          "name": "Msg Info",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 32,
          "skip": false,
          "hide": true
      },
      {
          "id": "pdd",
          "name": "PDD",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 33,
          "skip": false,
          "hide": true
      },
      {
          "id": "pid_user",
          "name": "PID user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 34,
          "skip": false,
          "hide": true
      },
      {
          "id": "pl",
          "name": "Packet Loss",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 35,
          "skip": false,
          "hide": true
      },
      {
          "id": "region_id",
          "name": "Regiond ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 36,
          "skip": false,
          "hide": true
      },
      {
          "id": "rtp_stat_a",
          "name": "RTP Stats A",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 37,
          "skip": false,
          "hide": true
      },
      {
          "id": "rtp_stat_b",
          "name": "RTP Stats B",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 38,
          "skip": false,
          "hide": true
      },
      {
          "id": "ruri_user",
          "name": "RURI User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 39,
          "skip": false,
          "hide": false
      },
      {
          "id": "server_type_in",
          "name": "Server Type IN",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 40,
          "skip": false,
          "hide": true
      },
      {
          "id": "server_type_out",
          "name": "Server Type OUT",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 41,
          "skip": false,
          "hide": true
      },
      {
          "id": "source_ip",
          "name": "Source IP",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 42,
          "skip": false,
          "hide": false
      },
      {
          "id": "source_port",
          "name": "Source Port",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 43,
          "skip": false,
          "hide": false
      },
      {
          "id": "srd",
          "name": "SRD",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 44,
          "skip": false,
          "hide": true
      },
      {
          "id": "sss",
          "name": "SSS",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 45,
          "skip": false,
          "hide": true
      },
      {
          "id": "status",
          "name": "Status",
          "type": "integer",
          "index": "none",
          "form_type": "select",
          "form_default": [
              {
                  "name": "Init",
                  "value": 1
              },
              {
                  "name": "Unauthorized",
                  "value": 2
              },
              {
                  "name": "Progress",
                  "value": 3
              },
              {
                  "name": "Ringing",
                  "value": 4
              },
              {
                  "name": "Connected",
                  "value": 5
              },
              {
                  "name": "Moved",
                  "value": 6
              },
              {
                  "name": "Busy",
                  "value": 7
              },
              {
                  "name": "User Failure",
                  "value": 8
              },
              {
                  "name": "Hard Failure",
                  "value": 9
              },
              {
                  "name": "Finished",
                  "value": 10
              },
              {
                  "name": "Canceled",
                  "value": 11
              },
              {
                  "name": "Timed Out",
                  "value": 12
              },
              {
                  "name": "Bad Termination",
                  "value": 13
              },
              {
                  "name": "Declined",
                  "value": 14
              },
              {
                  "name": "Unknown",
                  "value": 15
              }
          ],
          "position": 46,
          "skip": false,
          "hide": false
      },
      {
          "id": "table",
          "name": "Table",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 47,
          "skip": false,
          "hide": true
      },
      {
          "id": "termcode",
          "name": "Term. Code",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 48,
          "skip": false,
          "hide": true
      },
      {
          "id": "to_user",
          "name": "To User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 49,
          "skip": false,
          "hide": true
      },
      {
          "id": "uas",
          "name": "UAS",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 50,
          "skip": false,
          "hide": true
      },
      {
          "id": "update_ts",
          "name": "Update TS",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 51,
          "skip": false,
          "hide": true
      },
      {
          "id": "usergroup",
          "name": "User Group",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 52,
          "skip": false,
          "hide": true
      },
      {
          "id": "vlan",
          "name": "Vlan",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 53,
          "skip": false,
          "hide": true
      },
      {
          "id": "vqr_a",
          "name": "VQR A",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 54,
          "skip": false,
          "hide": true
      },
      {
          "id": "vqr_b",
          "name": "VQR B",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 55,
          "skip": false,
          "hide": true
      },
      {
          "id": "vst",
          "name": "Vst",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 56,
          "skip": false,
          "hide": true
      },
      {
          "id": "xgroup",
          "name": "XGroup",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 57,
          "skip": false,
          "hide": true
      }
]`

/* CDR Call */
var FieldsMapping60call = `[    
      {
          "id": "uuid",
          "type": "string",
          "index": "secondary",
          "name": "Session ID",
          "form_type": "input",
          "position": 1,
          "sid_type": true,
          "hide": false
      },
      {
        "id": "cdr_start",
        "type": "integer",        
        "name": "cdr_start",
        "form_type": "input",
        "system_param": true,
        "mapping": "param.range.time",
        "position": 1,
        "skip": true,
        "hide": true
      },
      {
          "id": "callid",
          "name": "Call-ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 2,
          "sid_type": true,
          "skip": false,
          "hide": false
      },     
      {
          "id": "captid",
          "name": "Capt ID",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 3,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_connect",
          "name": "CDR connect",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 4,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_progress",
          "name": "CDR progress",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 5,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_ringing",
          "name": "CDR ringing",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 6,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_start",
          "name": "CDR start",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 7,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_stop",
          "name": "CDR stop",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 8,
          "skip": false,
          "hide": true
      },
      {
          "id": "contact_user",
          "name": "Contact user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 9,
          "skip": false,
          "hide": true
      },
      {
          "id": "correlation_id",
          "name": "Correlation ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 10,
          "skip": false,
          "hide": true
      },
      {
          "id": "Correlations",
          "name": "Contact user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 11,
          "skip": false,
          "hide": true
      },
      {
          "id": "create_date",
          "name": "Create date",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 12,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_1",
          "name": "custom_1",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 13,
          "skip": false,
          "hide": true
      },      
      {
          "id": "data",
          "name": "Data",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 19,
          "skip": false,
          "hide": true
      },
      {
          "id": "dest_cc",
          "name": "Dest. Country",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 20,
          "skip": false,
          "hide": true
      },
      {
          "id": "destination_ip",
          "name": "Dest IP",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 21,
          "skip": false,
          "hide": false
      },
      {
          "id": "destination_port",
          "name": "Dest Port",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 22,
          "skip": false,
          "hide": false
      },
      {
          "id": "duration",
          "name": "Duration",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 23,
          "skip": false,
          "hide": false
      },
      {
          "id": "event",
          "name": "Event",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 24,
          "skip": false,
          "hide": true
      },
      {
          "id": "family",
          "name": "Family",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 25,
          "skip": false,
          "hide": true
      },
      {
          "id": "from_user",
          "name": "From User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 26,
          "skip": false,
          "hide": false
      },
      {
          "id": "geo_cc",
          "name": "Geo Country",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 27,
          "skip": false,
          "hide": true
      },
      {
          "id": "ipgroup_in",
          "name": "IP Group In",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 28,
          "skip": false,
          "hide": true
      },
      {
          "id": "ipgroup_out",
          "name": "IP Group Out",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 29,
          "skip": false,
          "hide": true
      },
      {
          "id": "jitter",
          "name": "Jitter",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 30,
          "skip": false,
          "hide": true
      },
      {
          "id": "mos",
          "name": "Mos",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 31,
          "skip": false,
          "hide": false
      },
      {
          "id": "msg_info",
          "name": "Msg Info",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 32,
          "skip": false,
          "hide": true
      },
      {
          "id": "pdd",
          "name": "PDD",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 33,
          "skip": false,
          "hide": true
      },
      {
          "id": "pid_user",
          "name": "PID user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 34,
          "skip": false,
          "hide": true
      },
      {
          "id": "pl",
          "name": "Packet Loss",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 35,
          "skip": false,
          "hide": true
      },
      {
          "id": "region_id",
          "name": "Regiond ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 36,
          "skip": false,
          "hide": true
      },
      {
          "id": "rtp_stat_a",
          "name": "RTP Stats A",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 37,
          "skip": false,
          "hide": true
      },
      {
          "id": "rtp_stat_b",
          "name": "RTP Stats B",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 38,
          "skip": false,
          "hide": true
      },
      {
          "id": "ruri_user",
          "name": "RURI User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 39,
          "skip": false,
          "hide": false
      },
      {
          "id": "server_type_in",
          "name": "Server Type IN",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 40,
          "skip": false,
          "hide": true
      },
      {
          "id": "server_type_out",
          "name": "Server Type OUT",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 41,
          "skip": false,
          "hide": true
      },
      {
          "id": "source_ip",
          "name": "Source IP",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 42,
          "skip": false,
          "hide": false
      },
      {
          "id": "source_port",
          "name": "Source Port",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 43,
          "skip": false,
          "hide": false
      },
      {
          "id": "srd",
          "name": "SRD",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 44,
          "skip": false,
          "hide": true
      },
      {
          "id": "sss",
          "name": "SSS",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 45,
          "skip": false,
          "hide": true
      },
      {
          "id": "status",
          "name": "Status",
          "type": "integer",
          "index": "none",
          "form_type": "select",
          "form_default": [
              {
                  "name": "Init",
                  "value": 1
              },
              {
                  "name": "Unauthorized",
                  "value": 2
              },
              {
                  "name": "Progress",
                  "value": 3
              },
              {
                  "name": "Ringing",
                  "value": 4
              },
              {
                  "name": "Connected",
                  "value": 5
              },
              {
                  "name": "Moved",
                  "value": 6
              },
              {
                  "name": "Busy",
                  "value": 7
              },
              {
                  "name": "User Failure",
                  "value": 8
              },
              {
                  "name": "Hard Failure",
                  "value": 9
              },
              {
                  "name": "Finished",
                  "value": 10
              },
              {
                  "name": "Canceled",
                  "value": 11
              },
              {
                  "name": "Timed Out",
                  "value": 12
              },
              {
                  "name": "Bad Termination",
                  "value": 13
              },
              {
                  "name": "Declined",
                  "value": 14
              },
              {
                  "name": "Unknown",
                  "value": 15
              }
          ],
          "position": 46,
          "skip": false,
          "hide": false
      },
      {
          "id": "table",
          "name": "Table",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 47,
          "skip": false,
          "hide": true
      },
      {
          "id": "termcode",
          "name": "Term. Code",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 48,
          "skip": false,
          "hide": true
      },
      {
          "id": "to_user",
          "name": "To User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 49,
          "skip": false,
          "hide": true
      },
      {
          "id": "uas",
          "name": "UAS",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 50,
          "skip": false,
          "hide": true
      },
      {
          "id": "update_ts",
          "name": "Update TS",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 51,
          "skip": false,
          "hide": true
      },
      {
          "id": "usergroup",
          "name": "User Group",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 52,
          "skip": false,
          "hide": true
      },
      {
          "id": "vlan",
          "name": "Vlan",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 53,
          "skip": false,
          "hide": true
      },
      {
          "id": "vqr_a",
          "name": "VQR A",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 54,
          "skip": false,
          "hide": true
      },
      {
          "id": "vqr_b",
          "name": "VQR B",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 55,
          "skip": false,
          "hide": true
      },
      {
          "id": "vst",
          "name": "Vst",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 56,
          "skip": false,
          "hide": true
      },
      {
          "id": "xgroup",
          "name": "XGroup",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 57,
          "skip": false,
          "hide": true
      }
]`

/* CDR Call */
var FieldsMapping60registration = `[    
      {
          "id": "uuid",
          "type": "string",
          "index": "secondary",
          "name": "Session ID",
          "form_type": "input",
          "position": 1,
          "sid_type": true,
          "hide": false
      },
      {
        "id": "cdr_start",
        "type": "integer",        
        "name": "cdr_start",
        "form_type": "input",
        "system_param": true,
        "mapping": "param.range.time",
        "position": 1,
        "skip": true,
        "hide": true
      },
      {
          "id": "callid",
          "name": "Call-ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 2,
          "sid_type": true,
          "skip": false,
          "hide": false
      },    
      {
          "id": "captid",
          "name": "Capt ID",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 3,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_connect",
          "name": "CDR connect",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 4,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_progress",
          "name": "CDR progress",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 5,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_ringing",
          "name": "CDR ringing",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 6,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_start",
          "name": "CDR start",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 7,
          "skip": false,
          "hide": true
      },
      {
          "id": "cdr_stop",
          "name": "CDR stop",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 8,
          "skip": false,
          "hide": true
      },
      {
          "id": "contact_user",
          "name": "Contact user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 9,
          "skip": false,
          "hide": true
      },
      {
          "id": "correlation_id",
          "name": "Correlation ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 10,
          "skip": false,
          "hide": true
      },
      {
          "id": "Correlations",
          "name": "Contact user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 11,
          "skip": false,
          "hide": true
      },
      {
          "id": "create_date",
          "name": "Create date",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 12,
          "skip": false,
          "hide": true
      },
      {
          "id": "custom_1",
          "name": "custom_1",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 13,
          "skip": false,
          "hide": true
      },      
      {
          "id": "data",
          "name": "Data",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 19,
          "skip": false,
          "hide": true
      },
      {
          "id": "dest_cc",
          "name": "Dest. Country",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 20,
          "skip": false,
          "hide": true
      },
      {
          "id": "destination_ip",
          "name": "Dest IP",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 21,
          "skip": false,
          "hide": false
      },
      {
          "id": "destination_port",
          "name": "Dest Port",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 22,
          "skip": false,
          "hide": false
      },
      {
          "id": "duration",
          "name": "Duration",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 23,
          "skip": false,
          "hide": false
      },
      {
          "id": "event",
          "name": "Event",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 24,
          "skip": false,
          "hide": true
      },
      {
          "id": "family",
          "name": "Family",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 25,
          "skip": false,
          "hide": true
      },
      {
          "id": "from_user",
          "name": "From User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 26,
          "skip": false,
          "hide": false
      },
      {
          "id": "geo_cc",
          "name": "Geo Country",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 27,
          "skip": false,
          "hide": true
      },
      {
          "id": "ipgroup_in",
          "name": "IP Group In",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 28,
          "skip": false,
          "hide": true
      },
      {
          "id": "ipgroup_out",
          "name": "IP Group Out",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 29,
          "skip": false,
          "hide": true
      },
      {
          "id": "jitter",
          "name": "Jitter",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 30,
          "skip": false,
          "hide": true
      },
      {
          "id": "mos",
          "name": "Mos",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 31,
          "skip": false,
          "hide": false
      },
      {
          "id": "msg_info",
          "name": "Msg Info",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 32,
          "skip": false,
          "hide": true
      },
      {
          "id": "pdd",
          "name": "PDD",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 33,
          "skip": false,
          "hide": true
      },
      {
          "id": "pid_user",
          "name": "PID user",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 34,
          "skip": false,
          "hide": true
      },
      {
          "id": "pl",
          "name": "Packet Loss",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 35,
          "skip": false,
          "hide": true
      },
      {
          "id": "region_id",
          "name": "Regiond ID",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 36,
          "skip": false,
          "hide": true
      },
      {
          "id": "rtp_stat_a",
          "name": "RTP Stats A",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 37,
          "skip": false,
          "hide": true
      },
      {
          "id": "rtp_stat_b",
          "name": "RTP Stats B",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 38,
          "skip": false,
          "hide": true
      },
      {
          "id": "ruri_user",
          "name": "RURI User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 39,
          "skip": false,
          "hide": false
      },
      {
          "id": "server_type_in",
          "name": "Server Type IN",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 40,
          "skip": false,
          "hide": true
      },
      {
          "id": "server_type_out",
          "name": "Server Type OUT",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 41,
          "skip": false,
          "hide": true
      },
      {
          "id": "source_ip",
          "name": "Source IP",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 42,
          "skip": false,
          "hide": false
      },
      {
          "id": "source_port",
          "name": "Source Port",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 43,
          "skip": false,
          "hide": false
      },
      {
          "id": "srd",
          "name": "SRD",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 44,
          "skip": false,
          "hide": true
      },
      {
          "id": "sss",
          "name": "SSS",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 45,
          "skip": false,
          "hide": true
      },
      {
          "id": "status",
          "name": "Status",
          "type": "integer",
          "index": "none",
          "form_type": "select",
          "form_default": [
              {
                  "name": "Init",
                  "value": 1
              },
              {
                  "name": "Unauthorized",
                  "value": 2
              },
              {
                  "name": "Progress",
                  "value": 3
              },
              {
                  "name": "Ringing",
                  "value": 4
              },
              {
                  "name": "Connected",
                  "value": 5
              },
              {
                  "name": "Moved",
                  "value": 6
              },
              {
                  "name": "Busy",
                  "value": 7
              },
              {
                  "name": "User Failure",
                  "value": 8
              },
              {
                  "name": "Hard Failure",
                  "value": 9
              },
              {
                  "name": "Finished",
                  "value": 10
              },
              {
                  "name": "Canceled",
                  "value": 11
              },
              {
                  "name": "Timed Out",
                  "value": 12
              },
              {
                  "name": "Bad Termination",
                  "value": 13
              },
              {
                  "name": "Declined",
                  "value": 14
              },
              {
                  "name": "Unknown",
                  "value": 15
              }
          ],
          "position": 46,
          "skip": false,
          "hide": false
      },
      {
          "id": "table",
          "name": "Table",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 47,
          "skip": false,
          "hide": true
      },
      {
          "id": "termcode",
          "name": "Term. Code",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 48,
          "skip": false,
          "hide": true
      },
      {
          "id": "to_user",
          "name": "To User",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 49,
          "skip": false,
          "hide": true
      },
      {
          "id": "uas",
          "name": "UAS",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 50,
          "skip": false,
          "hide": true
      },
      {
          "id": "update_ts",
          "name": "Update TS",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 51,
          "skip": false,
          "hide": true
      },
      {
          "id": "usergroup",
          "name": "User Group",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 52,
          "skip": false,
          "hide": true
      },
      {
          "id": "vlan",
          "name": "Vlan",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 53,
          "skip": false,
          "hide": true
      },
      {
          "id": "vqr_a",
          "name": "VQR A",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 54,
          "skip": false,
          "hide": true
      },
      {
          "id": "vqr_b",
          "name": "VQR B",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 55,
          "skip": false,
          "hide": true
      },
      {
          "id": "vst",
          "name": "Vst",
          "type": "integer",
          "index": "none",
          "form_type": "input",
          "position": 56,
          "skip": false,
          "hide": true
      },
      {
          "id": "xgroup",
          "name": "XGroup",
          "type": "string",
          "index": "none",
          "form_type": "input",
          "position": 57,
          "skip": false,
          "hide": true
      }
]`

var FieldsMapping100default = `[
    {
      "id": "sid",
      "type": "string",
      "index": "secondary",
      "name": "Session ID",
      "form_type": "input",
      "position": 1,
      "skip": false,
      "hide": false,
      "sid_type": true
    },
    {
      "id": "protocol_header.correlation_id",
      "name": "Correlation ID",
      "type": "string",
      "index": "none",
      "form_type": "input",
      "position": 2,
      "skip": false,
      "hide": true,
      "sid_type": true
    },
    {
      "id": "protocol_header.srcIp",
      "name": "Source IP",
      "type": "string",
      "index": "none",
      "form_type": "input",
      "position": 3,
      "skip": false,
      "hide": false
    },
    {
      "id": "protocol_header.srcPort",
      "name": "Src Port",
      "type": "integer",
      "index": "none",
      "form_type": "input",
      "position": 4,
      "skip": false,
      "hide": false
    },
    {
      "id": "protocol_header.dstIp",
      "name": "Destination IP",
      "type": "string",
      "index": "none",
      "form_type": "input",
      "position": 5,
      "skip": false,
      "hide": false
    },
    {
      "id": "protocol_header.dstPort",
      "name": "Dst Port",
      "type": "integer",
      "index": "none",
      "form_type": "input",
      "position": 6,
      "skip": false,
      "hide": false
    },
    {
      "id": "protocol_header.timeSeconds",
      "name": "Timeseconds",
      "type": "integer",
      "index": "none",
      "form_type": "input",
      "position": 7,
      "skip": false,
      "hide": true
    },
    {
      "id": "protocol_header.timeUseconds",
      "name": "Usecond time",
      "type": "integer",
      "index": "none",
      "form_type": "input",
      "position": 8,
      "skip": false,
      "hide": true
    },
    {
      "id": "protocol_header.payloadType",
      "name": "Payload type",
      "type": "integer",
      "index": "none",
      "form_type": "input",
      "position": 9,
      "skip": false,
      "hide": true
    },
    {
      "id": "protocol_header.captureId",
      "name": "Capture ID",
      "type": "integer",
      "index": "none",
      "form_type": "input",
      "position": 10,
      "skip": false,
      "hide": true
    },
    {
      "id": "protocol_header.capturePass",
      "name": "Capture Pass",
      "type": "string",
      "index": "none",
      "form_type": "input",
      "position": 12,
      "skip": true,
      "hide": true
    },
    {
      "id": "protocol_header.protocolFamily",
      "name": "Proto Family",
      "type": "integer",
      "index": "none",
      "form_type": "input",
      "position": 13,
      "skip": false,
      "hide": true
    },
    {
      "id": "protocol_header.protocol",
      "name": "Protocol Type",
      "type": "integer",
      "index": "none",
      "form_type": "input",
      "position": 14,
      "skip": false,
      "hide": true
    },
    {
      "id": "raw",
      "name": "RAW",
      "type": "string",
      "index": "none",
      "form_type": "input",
      "position": 15,
      "skip": true,
      "hide": true
    }
]`

var FieldsMapping1000default = `[
  {
    "id": "sid",
    "type": "string",
    "index": "secondary",
    "name": "Session ID",
    "form_type": "input",
    "position": 1,
    "skip": false,
    "hide": false,
    "sid_type": true
  },
  {
    "id": "protocol_header.address",
    "name": "Proto Address",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 2,
    "skip": false,
    "hide": false
  },
  {
    "id": "data_header.family",
    "name": "Family",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 3,
    "skip": false,
    "hide": true
  },
  {
    "id": "protocol_header.srcPort",
    "name": "Protocol port",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 4,
    "skip": false,
    "hide": false
  },
  {
    "id": "data_header.type",
    "name": "Data type",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 5,
    "skip": false,
    "hide": true
  },
  {
    "id": "data_header.handle",
    "name": "Data Handle",
    "type": "integer",
    "index": "none",
    "form_type": "input",
    "position": 6,
    "skip": false,
    "hide": true
  },
  {
    "id": "data_header.event",
    "name": "Data Event",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 7,
    "skip": false,
    "hide": false
  },
  {
    "id": "data_header.medium",
    "name": "Data Medium",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 8,
    "skip": false,
    "hide": true
  },
  {
    "id": "data_header.source",
    "name": "Data Source",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 9,
    "skip": false,
    "hide": false
  },
  {
    "id": "data_header.session",
    "name": "Data Session",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 10,
    "skip": false,
    "hide": false
  },
  {
    "id": "raw",
    "name": "RAW",
    "type": "string",
    "index": "none",
    "form_type": "input",
    "position": 11,
    "skip": true,
    "hide": true
  }
]`

var FieldsMapping2000loki = `[
	{
		  "id": "micro_ts",
		  "name": "Timeseconds",
		  "type": "integer",
		  "index": "none",
		  "form_type": "input",
		  "date_field": true,
		  "position": 1,
		  "skip": false,
		  "hide": false
	},{
		  "id": "custom_1",
		  "name": "Message",
		  "type": "string",
		  "index": "none",
		  "form_type": "input",
		  "position": 2,
		  "skip": false,
		  "hide": false,
		  "autoheight": true
	},{
		  "id": "custom_2",
		  "name": "Labels",
		  "type": "string",
		  "index": "none",
		  "form_type": "input",
		  "position": 3,
		  "skip": false,
		  "hide": false
	}
]`

var CorrelationMapping1callold = `[		
    {
        "source_field": "data_header.callid",
        "lookup_id": 100,
        "lookup_profile": "call",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "data_header.callid",
        "lookup_id": 5,
        "lookup_profile": "call",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "protocol_header.correlation_id",
        "lookup_id": 1,
        "lookup_profile": "call",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "data_header.callid",
        "lookup_id": 1,
        "lookup_profile": "call",
        "lookup_field": "data_header->>'callid'",
        "lookup_range": [
            -300,
            200
        ],
        "input_function_js": "var returnData=[]; for (var i = 0; i < data.length; i++) { returnData.push(data[i]+'_b2b-1'); }; returnData;"
    }
	]	
`

var CorrelationMapping1registrationold = `[		
    {
        "source_field": "data_header.callid",
        "lookup_id": 100,
        "lookup_profile": "registration",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "data_header.callid",
        "lookup_id": 5,
        "lookup_profile": "registration",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "protocol_header.correlation_id",
        "lookup_id": 1,
        "lookup_profile": "registration",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    }
	]
`

var CorrelationMapping1defaultold = `[
    {
        "source_field": "data_header.callid",
        "lookup_id": 100,
        "lookup_profile": "default",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "data_header.callid",
        "lookup_id": 5,
        "lookup_profile": "default",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "protocol_header.correlation_id",
        "lookup_id": 1,
        "lookup_profile": "default",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    }
	]
`

/* new */

var CorrelationMapping1call = `[		
    {
        "source_field": "data_header.callid",
        "lookup_id": 100,
        "lookup_profile": "call",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "data_header.callid",
        "lookup_id": 5,
        "lookup_profile": "call",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "protocol_header.correlation_id",
        "lookup_id": 1,
        "lookup_profile": "call",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "data_header.callid",
        "lookup_id": 1,
        "lookup_profile": "call",
        "lookup_field": "data_header->>'callid'",
        "lookup_range": [
            -300,
            200
        ],
        "input_function_js": "var returnData=[]; for (var i = 0; i < data.length; i++) { returnData.push(data[i]+'_b2b-1'); }; returnData;"
    }
	]	
`

var CorrelationMapping1registration = `[		
    {
        "source_field": "data_header.callid",
        "lookup_id": 100,
        "lookup_profile": "registration",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "data_header.callid",
        "lookup_id": 5,
        "lookup_profile": "registration",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "protocol_header.correlation_id",
        "lookup_id": 1,
        "lookup_profile": "registration",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    }
	]
`

var CorrelationMapping1default = `[
    {
        "source_field": "data_header.callid",
        "lookup_id": 100,
        "lookup_profile": "default",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "data_header.callid",
        "lookup_id": 5,
        "lookup_profile": "default",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "protocol_header.correlation_id",
        "lookup_id": 1,
        "lookup_profile": "default",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    }
	]
`

var CorrelationMapping60call = `[
    {
        "source_field": "data_header.callid",
        "lookup_id": 100,
        "lookup_profile": "default",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "data_header.callid",
        "lookup_id": 5,
        "lookup_profile": "default",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    },
    {
        "source_field": "protocol_header.correlation_id",
        "lookup_id": 1,
        "lookup_profile": "default",
        "lookup_field": "sid",
        "lookup_range": [
            -300,
            200
        ]
    }
	]
`

var CorrelationMapping60registration = `[
    {
        "source_field": "сallid",
        "lookup_id": 1,
        "lookup_table": "hep_proto_1_registration",
        "lookup_profile": "registration",
        "lookup_field": "sid",
        "partition_field": "record_datetime",
        "timerange_field": "create_ts",
        "lookup_range": [
            -300,
            200
        ]
    }
]`

/* OLD SCHEMA */
var CorrelationMapping60callold = `[
    {
        "source_field": "callid",
        "lookup_id": 1,
        "lookup_table": "sip_messages_call",
        "lookup_profile": "call_h20",
        "lookup_field": "callid",
        "partition_field": "record_datetime",
        "timerange_field": "create_ts",
        "lookup_range": [
            -300,
            200
        ]
    }
]`

var CorrelationMapping60registrationold = `[
    {
        "source_field": "сallid",
        "lookup_id": 1,
        "lookup_table": "sip_messages_registration",
        "lookup_profile": "registration_h20",
        "lookup_field": "callid",
        "partition_field": "record_datetime",
        "timerange_field": "create_ts",
        "lookup_range": [
            -300,
            200
        ]
    }
]`

var CorrelationMapping100default = `[
		{
			"source_field": "sid",
			"lookup_id": 1,
			"lookup_profile": "call",
			"lookup_field": "data_header.callid",
			"lookup_range": [-300, 200]
		},
		{
			"source_field": "sid",
			"lookup_id": 1,
			"lookup_profile": "registration",
			"lookup_field": "data_header.callid",
			"lookup_range": [-300, 200]
		},{
			"source_field": "sid",
			"lookup_id": 1,
			"lookup_profile": "default",
			"lookup_field": "data_header.callid",
			"lookup_range": [-300, 200]
		}
	]	
`

var CorrelationMapping100defaultold = `[
		{
			"source_field": "sid",
			"lookup_id": 1,
			"lookup_profile": "call",
			"lookup_field": "data_header.callid",
			"lookup_range": [-300, 200]
		},
		{
			"source_field": "sid",
			"lookup_id": 1,
			"lookup_profile": "registration",
			"lookup_field": "data_header.callid",
			"lookup_range": [-300, 200]
		},{
			"source_field": "sid",
			"lookup_id": 1,
			"lookup_profile": "default",
			"lookup_field": "data_header.callid",
			"lookup_range": [-300, 200]
		}
	]	
`

var CorrelationMapping34defaultold = `[
		{
			"source_field": "sid",
			"lookup_id": 1,
			"lookup_profile": "call",
			"lookup_field": "data_header.callid",
			"lookup_range": [-300, 200]
		}
	]
`
var CorrelationMapping34default = `[
		{
			"source_field": "sid",
			"lookup_id": 1,
			"lookup_profile": "call",
			"lookup_field": "data_header.callid",
			"lookup_range": [-300, 200]
		}
	]
`

var CorrelationMapping5default = `[
		{
			"source_field": "sid",
			"lookup_id": 1,
			"lookup_profile": "call",
			"lookup_field": "data_header.callid",
			"lookup_range": [-300, 200]
		}
	]
`

var CorrelationMapping1000default = `[
		{
			"source_field": "sid",
			"lookup_id": 1,
			"lookup_profile": "call",
			"lookup_field": "data_header.callid",
			"lookup_range": [-300, 200]
		}
	]
`
