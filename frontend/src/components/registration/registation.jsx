import React, { useState } from 'react';
import { TextField, Button, Dialog, DialogTitle, DialogContent, DialogActions, Snackbar } from '@material-ui/core';
import axios from 'axios';
import {useNavigate} from "react-router-dom";

const RegistrationForm = () => {
    const [open, setOpen] = useState(true);
    const [name, setName] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const navigate = useNavigate(); // Инициализация хука useHistory
    const handleClose = () => {
        setOpen(false);
        navigate("");
    };

    const handleNameChange = (e) => {
        setName(e.target.value);
    };

    const handleUsernameChange = (e) => {
        setUsername(e.target.value);
    };

    const handlePasswordChange = (e) => {
        setPassword(e.target.value);
    };

    const handleEmailChange = (e) => {
        setEmail(e.target.value);
    };


    const handleSubmit = (e) => {
        e.preventDefault();

// Отправка данных на сервер для регистрации
        const data = {
            name: name,
            username: username,
            password: password,
            email: email
        };
        const headers = {
            'Content-Type': 'application/json'
        };

        axios.post('http://localhost:4000/api/auth/signup', data, {headers})
            .then((response) => {
                setSuccess(true);
                console.log(response.data)
                handleClose();
                navigate("");
            })
            .catch((error) => {
                console.log(error)
                setError(true);
                handleClose();
                navigate("");
            });
    };

    return (
        <>
            <Dialog open={open} onClose={handleClose}>
                <DialogTitle style={{textAlign: "center"}}>Registration</DialogTitle>
                <DialogContent style={{ display: 'flex', flexDirection: 'column' }}>
                    <TextField label="Name" value={name} onChange={handleNameChange} />
                    <TextField label="Username" value={username} onChange={handleUsernameChange} />
                    <TextField label="Password" type="password" value={password} onChange={handlePasswordChange} />
                    <TextField label="Email" type="email" value={email} onChange={handleEmailChange} />
                </DialogContent>
                <DialogActions style={{ justifyContent: 'space-between' }}>
                    <Button onClick={handleClose}>Cancel</Button>
                    <Button onClick={handleSubmit} color="primary">Register</Button>
                </DialogActions>
            </Dialog>
            <Snackbar open={success} onClose={() => setSuccess(false)} message="Registration successful!" />
            <Snackbar open={error} onClose={() => setError(false)} message="Registration failed. Please try again." />
        </>
    );
};

export default RegistrationForm;