import {createAuthProvider} from "react-token-auth";
import axios from "axios";
import {Button, Dialog, DialogActions, DialogContent, DialogTitle, Snackbar, TextField} from "@material-ui/core";
import {SetStateAction, useContext, useState} from "react";
import {useNavigate} from 'react-router-dom';
import navbar from "../navbar/navbar";

type Session = { accessToken: string };

export const {useAuth, authFetch, login, logout, getSession} = createAuthProvider<Session>({
    getAccessToken: session => session.accessToken,
    storage: localStorage,
});

const LoginForm = () => {
    const [open, setOpen] = useState(true);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);

    const navigate = useNavigate(); // Инициализация хука useHistory


    const handleClose = () => {
        setOpen(false);
        navigate('/');
    };


    const handleUsernameChange = (e: { target: { value: SetStateAction<string>; }; }) => {
        setUsername(e.target.value);
    };

    const handlePasswordChange = (e: { target: { value: SetStateAction<string>; }; }) => {
        setPassword(e.target.value);
    };


    const handleSubmit = (e: { preventDefault: () => void; }) => {
        e.preventDefault();

// Отправка данных на сервер для аутентификации
        const data = {
            username: username,
            password: password
        };

        axios.post('http://localhost:4000/api/auth/login', data)
            .then((response) => {
                handleClose();
                let token = response.data.data.accessToken
                login(token)
                localStorage.setItem('REACT_TOKEN_AUTH', token)
                console.log(response.data);
                navigate('/', { state: { success: true } })
                setSuccess(true);
            })
            .catch((error) => {
                console.log(error);
                handleClose();
                navigate('/', { state: { error: true } })
                setError(true);
            });
    };

    return (
        <>
            <Dialog open={open} onClose={handleClose}>
                <DialogTitle style={{textAlign: "center"}}>Login</DialogTitle>
                <DialogContent style={{display: 'flex', flexDirection: 'column'}}>
                    <TextField label="Username" value={username} onChange={handleUsernameChange}/>
                    <TextField label="Password" type="password" value={password} onChange={handlePasswordChange}/>
                </DialogContent>
                <DialogActions style={{justifyContent: 'space-between'}}>
                    <Button onClick={handleClose}>Cancel</Button>
                    <Button onClick={handleSubmit} color="primary">Login</Button>
                </DialogActions>
            </Dialog>
        </>
    );
};

export default LoginForm;