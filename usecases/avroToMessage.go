package usecases

import (
	"fmt"
	"github.com/linkedin/goavro/v2"
	"github.com/rafaeldblima/requests-tracing-poc/domain"
)

// ConvertAvroToMessage converte abro para message do sistema
func ConvertAvroToMessage(avroData []byte, codec *goavro.Codec) (*domain.Message, error) {
	native, _, err := codec.NativeFromBinary(avroData)
	if err != nil {
		return nil, err
	}

	// Converte o mapa nativo para a estrutura de mensagem do domínio
	nativeMap, ok := native.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("erro na conversão de avro para mapa nativo")
	}

	message := &domain.Message{
		RequestID:   nativeMap["request_id"].(string),
		AccountID:   nativeMap["account_id"].(string),
		SessionID:   nativeMap["session_id"].(string),
		Request:     nativeMap["request"].(string),
		Response:    nativeMap["response"].(string),
		TargetURL:   nativeMap["target_url"].(string),
		ServiceName: nativeMap["service_name"].(string),
		Timestamp:   nativeMap["timestamp"].(string),
	}

	return message, nil
}
