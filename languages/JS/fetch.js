// Обработчик формы загрузки изображения при создании и редактировании поста
window.addEventListener('DOMContentLoaded', () => {

    let form = document.querySelector('#uploadForm');

    function addItem(item_url, item_title) {
        let row = document.createElement('div');
        row.innerHTML = `<a href="${item_url}" target="_blank">${item_title}</a>`;
        document.querySelector('#uploadedFiles').appendChild(row);
    };

    async function send_request(url, data={}, headers={}) {
        const resp = await fetch(url, {
            method: 'POST',
            headers: headers,
            body: data,
        });

        if (!resp.ok){
            throw new Error(`Could not fetch ${url}, status: ${resp.status}`);
        };

        return await resp.text();
    };

    function upload(e) {
        e.preventDefault();  // сбрасывает действие по умолчанию - перезагрузку страницы

        let filename = form.inputGroupFile02.files[0].name;
        let formData = new FormData(form);

        send_request('/upload', data=formData)
            .then(url => addItem(url, filename))
            .catch(err => console.error('No response from server'));
    };

    form.addEventListener('submit', (e) => upload(e));
});
