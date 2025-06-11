document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("chatForm");
  const input = document.getElementById("userInput");
  const messages = document.getElementById("chatMessages");

  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const query = input.value.trim();
    if (!query) return;

    // Добавляем сообщение пользователя
    addMessage(query, "user");
    input.value = "";

    // Создаем EventSource
    const eventSource = new EventSource(`http://localhost:8000/api/v1/generation?query=${encodeURIComponent(query)}`);
    const botMessage = addMessage("▌", "bot"); // Индикатор генерации

    eventSource.onmessage = (event) => {
      if (event.data === "[DONE]") {
        eventSource.close();
        // Убираем индикатор
        if (botMessage.textContent.endsWith("▌")) {
          botMessage.textContent = botMessage.textContent.slice(0, -1);
        }
      } else {
        // Заменяем индикатор на данные
        if (botMessage.textContent.endsWith("▌")) {
          botMessage.textContent = event.data;
        } else {
          botMessage.textContent += event.data;
        }
        messages.scrollTop = messages.scrollHeight;
      }
    };

    eventSource.onerror = () => {
      // Не выводим ошибку, если это нормальное завершение
      if (botMessage.textContent && !botMessage.textContent.includes("[DONE]")) {
        console.error("SSE connection error");
        if (botMessage.textContent.endsWith("▌")) {
          botMessage.textContent = botMessage.textContent.slice(0, -1);
        }
        botMessage.textContent += "\n⚠ Соединение прервано";
      }
      eventSource.close();
    };
  });
});

function addMessage(text, sender = "user") {
  const message = document.createElement("div");
  message.className = `message ${sender}`;
  message.textContent = text;
  document.getElementById("chatMessages").appendChild(message);
  return message;
}