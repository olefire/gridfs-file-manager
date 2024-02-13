import React, {useState, useEffect} from 'react';
import {defaultStyles, FileIcon} from "react-file-icon";
import Button from "@material-ui/core/Button";
import axios from "axios";
import {handleDownload} from "../../utils/handleDownload";
import {ext} from "../../utils/extension";
import {Link} from "react-router-dom";

const downloadUrl = "http://localhost:4000/api/media/protected/"

const MyFiles = () => {
    const [files, setFiles] = useState([]);

    let config = {
        headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('REACT_TOKEN_AUTH')
        }
    }

    useEffect(() => {
        axios.get('http://localhost:4000/api/file/private', config)
            .then((response) => {
                setFiles(response.data.data.files);
                console.log(response)
            })
            .catch((error) => {
                console.log(error)
            });
    }, [config]);


    return (
        <>
            <div style={{textAlign: "center"}}>
                <Button><Link to="/share" style={{fontSize: "24px", textDecoration: "none"}}> SHARE </Link></Button>
                <Button><Link to="/delete" style={{fontSize: "24px", textDecoration: "none"}}> DELETE </Link></Button>
            </div>
            <div style={{display: "flex", flexWrap: "wrap"}}>
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

export default MyFiles;