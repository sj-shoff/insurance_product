function addParameter() {
    const container = document.getElementById('parameters-container');
    const template = document.getElementById('parameter-template').content.cloneNode(true);
    container.insertBefore(template, container.lastElementChild);
}

function clearForms() {
    document.querySelectorAll('input, textarea').forEach(field => {
        field.value = '';
    });

    const container = document.getElementById('parameters-container');
    container.querySelectorAll('.grid-cols-2').forEach(row => row.remove());
}

async function saveData() {
    const data = {
        productName: document.querySelector('input[placeholder="Наименование продукта"]').value.trim(),
        startDate: document.querySelectorAll('input[type="date"]')[0]?.value || "",
        endDate: document.querySelectorAll('input[type="date"]')[1]?.value || "",
        updateDate: document.querySelectorAll('input[type="date"]')[2]?.value || "",
        description: document.querySelector('textarea[placeholder="Описание версии"]').value.trim(),
        mandatoryParams: {
            prefixSeries: document.querySelector('input[placeholder="Префикс серии"]').value.trim(),
            postfixSeries: document.querySelector('input[placeholder="Постфикс серии"]').value.trim(),
            prefixNumber: document.querySelector('input[placeholder="Префикс номера"]').value.trim(),
            postfixNumber: document.querySelector('input[placeholder="Постфикс номера"]').value.trim(),
            numerator: document.querySelector('input[placeholder="Нумератор"]').value.trim(),
            customMethod: document.querySelector('input[placeholder="Метод кастомного регулирования номера договора"]').value.trim(),
        },
        costFormula: document.querySelector('input[placeholder="Формула для расчета стоимости"]').value.trim(),
        individualParams: {}
    };

    // Обработка строк с параметрами и значениями
    document.querySelectorAll('#parameters-container .grid-cols-2').forEach(row => {
        const parameter = row.querySelector('input[name="parameter[]"]').value.trim();
        const value = row.querySelector('input[name="value[]"]').value.trim();
        if (parameter && value) {
            data.individualParams[parameter] = value; // Добавляем только заполненные строки
        }
    });

    try {
        const response = await fetch('/save', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data, null, 4) // Форматируем JSON для читаемости
        });
        if (response.ok) {
            alert('Данные успешно сохранены!');
        } else {
            alert('Ошибка при сохранении данных.');
        }
    } catch {
        alert('Ошибка соединения с сервером.');
    }
}
