import React, {useState, useEffect} from 'react';
import axios from 'axios';
import {defaultStyles, FileIcon} from "react-file-icon";
import Button from "@material-ui/core/Button";
import {ext} from "../../utils/extension";
import {handleDownload} from "../../utils/handleDownload";

const downloadUrl = "http://localhost:4000/api/media/"
const CommonFiles = () => {
    const [files, setFiles] = useState([]);

    useEffect(() => {
// Получение данных с эндпоинта
        axios.get('http://localhost:4000/api/file/public')
            .then((response) => {
// Установка полученных данных в состояние компонента
                setFiles(response.data.data.files);
                console.log(response.data)
            })
            .catch((error) => {
                console.log(error)
            });
    }, []);

    return (
        <>
            <div style={{display: "flex", flexWrap: "wrap"}}>
                {/* Отображение данных в массиве files */}
                {files?.map((file) => (
                    <div style={{margin: "75px"}}>
                        <div style={{width: '100px', height: '100px', position: "relative"}}><FileIcon
                            extension={ext(file.filename)} {...defaultStyles.bmp} />
                            <div style={{position: "absolute", left: "50%", transform: "translateX(-50%)"}}><Button
                                variant="contained" color="primary"
                                onClick={() => handleDownload(downloadUrl, file)}>
                                {file.filename}
                            </Button>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </>
    );
};

export default CommonFiles;