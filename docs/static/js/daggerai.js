(function() {
  if (typeof window === 'undefined') return;
  if (typeof window.difyChatbot !== 'undefined') return;

  // Create configuration
  window.difyChatbotConfig = {
    token: 'OqnZo1bkX29jclQb',
    baseUrl: 'http://dify-2025-01-16.fly.dev'
  };

  // Create and append script
  var script = document.createElement('script');
  script.src = 'http://dify-2025-01-16.fly.dev/embed.min.js';
  script.id = 'OqnZo1bkX29jclQb';
  script.defer = true;
  document.head.appendChild(script);

  // Create and append styles
  var style = document.createElement('style');
  style.textContent = `
    #dify-chatbot-bubble-button {
      background-color: #1C64F2 !important;
    }
    #dify-chatbot-bubble-window {
      width: 56rem !important;
      height: 40rem !important;
    }
  `;
  document.head.appendChild(style);
})();
