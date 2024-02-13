import React from 'react';
import './navbar.css'
import {Link} from "react-router-dom";

const Navbar = () => {
    return (
            <div className="navbar">
                <div className="navbar__header">GRIDFS CLOUD</div>
                <div className="navbar__common"><Link to="/common" style={{textDecoration: "none"}}> Common files</Link>
                </div>
                <div className="navbar__login"><Link to="/login" style={{textDecoration: "none"}}>Sign in</Link></div>
                <div className="navbar__registration"><Link to="/registration"
                                                            style={{textDecoration: "none"}}>Register</Link></div>
            </div>
    );
};

export default Navbar;