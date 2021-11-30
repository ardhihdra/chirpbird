// import logo from '../assets/img/logo/logo-horizontal.png';
// import LoginForm from './LoginForm';
// import ContactInfo from '../components/ContactInfo';
import Input from '../components/Input';
import React, { useState } from 'react';
import { useParams, useNavigate } from "react-router-dom";
import menuImg from '../assets/img/icons/menu.png';
import menuDotImg from '../assets/img/icons/menu-dot.png';
import userImg from '../assets/img/icons/010-user.png';

import './Dashboard.css';

export default function Home() {
    const [isProfileModalShow, setProfileModalShow] = useState(0)
    const { key } = useParams();

    return (
        <div>
            <div className="header-menu">
                <div className="ds-pt-1 ds-mr-5">
                    <img src={menuImg} height="30px"/>
                </div>
                <div className="ds-mr-5 search-input">
                    <Input placeholder="Search"></Input>
                </div>
                <div className="ds-ml-1 ds-pt-1 room-pict">
                    <img src={userImg} height="30px"/>
                </div>
                <div className="ds-ml-5 ds-mt-1 room-title">
                    <div>room title</div>
                    <div className="txt-desc-meta">room desc</div>
                </div>
                <div className="ds-mr-5 room-search">
                    <Input className="transparent night-mode" placeholder="Search in chat"></Input>
                </div>
                <div className="ds-pt-4 ds-mr-5 ds-pr-4">
                    <img src={menuDotImg} height="20px"/>
                </div>
            </div>
            <div className="dashboard-container">
                <div className="left-menu">
                    {[1,2,3].map(el => {
                        return (
                            <div className="room-preview-box ds-flex">
                                <div className="ds-ml-1 ds-mt-1 room-pict">
                                    <img src={userImg} height="40px"/>
                                </div>
                                <div className="ds-ml-5 ds-mt-1 room-title">
                                    <div>room title {el}</div>
                                    <div className="ds-flex">
                                        <div className="txt-desc-sm">room desc :</div>
                                        <div className="txt-desc-meta ds-ml-2">room chatting preview</div>
                                    </div>
                                </div>
                            </div>
                        )
                    })}
                </div>
                <div className="chatbox">chatbox</div>
            </div>
            {/* <div className="modal">modal</div> */}
        </div>
    )
}