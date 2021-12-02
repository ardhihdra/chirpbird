// import logo from '../assets/img/logo/logo-horizontal.png';
// import LoginForm from './LoginForm';
// import ContactInfo from '../components/ContactInfo';
import Input from '../components/Input';
import React, { useState } from 'react';
import { useParams, useNavigate } from "react-router-dom";
import menuImg from '../assets/img/icons/menu.png';
import menuDotImg from '../assets/img/icons/menu-dot.png';
import userImg from '../assets/img/icons/010-user.png';
import attachImg from '../assets/img/icons/attachment.png';
import submitImg from '../assets/img/icons/submit.png';

import './Dashboard.css';

export default function Home() {
    const [isProfileModalShow, setProfileModalShow] = useState(0)
    const [menu, setMenu] = useState(false)
    const { id } = useParams();
    const token = sessionStorage.getItem('token');
    const userinfo = localStorage.getItem('userinfo');
    const activeRooms = [2,2,3]
    
    const openMenu = (e) => {
        setMenu(!menu)
    }


    const menuRender = (
        <div className="menu ds-p-5">
            <div className="ds-m-3">Create New Room</div>
            <Input placeholder="Room name"/>
            <Input placeholder="Add room member (username)"/>
        </div>
    )

    const activeRoomRender = (
        activeRooms.map(el => {
            return (
                <div className="room-preview-box ds-flex">
                    <div className="ds-ml-1 ds-mt-1 ds-mb-1 room-pict">
                        <img alt="room-pict" src={userImg} height="40"/>
                    </div>
                    <div className="ds-ml-5 ds-mt-1 room-title">
                        <div>room title {el}</div>
                        <div className="ds-flex">
                            <div className="txt-desc-sm">room desc :</div>
                            <div className="txt-desc-meta ds-ml-2">room chatting preview</div>
                        </div>
                    </div>
                    <div>
                        <div className="txt-desc-meta-sm">12/1/2021</div>
                        <div className="room-unread ds-mt-2 ds-ml-3">{el+10}</div>
                    </div>
                </div>
            )
        })
    )
    return (
        <div>
            <div className="header-menu">
                <div className="ds-mr-5 icon-bg">
                    <img alt="profile-menu" src={menuImg} height="28" onClick={openMenu}/>
                </div>
                <div className="ds-mr-5 search-input">
                    <Input placeholder="Search"></Input>
                </div>
                <div className="ds-ml-1 ds-mt-1 ds-mb-1 room-pict">
                    <img alt="profile-pict" src={userImg} height="30"/>
                </div>
                <div className="ds-ml-5 ds-mt-1 room-title">
                    <div>room title</div>
                    <div className="txt-desc-meta">room desc</div>
                </div>
                <div className="ds-mr-5 room-search">
                    <Input className="transparent night-mode" placeholder="Search in chat"></Input>
                </div>
                <div className="ds-mt-1 ds-mr-5 icon-bg">
                    <img alt="rooms-menu" src={menuDotImg} height="22px"/>
                </div>
            </div>
            <div className="dashboard-container ds-slide-in">
                <div className="left-container ds-fade-in">
                    {menu ? (menuRender) : (activeRoomRender)}
                </div>
                <div className="chat-container ds-fade-in">
                    <div className="chatbox ds-flex">
                        <div className="ds-ml-5 ds-mt-1 ds-mb-1 room-pict">
                            <img alt="friends-pict" src={userImg} height="40"/>
                        </div>
                        <div className="ds-ml-5 ds-mt-1">
                            <div className="chat-username">friends name</div>
                            <div className="ds-flex chat-content">
                                <div className="txt-desc-sm">room chat text {userinfo}</div>
                            </div>
                        </div>
                        <div>
                            <div className="txt-desc-meta-sm ds-mt-2">12/1/2021</div>
                        </div>
                    </div>

                    <div className="texting-box ds-flex">
                        <div className="ds-m-3 ds-mt-4 icon-bg">
                            <img className="" alt="friends-pict" src={attachImg} height="26"/>
                        </div>
                        <div className="ds-mt-3">
                            <Input className="texting-input" placeholder="Write your thought" />
                        </div>
                        <div className="ds-m-3 icon-bg">
                            <img alt="friends-pict" src={submitImg} height="28"/>
                        </div>
                    </div>
                </div>
            </div>
            {/* <div className="modal">modal</div> */}
        </div>
    )
}