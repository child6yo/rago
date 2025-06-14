document.addEventListener('DOMContentLoaded', () => {
  // DOM элементы
  const form = document.getElementById('chatForm');
  const input = document.getElementById('userInput');
  const messages = document.getElementById('chatMessages');
  const sendButton = document.getElementById('sendButton');
  const statusIndicator = document.getElementById('statusIndicator');
  
  // Текущее состояние
  let isGenerating = false;
  let eventSource = null;
  let botMessage = null;
  let accumulatedText = '';
  
  // Отправка сообщения
  form.addEventListener('submit', async (e) => {
    e.preventDefault();
    const query = input.value.trim();
    if (!query || isGenerating) return;
    
    // Добавляем сообщение пользователя
    addMessage(query, 'user');
    input.value = '';
    sendButton.disabled = true;
    
    // Создаем контейнер для сообщения бота
    botMessage = createBotMessage();
    messages.appendChild(botMessage);
    messages.scrollTop = messages.scrollHeight;
    
    // Статус подключения
    updateStatus('Подключаемся к серверу...');
    isGenerating = true;
    accumulatedText = '';
    
    try {
      // Создаем EventSource для SSE соединения
      eventSource = new EventSource(`http://localhost:8000/api/v1/generation?query=${encodeURIComponent(query)}`);
      
      // Обработка входящих сообщений
      eventSource.addEventListener('message', (event) => {
        updateStatus('Получаем ответ...');
        
        // Добавляем новый чанк к накопленному тексту
        accumulatedText += event.data;
        
        // Обновляем содержимое сообщения
        const contentElement = botMessage.querySelector('.bot-message-content');
        contentElement.innerHTML = accumulatedText;
        
        // Прокрутка вниз
        messages.scrollTop = messages.scrollHeight;
      });
      
      // Обработка события завершения
      eventSource.addEventListener('end', (event) => {
        updateStatus('Готово');
        if (event.data === '[DONE]') {
          completeGeneration();
        }
      });
      
      // Обработка ошибок
      eventSource.addEventListener('error', (event) => {
        console.error('SSE error:', event);
        updateStatus('Ошибка соединения');
        addErrorIndicator();
      });
      
      eventSource.onerror = () => {
        updateStatus('Соединение прервано');
        addErrorIndicator();
        completeGeneration();
      };
      
    } catch (error) {
      console.error('Connection error:', error);
      updateStatus('Ошибка подключения');
      addErrorIndicator();
      completeGeneration();
    }
  });
  
  // Создание сообщения бота
  function createBotMessage() {
    const message = document.createElement('div');
    message.className = 'message bot';
    
    // Аватар
    const avatar = document.createElement('div');
    avatar.className = 'message-avatar';
    avatar.innerHTML = `
      <svg width="14" height="14" viewBox="0 0 24 24">
        <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm-1-13.5v7l6-3.5z"/>
      </svg>
    `;
    
    // Контент
    const content = document.createElement('div');
    content.className = 'message-content';
    
    const botContent = document.createElement('div');
    botContent.className = 'bot-message-content';
    botContent.textContent = ''; // Начальный пустой текст
    
    // Добавляем индикатор набора текста
    const typingIndicator = document.createElement('div');
    typingIndicator.className = 'typing-indicator';
    typingIndicator.innerHTML = `
      <div class="typing-dot"></div>
      <div class="typing-dot"></div>
      <div class="typing-dot"></div>
    `;
    
    botContent.appendChild(typingIndicator);
    content.appendChild(botContent);
    
    message.appendChild(avatar);
    message.appendChild(content);
    
    return message;
  }
  
  // Завершение генерации
  function completeGeneration() {
    isGenerating = false;
    sendButton.disabled = false;
    
    if (eventSource) {
      eventSource.close();
      eventSource = null;
    }
    
    // Удаляем индикатор набора текста
    const typingIndicator = botMessage.querySelector('.typing-indicator');
    if (typingIndicator) {
      typingIndicator.remove();
    }
    
    botMessage = null;
  }
  
  // Добавление индикатора ошибки
  function addErrorIndicator() {
    if (!botMessage) return;
    
    const errorElement = document.createElement('div');
    errorElement.className = 'error-indicator';
    errorElement.textContent = ' (соединение прервано)';
    errorElement.style.color = 'var(--error)';
    errorElement.style.display = 'inline';
    errorElement.style.fontSize = '0.9em';
    errorElement.style.marginLeft = '0.5em';
    
    const content = botMessage.querySelector('.bot-message-content');
    if (content) {
      content.appendChild(errorElement);
    }
  }
  
  // Обновление статуса
  function updateStatus(text) {
    statusIndicator.textContent = text;
  }
  
  // Добавление сообщения
  function addMessage(text, sender = 'user') {
    const message = document.createElement('div');
    message.className = `message ${sender}`;
    
    // Аватар
    const avatar = document.createElement('div');
    avatar.className = 'message-avatar';
    if (sender === 'user') {
      avatar.innerHTML = `
        <svg width="14" height="14" viewBox="0 0 24 24">
          <path fill="currentColor" d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
        </svg>
      `;
    } else {
      avatar.innerHTML = `
        <svg width="14" height="14" viewBox="0 0 24 24">
          <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm-1-13.5v7l6-3.5z"/>
        </svg>
      `;
    }
    
    // Контент
    const content = document.createElement('div');
    content.className = 'message-content';
    content.textContent = text;
    
    message.appendChild(avatar);
    message.appendChild(content);
    messages.appendChild(message);
    
    // Прокрутка вниз
    messages.scrollTop = messages.scrollHeight;
    
    return message;
  }
  
  // Фокус на поле ввода при загрузке
  input.focus();
});