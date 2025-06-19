#!/bin/bash

echo "Starting Ollama server in background..."
ollama serve &

echo "Waiting for server to start..."
sleep 10

if ! ollama list >/dev/null 2>&1; then
  echo "Ollama server failed to start!"
  exit 1
fi

echo "Pulling required models..."
ollama pull nomic-embed-text:v1.5 && \
ollama pull gemma3:latest

if ollama list | grep -q "gemma3\|nomic-embed-text"; then
  echo "Models loaded successfully!"
else
  echo "Failed to load models!"
  exit 1
fi

echo "Ollama is ready!"
wait