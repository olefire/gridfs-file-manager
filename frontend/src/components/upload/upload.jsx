import React, {useState} from 'react';
import Button from '@material-ui/core/Button';
import axios from 'axios';
import {Checkbox, FormControlLabel, Snackbar} from "@material-ui/core";

export const UploadFile = () => {
    const [success, setSuccess] = useState(false)
    const [error, setError] = useState(false)
    const [isPublic, setIsPublic] = useState(false);
    const [file, setFile] = useState();
    const [url, setUrl] = useState('http://localhost:4000/api/media/protected/?isPublic=false')
    let config = {
        headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('REACT_TOKEN_AUTH'),
            'Content-Type': 'multipart/form-data'
        }
    }
    const handleCheck = (event) => {
        setIsPublic(event.target.checked);
        if (event.target.checked === true) {
            setUrl('http://localhost:4000/api/media/protected/?isPublic=true')
        }
        else {
            setUrl('http://localhost:4000/api/media/protected/?isPublic=false')
        }
    };

    const handleFileChange = (e) => {
        if (e.target.files) {
            setFile(e.target.files[0]);
        }
    };

    const handleUpload = () => {
        const formData = new FormData();
        formData.append('file', file);

        axios.post(url, formData, config)
            .then((response) => {
                console.log(response)
                setSuccess(true)
            })
            .catch((error) => {
                console.log(formData)
                console.log(error)
                setError(true)
            });
    };

    return (
        <div style={{textAlign: "center", marginTop: "30px"}}>
            <input autoFocus={true} type="file" onChange={(e) => handleFileChange(e)}/>
            <FormControlLabel
                control={
            <Checkbox checked={isPublic} onChange={(e) => handleCheck(e)}/>
                }
                label="send public file"
            />
            <div className="btn" style={{marginTop: "10px"}}><Button variant="contained" color="primary" onClick={handleUpload} style={{fontSize: "24px"}}>Upload File</Button></div>
            <Snackbar open={success} onClose={() => setSuccess(false)} message="File successfully uploaded!"/>
            <Snackbar open={error} onClose={() => setSuccess(false)} message="Some error while uploading file"/>
        </div>
    );
};