import datetime
import random
import uuid
from io import BytesIO

import fastavro
from confluent_kafka import Producer

# Configurações do Kafka
kafka_config = {
    'bootstrap.servers': 'localhost:9092',  # Substitua pelo endereço do seu servidor Kafka
    'client.id': 'avro-producer'
}

# Esquema Avro (substitua pelo seu próprio esquema)
avro_schema = {
    'type': 'record',
    'name': 'Message',
    'fields': [
        {'name': 'request_id', 'type': 'string'},
        {'name': 'account_id', 'type': 'string'},
        {'name': 'session_id', 'type': 'string'},
        {'name': 'request', 'type': 'string'},
        {'name': 'response', 'type': 'string'},
        {'name': 'target_url', 'type': 'string'},
        {'name': 'service_name', 'type': 'string'},
        {'name': 'timestamp', "type": "string", "logicalType": "timestamp-millis"}
    ]
}


def generate_random_message():
    # Gere dados de exemplo para a mensagem
    request_id = str(uuid.uuid4())
    account_id = str(uuid.uuid4())
    session_id = str(uuid.uuid4())
    request = f"Request {random.randint(1, 10)}"
    response = f"Response {random.randint(1, 10)}"
    target_url = f"https://example.com/{random.randint(1, 5)}"
    service_name = "example_service"
    agora = datetime.datetime.now()
    timestamp = agora.strftime('%Y-%m-%dT%H:%M:%S.%f')[:-3] + 'Z'

    message_data = {
        'request_id': request_id,
        'account_id': account_id,
        'session_id': session_id,
        'request': request,
        'response': response,
        'target_url': target_url,
        'service_name': service_name,
        'timestamp': timestamp
    }

    return message_data


def produce_avro_messages(topic_name, message_qty):
    producer = Producer(kafka_config)

    for _ in range(message_qty):
        bytes_writer = BytesIO()
        message_data = generate_random_message()
        fastavro.schemaless_writer(bytes_writer, avro_schema, message_data)

        producer.produce(topic_name, key=message_data['request_id'], value=bytes_writer.getvalue())

    producer.flush()


if __name__ == "__main__":
    topic = 'tracing-user'  # Substitua pelo nome do seu tópico Kafka
    num_messages = 100  # Defina o número de mensagens a serem geradas

    produce_avro_messages(topic, num_messages)
