package tx

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateEventsMap(t *testing.T) {
	binaryEvents, err := getEventsData([]byte(jsonTx))
	if err != nil {
		t.Fatal(err)
	}
	m := createEventsMap(binaryEvents)
	// for k, v := range m {
	// fmt.Printf("key %s | has values: %s", k, v)
	// }
	fmt.Println(m["message"])
}

func TestParseTx(t *testing.T) {
	tx, err := ParseTx([]byte(jsonTx))
	if err != nil {
		t.Fatal(err)
	}
	bz, err := json.MarshalIndent(tx, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bz))
}

const jsonTx = `{"jsonrpc":"2.0","id":0,"result":{"query":"tm.event = 'Tx'","data":{"type":"tendermint/event/Tx","value":{"TxResult":{"height":"1114","index":0,"tx":"KCgWqQpyXSjKCAoIdHJhbnNmZXISCmliY29uZXhmZXIY2QgiJQocdHJhbnNmZXIvaWJjemVyb3hmZXIvbjB0b2tlbhIFMTAwMDAqFOlsTSvORpXrHiRYAgq6MMq7in4BMhQWdu2a1IfSZVwCNE5DKx6TxsbCkjgBEhMKDQoFc3Rha2USBDUwMDAQwJoMGmoKJuta6YchAxZKRZ4clnaYlUPr2bpKF6ML0OhwYiXvehI741Sy5+T1EkB96oxW8OQsNGIUK6nAcuTlBSbaLktF+c8fAmypu/I+6kpJ6IUpYq/PVCHY0QZLk1aAnFX8NogOS0GTo07iMaYb","result":{"log":"[{\"msg_index\":0,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"transfer\"},{\"key\":\"sender\",\"value\":\"cosmos1a9ky627wg627k83ytqpq4w3se2ac5lspkhz8rx\"},{\"key\":\"sender\",\"value\":\"cosmos1a9ky627wg627k83ytqpq4w3se2ac5lspkhz8rx\"},{\"key\":\"module\",\"value\":\"ibc_transfer\"},{\"key\":\"sender\",\"value\":\"cosmos1a9ky627wg627k83ytqpq4w3se2ac5lspkhz8rx\"},{\"key\":\"receiver\",\"value\":\"cosmos1zemwmxk5slfx2hqzx38yx2c7j0rvds5jelkgdt\"}]},{\"type\":\"send_packet\",\"attributes\":[{\"key\":\"packet_data\",\"value\":\"{\\\"type\\\":\\\"ibc/transfer/PacketDataTransfer\\\",\\\"value\\\":{\\\"amount\\\":[{\\\"amount\\\":\\\"10000\\\",\\\"denom\\\":\\\"transfer/ibczeroxfer/n0token\\\"}],\\\"receiver\\\":\\\"cosmos1zemwmxk5slfx2hqzx38yx2c7j0rvds5jelkgdt\\\",\\\"sender\\\":\\\"cosmos1a9ky627wg627k83ytqpq4w3se2ac5lspkhz8rx\\\",\\\"source\\\":true,\\\"timeout\\\":\\\"2113\\\"}}\"},{\"key\":\"packet_timeout\",\"value\":\"2113\"},{\"key\":\"packet_sequence\",\"value\":\"2\"},{\"key\":\"packet_src_port\",\"value\":\"transfer\"},{\"key\":\"packet_src_channel\",\"value\":\"ibconexfer\"},{\"key\":\"packet_dst_port\",\"value\":\"transfer\"},{\"key\":\"packet_dst_channel\",\"value\":\"ibczeroxfer\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta\"},{\"key\":\"amount\",\"value\":\"5000stake\"},{\"key\":\"recipient\",\"value\":\"cosmos1qmvw4k4rgu53066yjaz03m83uzvexhytsjq4er\"},{\"key\":\"amount\",\"value\":\"10000n0token\"}]}]}]","gas_wanted":"200000","gas_used":"67678","events":[{"type":"message","attributes":[{"key":"YWN0aW9u","value":"dHJhbnNmZXI="}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y29zbW9zMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsc2VycXRh"},{"key":"YW1vdW50","value":"NTAwMHN0YWtl"}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y29zbW9zMWE5a3k2Mjd3ZzYyN2s4M3l0cXBxNHczc2UyYWM1bHNwa2h6OHJ4"}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y29zbW9zMXFtdnc0azRyZ3U1MzA2NnlqYXowM204M3V6dmV4aHl0c2pxNGVy"},{"key":"YW1vdW50","value":"MTAwMDBuMHRva2Vu"}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y29zbW9zMWE5a3k2Mjd3ZzYyN2s4M3l0cXBxNHczc2UyYWM1bHNwa2h6OHJ4"}]},{"type":"send_packet","attributes":[{"key":"cGFja2V0X2RhdGE=","value":"eyJ0eXBlIjoiaWJjL3RyYW5zZmVyL1BhY2tldERhdGFUcmFuc2ZlciIsInZhbHVlIjp7ImFtb3VudCI6W3siYW1vdW50IjoiMTAwMDAiLCJkZW5vbSI6InRyYW5zZmVyL2liY3plcm94ZmVyL24wdG9rZW4ifV0sInJlY2VpdmVyIjoiY29zbW9zMXplbXdteGs1c2xmeDJocXp4Mzh5eDJjN2owcnZkczVqZWxrZ2R0Iiwic2VuZGVyIjoiY29zbW9zMWE5a3k2Mjd3ZzYyN2s4M3l0cXBxNHczc2UyYWM1bHNwa2h6OHJ4Iiwic291cmNlIjp0cnVlLCJ0aW1lb3V0IjoiMjExMyJ9fQ=="},{"key":"cGFja2V0X3RpbWVvdXQ=","value":"MjExMw=="},{"key":"cGFja2V0X3NlcXVlbmNl","value":"Mg=="},{"key":"cGFja2V0X3NyY19wb3J0","value":"dHJhbnNmZXI="},{"key":"cGFja2V0X3NyY19jaGFubmVs","value":"aWJjb25leGZlcg=="},{"key":"cGFja2V0X2RzdF9wb3J0","value":"dHJhbnNmZXI="},{"key":"cGFja2V0X2RzdF9jaGFubmVs","value":"aWJjemVyb3hmZXI="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"aWJjX3RyYW5zZmVy"},{"key":"c2VuZGVy","value":"Y29zbW9zMWE5a3k2Mjd3ZzYyN2s4M3l0cXBxNHczc2UyYWM1bHNwa2h6OHJ4"},{"key":"cmVjZWl2ZXI=","value":"Y29zbW9zMXplbXdteGs1c2xmeDJocXp4Mzh5eDJjN2owcnZkczVqZWxrZ2R0"}]}]}}}},"events":{"send_packet.packet_timeout":["2113"],"send_packet.packet_sequence":["2"],"send_packet.packet_src_channel":["ibconexfer"],"send_packet.packet_dst_port":["transfer"],"transfer.amount":["5000stake","10000n0token"],"message.module":["ibc_transfer"],"message.sender":["cosmos1a9ky627wg627k83ytqpq4w3se2ac5lspkhz8rx","cosmos1a9ky627wg627k83ytqpq4w3se2ac5lspkhz8rx","cosmos1a9ky627wg627k83ytqpq4w3se2ac5lspkhz8rx"],"send_packet.packet_data":["{\"type\":\"ibc/transfer/PacketDataTransfer\",\"value\":{\"amount\":[{\"amount\":\"10000\",\"denom\":\"transfer/ibczeroxfer/n0token\"}],\"receiver\":\"cosmos1zemwmxk5slfx2hqzx38yx2c7j0rvds5jelkgdt\",\"sender\":\"cosmos1a9ky627wg627k83ytqpq4w3se2ac5lspkhz8rx\",\"source\":true,\"timeout\":\"2113\"}}"],"tm.event":["Tx"],"tx.hash":["0E95B9EF7FFAF66DB41BBD66997C4CF06698B342E9576B7F52A16756743247AD"],"message.action":["transfer"],"transfer.recipient":["cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta","cosmos1qmvw4k4rgu53066yjaz03m83uzvexhytsjq4er"],"send_packet.packet_src_port":["transfer"],"send_packet.packet_dst_channel":["ibczeroxfer"],"message.receiver":["cosmos1zemwmxk5slfx2hqzx38yx2c7j0rvds5jelkgdt"],"tx.height":["1114"]}}}
`
