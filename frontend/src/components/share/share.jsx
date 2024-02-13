import React, {useEffect, useState} from 'react';
import Modal from 'react-modal';
import {defaultStyles, FileIcon} from "react-file-icon";
import Button from '@material-ui/core/Button';
import axios from "axios";
import {ext} from "../../utils/extension";
import {useNavigate} from "react-router-dom";
import "./share.css"
import {Snackbar, TextField} from "@material-ui/core";

export const ShareFiles = () => {
    const [files, setFiles] = useState([]);
    const navigate = useNavigate();
    const [usernames, setUsernames] = useState([]);
    const [showModal, setShowModal] = useState(true);
    const [share, setShare] = useState(Array.from({length: files.length}, () => false));
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);

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
    const toggleModal = () => {
        navigate("/");
        setShowModal(!showModal);

    };

    const handleClick = (index) => {
        const newColors = [...share];
        newColors[index] = !newColors[index];
        setShare(newColors);
    };

    const addUsername = () => {
        const newUsers = [...usernames]
        newUsers.push("")
        setUsernames(newUsers)
    }

    const removeLastUsername = () => {
        const newUsers = [...usernames]
        newUsers.pop()
        setUsernames(newUsers)
    }
    const handleUsernameChange = (username, index) => {
        const newUsers = [...usernames]
        newUsers[index] = username
        setUsernames(newUsers)
        console.log(usernames)
        console.log(username)
    }

    const shareFiles = () => {
        const shareFiles = [];
        for (let i = 0; i < share?.length; i++) {
            if (share[i] === true) {
                shareFiles.push(files[i]["_id"]);
            }
        }

        const request = {
            "users": usernames,
            "files": shareFiles
        };

        if (usernames.length === 0 || shareFiles.length === 0) {
            setError(true)
            return
        }

        axios.patch('http://localhost:4000/api/file/share', request, config)
            .then((response) => {
                console.log(response)
                setSuccess(true)
                navigate("")
            })
            .catch((error) => {
                console.log(error)
                setError(true)
                navigate("")
            });
    }

    return (
        <div>
            <Modal isOpen={showModal} onRequestClose={toggleModal}>
                <h1 style={{textAlign: "center"}}>Share the files</h1>
                <div style={{display: "flex", flexWrap: "wrap"}}>
                    {files?.map((file, index) => (
                        <div style={{margin: "75px"}}>
                            <div style={{width: '100px', height: '100px', position: "relative"}}><FileIcon
                                extension={ext(file.filename)} {...defaultStyles.bmp} />
                                <div style={{position: "absolute", left: "50%", transform: "translateX(-50%)"}}>
                                    <Button
                                        variant="contained"
                                        className={share[index] ? "share" : ""}
                                        color={share[index] ? "secondary" : "primary"}
                                        onClick={() => handleClick(index)}>
                                        {file.filename}
                                    </Button>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
                <Button onClick={toggleModal}></Button>
                <div style={{marginTop: "50px"}}>
                    <Button onClick={addUsername}>add another username</Button>
                    {usernames.map((username, index) => (
                        <div className="username">
                            <TextField
                                onChange={(e) => handleUsernameChange(e.target.value, index)}
                                name="username"
                                label="Username"
                            />
                        </div>
                    ))}
                    <Button onClick={removeLastUsername}>remove last username</Button>
                </div>
                <div style={{textAlign: "center"}}><Button onClick={shareFiles} variant={"contained"}>share the
                    files</Button></div>

                <Snackbar open={success} onClose={() => setSuccess(false)} message="Successfully share files with users!"/>
                <Snackbar open={error} onClose={() => setSuccess(false)} message="Some error while sharing files"/>
            </Modal>
        </div>

    );
};

