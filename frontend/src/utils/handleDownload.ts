import {saveAs} from "file-saver";


export const handleDownload = (downloadURL: any, file: { [x: string]: any; filename: string | undefined; }) => {
    const xhr = new XMLHttpRequest();
    xhr.open('GET', downloadURL + file["_id"]);
    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('REACT_TOKEN_AUTH')); // Установка заголовка Authorization
    xhr.responseType = 'blob';

    xhr.onload = function() {
        if (xhr.status === 200) {
            const blob = xhr.response;
            saveAs(blob, file.filename);
        } else {
            console.log('Не удалось скачать файл');
        }
    };

    xhr.send();
};