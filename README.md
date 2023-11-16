# Requests Tracing PoC

Este é um projeto de prova de conceito (PoC) para demonstrar como consumir mensagens de um tópico Kafka, processá-las e indexá-las em um índice Elasticsearch. O projeto foi desenvolvido em Golang e segue uma arquitetura limpa (Clean Architecture) para manter um código organizado e de fácil manutenção.

## Pré-requisitos

Certifique-se de ter os seguintes pré-requisitos instalados em seu sistema:

- Go: [Instalação do Go](https://golang.org/doc/install)
- Docker: [Instalação do Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Instalação do Docker Compose](https://docs.docker.com/compose/install/)
- Kafka: [Instalação do Apache Kafka](https://kafka.apache.org/quickstart)
- Elasticsearch: [Instalação do Elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/current/install-elasticsearch.html)
- Kibana (opcional): [Instalação do Kibana](https://www.elastic.co/guide/en/kibana/current/install.html)
- Pipenv (para geração de mensagens de teste): [Instalação do Pipenv](https://pipenv.pypa.io/en/latest/)

## Configuração

Antes de executar o projeto, você precisará configurar as variáveis de ambiente:

```bash
# Configurações do Kafka
KAFKA_BOOTSTRAP_SERVERS=localhost:9092

# Configurações do Elasticsearch
ELASTICSEARCH_URL=https://localhost:9200 
# atenção ao https, caso use o compose junto do repositório é necessário que seja https
```

## Execução com Docker Compose

Você pode executar a infraestrutura necessária para o projeto usando o Docker Compose. Certifique-se de que o Docker Compose esteja instalado em seu sistema.

1. Navegue até o diretório raiz do projeto onde está localizado o arquivo `docker-compose.yml`.

2. Execute o seguinte comando para iniciar os serviços do Kafka, Elasticsearch e Kibana:

```bash
docker-compose up -d
```

Isso iniciará os contêineres Docker para Kafka, Elasticsearch e Kibana em segundo plano.

3. Após a inicialização bem-sucedida, você pode executar o projeto Golang com o seguinte comando:

```bash
go run main.go
```

Isso iniciará o consumidor Kafka que processará as mensagens e as indexará no Elasticsearch.

4. Para encerrar os serviços do Docker Compose, execute o seguinte comando:

```bash
docker-compose down
```

Isso encerrará os contêineres Docker.

## Geração de Mensagens de Teste

Para gerar mensagens de teste aleatórias, você pode usar o script Python `publish-message.py`. Certifique-se de ter o Pipenv instalado e execute os seguintes comandos:

```bash
pipenv install
pipenv run python publish-message.py
```

Isso gerará o número descrito no script de mensagens aleatórias e as publicará no tópico Kafka.

## Contribuição

Sinta-se à vontade para contribuir com este projeto. Você pode abrir problemas (issues) ou enviar solicitações de pull (pull requests) para melhorar o código.
