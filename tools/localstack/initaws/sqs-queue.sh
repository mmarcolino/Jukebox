#!/usr/bin/env bash

set -e

export AWS_DEFAULT_REGION=us-east-2

attrs="ReceiveMessageWaitTimeSeconds=10"

# Aguarda o LocalStack estar pronto
until awslocal sqs list-queues > /dev/null 2>&1; do
  echo "Aguardando o LocalStack subir..."
  sleep 2
done

# Cria as filas
awslocal sqs create-queue --queue-name music-queue --attributes "$attrs"

echo "Filas SQS criadas com sucesso!"