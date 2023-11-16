package usecases

import (
	"github.com/linkedin/goavro/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertAvroToMessage(t *testing.T) {
	// Define um codec Avro baseado no seu esquema
	codec, err := goavro.NewCodec(`
    {
        "type": "record",
        "name": "Message",
        "fields": [
            {"name": "request_id", "type": "string"},
            {"name": "account_id", "type": "string"},
            {"name": "session_id", "type": "string"},
            {"name": "request", "type": "string"},
            {"name": "response", "type": "string"},
            {"name": "target_url", "type": "string"},
            {"name": "service_name", "type": "string"},
            {"name": "timestamp", "type": "string", "logicalType": "timestamp-millis"}
        ]
    }`)
	assert.NoError(t, err)

	// Cria um mapa representando uma mensagem Avro
	avroDataMap := map[string]interface{}{
		"request_id":   "req-123",
		"account_id":   "acc-456",
		"session_id":   "sess-789",
		"request":      "test request",
		"response":     "test response",
		"target_url":   "http://example.com",
		"service_name": "test service",
		"timestamp":    "2023-11-15T14:30:00.000Z",
	}

	// Converte o mapa para dados binários Avro
	avroData, err := codec.BinaryFromNative(nil, avroDataMap)
	assert.NoError(t, err)

	// Converte a mensagem Avro para o objeto Message
	msg, err := ConvertAvroToMessage(avroData, codec)
	assert.NoError(t, err)
	assert.NotNil(t, msg)

	// Verifica se os campos estão corretos
	assert.Equal(t, "req-123", msg.RequestID)
	assert.Equal(t, "acc-456", msg.AccountID)
	assert.Equal(t, "sess-789", msg.SessionID)
	assert.Equal(t, "test request", msg.Request)
	assert.Equal(t, "test response", msg.Response)
	assert.Equal(t, "http://example.com", msg.TargetURL)
	assert.Equal(t, "test service", msg.ServiceName)
	assert.Equal(t, "2023-11-15T14:30:00.000Z", msg.Timestamp)
}
