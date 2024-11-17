const express = require('express');
const fs = require('fs');
const path = require('path');

const app = express();
const PORT = 3000;

// Путь к JSON-файлу
const DATA_FILE = path.join(__dirname, 'data.json');

// Middleware для обработки JSON
app.use(express.json());

// Обслуживание статических файлов
app.use(express.static(path.join(__dirname, 'public')));

// Маршрут для сохранения данных
app.post('/save', (req, res) => {
    const data = req.body;
    fs.writeFile(DATA_FILE, JSON.stringify(data, null, 4), (err) => {
        if (err) {
            console.error(err);
            return res.status(500).json({ message: 'Ошибка при сохранении данных' });
        }
        res.status(200).json({ message: 'Данные успешно сохранены' });
    });
});

// Запуск сервера
app.listen(PORT, () => {
    console.log(`Сервер запущен: http://localhost:${PORT}`);
});
