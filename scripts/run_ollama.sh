#!/bin/bash

ollama pull qwen3:0.6b
ollama pull nomic-embed-text:v1.5

echo "Starting Ollama server..."
ollama serve & 
sleep 10

echo "Ollama is ready, creating the model..."

ollama run qwen3:0.6b