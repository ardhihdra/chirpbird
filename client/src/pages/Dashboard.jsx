// import logo from '../assets/img/logo/logo-horizontal.png';
// import LoginForm from './LoginForm';
// import ContactInfo from '../components/ContactInfo';
import React, { useState } from 'react';
import { useParams, useNavigate } from "react-router-dom";

import './Dashboard.css';

export default function Home() {
    const [isProfileModalShow, setProfileModalShow] = useState(0)
    const { key } = useParams();

    return (
        <div>
            <div className="header-menu tes1">header</div>
            <div className="dashboard-container">
                <div className="left-menu tes2">leftmenu</div>
                <div className="chatbox">chatbox</div>
            </div>
            {/* <div className="modal">modal</div> */}
        </div>
    )
}