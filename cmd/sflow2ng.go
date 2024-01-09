package main

type Field struct {
	PEN    int    `json:"PEN"`
	Field  int    `json:"field"`
	Len    int    `json:"len"`
	Format string `json:"format"`
	Name   string `json:"name"`
	Descr  string `json:"descr"`
}

const (
	SFLOW_DEVICE_IP            = 0
	SFLOW_SAMPLES_GENERATED    = 1
	SFLOW_IF_INDEX             = 2
	SFLOW_IF_NAME              = 3
	SFLOW_IF_TYPE              = 4
	SFLOW_IF_SPEED             = 5
	SFLOW_IF_DIRECTION         = 6
	SFLOW_IF_ADMIN_STATUS      = 7
	SFLOW_IF_OPER_STATUS       = 8
	SFLOW_IF_IN_OCTETS         = 9
	SFLOW_IF_IN_PACKETS        = 10
	SFLOW_IF_IN_ERRORS         = 11
	SFLOW_IF_OUT_OCTETS        = 12
	SFLOW_IF_OUT_PACKETS       = 13
	SFLOW_IF_OUT_ERRORS        = 14
	SFLOW_IF_PROMISCUOUS_MODE  = 15
	SFLOW_IF_SAMPLING_INTERVAL = 34
)
const (
	SFLOW_FIELD_IN_BYTES                     = 1
	SFLOW_FIELD_IN_PKTS                      = 2
	SFLOW_FIELD_PROTOCOL                     = 4
	SFLOW_FIELD_L4_SRC_PORT                  = 7
	SFLOW_FIELD_INPUT_SNMP                   = 10
	SFLOW_FIELD_L4_DST_PORT                  = 11
	SFLOW_FIELD_OUTPUT_SNMP                  = 14
	SFLOW_FIELD_LAST_SWITCHED                = 21
	SFLOW_FIELD_FIRST_SWITCHED               = 22
	SFLOW_FIELD_OUT_BYTES                    = 23
	SFLOW_FIELD_OUT_PKTS                     = 24
	SFLOW_FIELD_IPV6_SRC_ADDR                = 27
	SFLOW_FIELD_IPV6_DST_ADDR                = 28
	SFLOW_FIELD_SAMPLING_INTERVAL            = 34
	SFLOW_FIELD_TOTAL_FLOWS_EXP              = 42
	SFLOW_FIELD_SRC_VLAN                     = 58
	SFLOW_FIELD_IN_SRC_MAC                   = 56
	SFLOW_FIELD_OUT_DST_MAC                  = 57
	SFLOW_FIELD_IP_PROTOCOL_VERSION          = 60
	SFLOW_FIELD_DIRECTION                    = 61
	SFLOW_FIELD_EXPORTER_IPV6_ADDRESS        = 131
	SFLOW_FIELD_POST_NAPT_SRC_TRANSPORT_PORT = 227
	SFLOW_FIELD_POST_NAPT_DST_TRANSPORT_PORT = 228
	SFLOW_FIELD_CLIENT_TCP_FLAGS             = 57550
	SFLOW_FIELD_SERVER_TCP_FLAGS             = 57551
	SFLOW_FIELD_NPROBE_IPV4_ADDRESS          = 57943
	SFLOW_FIELD_L7_PROTO                     = 57590
	SFLOW_FIELD_L7_CONFIDENCE                = 58032
	SFLOW_FIELD_NPROBE_INSTANCE_NAME         = 57594
	SFLOW_FIELD_L7_PROTO_RISK                = 57981
	SFLOW_FIELD_L7_RISK_SCORE                = 57999
)

