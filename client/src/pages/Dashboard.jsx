import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, useNavigate } from "react-router-dom";

import Input from '../components/Input';
import menuImg from '../assets/img/icons/menu.png';
import menuDotImg from '../assets/img/icons/menu-dot.png';
import userImg from '../assets/img/icons/010-user.png';
import attachImg from '../assets/img/icons/attachment.png';
import submitImg from '../assets/img/icons/submit.png';

import RoomForm from './RoomForm.jsx';
import RoomList from './RoomList.jsx';
import ChatItem from './ChatItem.jsx';
import RightMenu from './RightMenu.jsx';

import './Dashboard.css';
const MASTER_URL = `http://${process.env.REACT_APP_MASTER_URL}`

export default function Home(props) {
    let navigate = useNavigate();
    const [isProfileModalShow, setProfileModalShow] = useState(0)
    const [menu, setMenu] = useState(false)
    const [rightMenu, setRightMenu] = useState(false)
    const { id } = useParams();
    const userinfo = localStorage.getItem('userinfo');
    const payload = localStorage.getItem('access_token')
    const config = {
        headers: { 'Authorization': 'Bearer ' + localStorage.getItem('token') }
    }
    const activeRooms = [2,2,3]
    
    const openMenu = (e) => {
        setMenu(!menu)
    }

    const openRightMenu = () => {
        setRightMenu(!rightMenu)
    }

    useEffect(() => {
        //componentDidMount
        if(!userinfo) {
            navigate(`/login`)
        } else {
            const user = JSON.parse(userinfo)
            axios.get(`${MASTER_URL}/users?id=${user.id}`, {}, config)
                .then(resp => {
                    if(resp.error || resp instanceof Error) {
                        console.log("unauthorized")
                        navigate(`/login`)
                    }
                    initSocket(`ws://${process.env.REACT_APP_MASTER_URL}/messaging?access_token=${JSON.parse(payload)}`)
            })
        }
    }, [])

    const initSocket = (url) => {
        const socket = new WebSocket(url);
        socket.onopen = function() {
            console.log("socket opened")
            socket.send(JSON.stringify({"method":20, "timestamp":new Date().getTime()}));
        };
        socket.onmessage = function(event) {
            console.log(event.data);
        };
    }

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
                    <img alt="rooms-menu" src={menuDotImg} height="22px" onClick={openRightMenu}/>
                </div>
            </div>
            <div className="dashboard-container ds-slide-in">
                <div className="left-container ds-fade-in">
                    {menu ? (<RoomForm />) : (<RoomList rooms={activeRooms}/>)}
                </div>
                <div className="chat-container ds-fade-in">
                    <ChatItem name="friends name" img={userImg} 
                        message={`room chat text ${userinfo}`} date="12/1/2021"/>

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
                { rightMenu ? <RightMenu />: ''}
            </div>
            {/* <div className="modal">modal</div> */}
        </div>
    )
}