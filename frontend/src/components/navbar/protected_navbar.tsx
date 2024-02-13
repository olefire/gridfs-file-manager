import React from 'react';
import './navbar.css'
import {Link} from "react-router-dom";
import Button from "@material-ui/core/Button";
import {logout} from "../login/login";
import {useNavigate} from 'react-router-dom';

const ProtectedNavbar = () => {
    const currentRoute = window.location.pathname;
    const navigate = useNavigate(); // Инициализация хука useHistory
    const handleLogout = () => {
        logout()
        localStorage.clear()
        navigate('/')
    };

    return (
        <>
            <div className="navbar">
                <div className="navbar__header">GRIDFS CLOUD</div>
                <div className="navbar__logout" onClick={handleLogout}>Logout</div>
            </div>
            <div className="buttons" style={{display:"flex", justifyContent: "space-around"}}>
                <Button>
                    <Link to="/common" className={currentRoute === "/common" ? "common__active" : "common"}> Common files</Link>
                </Button>
                <Button>
                    <Link to="/my" className={currentRoute === "/my" ? "my__active" : "my"}> My files</Link>
                </Button>

                <Button>
                    <Link to="/shared" className={currentRoute === "/shared" ? "shared__active" : "shared"}> Shared files</Link>
                </Button>

                <Button>
                    <Link to="/upload" className={currentRoute === "/upload" ? "upload__active" : "upload"}> Upload file</Link>
                </Button>
            </div>
        </>
    );
};


export default ProtectedNavbar;