func getTemplateString() string {
	return `[{
        "PEN": 0,
        "field": 1,
        "len": 4,
        "format": "formatted_uint",
        "name": "IN_BYTES",
        "descr": "Incoming flow bytes (src->dst)"
    }, {
        "PEN": 0,
        "field": 2,
        "len": 4,
        "format": "formatted_uint",
        "name": "IN_PKTS",
        "descr": "Incoming flow packets (src->dst)"
    }, {
        "PEN": 0,
        "field": 4,
        "len": 1,
        "format": "uint",
        "name": "PROTOCOL",
        "descr": "IP protocol byte"
    }, {
        "PEN": 0,
        "field": 7,
        "len": 2,
        "format": "uint",
        "name": "L4_SRC_PORT",
        "descr": "IPv4 source port"
    }, {
        "PEN": 0,
        "field": 10,
        "len": 4,
        "format": "uint",
        "name": "INPUT_SNMP",
        "descr": "Input interface SNMP idx"
    }, {
        "PEN": 0,
        "field": 11,
        "len": 2,
        "format": "uint",
        "name": "L4_DST_PORT",
        "descr": "IPv4 destination port"
    }, {
        "PEN": 0,
        "field": 14,
        "len": 4,
        "format": "uint",
        "name": "OUTPUT_SNMP",
        "descr": "Output interface SNMP idx"
    }, {
        "PEN": 0,
        "field": 21,
        "len": 4,
        "format": "epoch",
        "name": "LAST_SWITCHED",
        "descr": "SysUptime (msec) of the last flow pkt"
    }, {
        "PEN": 0,
        "field": 22,
        "len": 4,
        "format": "epoch",
        "name": "FIRST_SWITCHED",
        "descr": "SysUptime (msec) of the first flow pkt"
    }, {
        "PEN": 0,
        "field": 23,
        "len": 4,
        "format": "formatted_uint",
        "name": "OUT_BYTES",
        "descr": "Outgoing flow bytes (dst->src)"
    }, {
        "PEN": 0,
        "field": 24,
        "len": 4,
        "format": "formatted_uint",
        "name": "OUT_PKTS",
        "descr": "Outgoing flow packets (dst->src)"
    }, {
        "PEN": 0,
        "field": 27,
        "len": 16,
        "format": "ipv6_address",
        "name": "IPV6_SRC_ADDR",
        "descr": "IPv6 source address"
    }, {
        "PEN": 0,
        "field": 28,
        "len": 16,
        "format": "ipv6_address",
        "name": "IPV6_DST_ADDR",
        "descr": "IPv6 destination address"
    }, {
        "PEN": 0,
        "field": 34,
        "len": 4,
        "format": "uint",
        "name": "SAMPLING_INTERVAL",
        "descr": "Sampling rate"
    }, {
        "PEN": 0,
        "field": 42,
        "len": 4,
        "format": "formatted_uint",
        "name": "TOTAL_FLOWS_EXP",
        "descr": "Total number of exported flows"
    }, {
        "PEN": 0,
        "field": 58,
        "len": 2,
        "format": "uint",
        "name": "SRC_VLAN",
        "descr": "Source VLAN (inner VLAN in QinQ)"
    }, {
        "PEN": 0,
        "field": 56,
        "len": 6,
        "format": "mac_address",
        "name": "IN_SRC_MAC",
        "descr": "Source MAC Address"
    }, {
        "PEN": 0,
        "field": 57,
        "len": 6,
        "format": "mac_address",
        "name": "OUT_DST_MAC",
        "descr": "Post Destination MAC Address"
    }, {
        "PEN": 0,
        "field": 60,
        "len": 1,
        "format": "uint",
        "name": "IP_PROTOCOL_VERSION",
        "descr": "[4=IPv4][6=IPv6]"
    }, {
        "PEN": 0,
        "field": 61,
        "len": 1,
        "format": "uint",
        "name": "DIRECTION",
        "descr": "Flow direction [0=RX, 1=TX]"
    }, {
        "PEN": 0,
        "field": 131,
        "len": 16,
        "format": "ipv6_address",
        "name": "EXPORTER_IPV6_ADDRESS",
        "descr": "Flow exporter IPv6 Address"
    }, {
        "PEN": 0,
        "field": 227,
        "len": 2,
        "format": "uint",
        "name": "POST_NAPT_SRC_TRANSPORT_PORT",
        "descr": "Post Napt Source Transport Port"
    }, {
        "PEN": 0,
        "field": 228,
        "len": 2,
        "format": "uint",
        "name": "POST_NAPT_DST_TRANSPORT_PORT",
        "descr": "Post Napt Destination Transport Port"
    }, {
        "PEN": 35632,
        "field": 57550,
        "len": 1,
        "format": "uint",
        "name": "CLIENT_TCP_FLAGS",
        "descr": "Cumulative of all client TCP flags"
    }, {
        "PEN": 35632,
        "field": 57551,
        "len": 1,
        "format": "uint",
        "name": "SERVER_TCP_FLAGS",
        "descr": "Cumulative of all server TCP flags"
    }, {
        "PEN": 35632,
        "field": 57943,
        "len": 4,
        "format": "ipv4_address",
        "name": "NPROBE_IPV4_ADDRESS",
        "descr": "IPv4 address of the host were nProbe runs"
    }, {
        "PEN": 35632,
        "field": 57590,
        "len": 2,
        "format": "uint",
        "name": "L7_PROTO",
        "descr": "Layer 7 Protocol (numeric)"
    }, {
        "PEN": 35632,
        "field": 58032,
        "len": 1,
        "format": "uint",
        "name": "L7_CONFIDENCE",
        "descr": "nDPI confidence"
    }, {
        "PEN": 35632,
        "field": 57594,
        "len": 32,
        "format": "ascii",
        "name": "NPROBE_INSTANCE_NAME",
        "descr": "nprobe instance name"
    }, {
        "PEN": 35632,
        "field": 57981,
        "len": 8,
        "format": "uint",
        "name": "L7_PROTO_RISK",
        "descr": "Layer 7 protocol risk (bitmap)"
    }, {
        "PEN": 35632,
        "field": 57999,
        "len": 2,
        "format": "uint",
        "name": "L7_RISK_SCORE",
        "descr": "Layer 7 flow Risk Score"
    }
]`

}
