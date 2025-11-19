🎯 Что где находится (понятно и по ролям)
Папка	Что хранится
/cmd/web	Запуск backend HTTP/API сервера
/cmd/wasm	Точка сборки frontend (Go→WASM)
/internal/models	User, Player, GameObject, Session
/internal/repository	Работа с SQLite/PostgreSQL
/internal/service	Бизнес-логика (регистрация, авторизация, игровой сервис)
/internal/handler	HTTP API (login, signup, game data, WebSocket)
/internal/game	Игровая логика (Go + WebAssembly)
/internal/wasm	Код, компилируемый в .wasm
/web	HTML, JS (минимум), CSS, wasm_exec.js, game.wasm