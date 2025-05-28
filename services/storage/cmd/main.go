package main

import (
	"sync"

	"github.com/child6yo/rago/services/storage/internal/app"
	"github.com/child6yo/rago/services/storage/internal/config"
)

func main() {
	// llm, err := ollama.New(ollama.WithModel("qwen3:0.6b"))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// embedder, err := embeddings.NewEmbedder(llm)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// conn := qdrant.NewQdrantConnection(embedder)

	// ctx := context.Background()
	// s, err := conn.Store.AddDocuments(ctx, []schema.Document{
	// 	{PageContent: "Протокол SSL был изначально разработан в компании Netscape для обеспечения безопасности электронной коммерции в Вебе"},
	// 	{PageContent: "Во время TLS-рукопожатия протокол также позволяет обеим сторонам проверить свою идентичность. В браузере этот механизм позволяет клиенту убедиться,"},
	// 	{PageContent: "кошки и собаки - хорошие домашние животные"},
	// 	{PageContent: "женщины - это зло"},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(s)

	// a, _ := conn.Store.SimilaritySearch(ctx, "милота", 2)
	// log.Println(a)
	// a, _ = conn.Store.SimilaritySearch(ctx, "сетевой протокол", 2)
	// log.Println(a)
	// a, _ = conn.Store.SimilaritySearch(ctx, "что такое TLS-рукопожатие", 2)
	// log.Println(a)

	cfg := config.InitConfig()
	app := app.CreateApplication(*cfg)
	app.StartApplication()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
