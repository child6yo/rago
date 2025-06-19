#!/bin/bash

KAFKA_TOPIC_CREATE_CMD="kafka-topics.sh --bootstrap-server kafka:9092 --create"

echo "Creating topics..."

$KAFKA_TOPIC_CREATE_CMD --topic raw-docs \
  --partitions 5 \
  --config retention.ms=604800000 \
  --config cleanup.policy=delete \
  --config min.insync.replicas=1

$KAFKA_TOPIC_CREATE_CMD --topic document-topic \
  --partitions 5 \
  --config retention.ms=604800000 \
  --config cleanup.policy=delete \
  --config min.insync.replicas=1

echo "Topics created successfully."