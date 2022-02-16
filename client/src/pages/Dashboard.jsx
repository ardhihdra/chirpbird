import React, { useState, useEffect, createRef } from 'react';
import axios from 'axios';
import { useNavigate } from "react-router-dom";

import Input from '../components/Input';
import { EVENT_TYPE, fetchGroups } from '../assets/js/data';
import menuImg from '../assets/img/icons/menu.png';
import settingImg from '../assets/img/icons/settings.png';
import addImg from '../assets/img/icons/011-add.png';
import userImg from '../assets/img/icons/010-user.png';
import attachImg from '../assets/img/icons/attachment.png';
import submitImg from '../assets/img/icons/submit.png';

import LeftMenu from './LeftMenu.jsx';
import RoomList from './RoomList.jsx';
import ChatItem from './ChatItem.jsx';
import RightMenu from './RightMenu.jsx';

import './Dashboard.css';
const MASTER_URL = `http://${process.env.REACT_APP_MASTER_URL}`

export default function Home(props) {
    let navigate = useNavigate();
    const textInput = createRef()

    const [menu, setMenu] = useState(false)
    const [newRoom, setNewRoom] = useState(false)
    const [rightMenu, setRightMenu] = useState(false)
    const [roomName, setRoomName] = useState()
    const [groupList, setGroupList] = useState([])
    const [activeGroup, setActiveGroup] = useState({})
    const [messagesPool, setMessagesPool] = useState({})
    const [socketConn, setSocketConn] = useState()
    const [wsReceived, setWsReceived] = useState('{}')

    const [chatText, setChatText] = useState()
    /** list of groupid 
     * [groupid]: {
     *  messages: { [msgid]: { data, isread, isdelivered } }, 
     *  startType: {}, endType: {}, join: {}, left: {}   
     * }
     * 
    */
    // const { id } = useParams();
    const userinfo = JSON.parse(localStorage.getItem('userinfo')) || {};
    const payload = sessionStorage.getItem('access_token')
    const token = sessionStorage.getItem('token')
    const config = {
        headers: { 'Authorization': 'Bearer ' + token }
    }
    
    const openMenu = (e) => {
        setMenu(!menu)
        setNewRoom(false)
    }

    const openRightMenu = () => {
        setRightMenu(!rightMenu)
    }

    const handleSearch = (event) => {
        const value = event.target.value
        setMenu(true)
        setRoomName(value)
    }

    const handleChange = (event) => {
        const value = event.target.value
        setChatText(value)
    }

    useEffect(() => {
        //componentDidMount
        if(!userinfo.id) navigate(`/login`)
        else {
            const user = userinfo
            axios.post(`${MASTER_URL}/users?id=${user.id}`, {}, config)
                .then(resp => {
                    if(resp.error || resp instanceof Error) throw resp
                    getGroups()
                    if(groupList.length === 1) setActiveGroup(groupList[0])
                    initSocket(`ws://${process.env.REACT_APP_MASTER_URL}/messaging?access_token=${payload}`)
            }).catch(err => {
                console.log("unauthorized", err)
                navigate(`/login`)
            })

        }
    }, [])

    useEffect(() => {
        parseSocketMessage(wsReceived, messagesPool)
    }, [wsReceived])

    const getGroups = () => {
        fetchGroups(userinfo).then(result => {
            setGroupList(result)
        }).catch(err => {
            console.log(err)
        })
    }

    const openNew = () => {
        setMenu(!menu)
        setNewRoom(true)
    }

    const initSocket = (url) => {
        const socket = new WebSocket(url)
        setSocketConn(socket)
        socket.onopen = function() {
            console.log("socket opened")
            socket.send(JSON.stringify({"method":20, "timestamp":new Date().getTime()}));
        };
        socket.onmessage = function(event) {
            console.log("on message")
            setWsReceived(event.data)
        };
        socket.onerror = function(event) {
            console.log("SOCKET ERROR", event)
            socket.close()
        }
        socket.onclose = function() {
            console.log("socket closed")
        }
    }

    const appendToPool = (source, id, group) => {
        setMessagesPool({
            ...source,
            [id]: group
        })
    }

    const parseSocketMessage = (data, source) => {
        const event = JSON.parse(data)
        const body = event.body
        if(!body || !source) {
            console.log("invalid ws body or source")
            return
        }
        console.log("incoming", event)
        let group = {...source[body.group_id]}
        if(!group) group = {}
        switch(event.type) {
            case EVENT_TYPE.EVENT_MESSAGE:
                group[body.message_id] = {timestamp: event.timestamp, ...body}
                appendToPool(source, body.group_id, group)
                break;
            case EVENT_TYPE.EVENT_MESSAGE_SENT:
                if(!group[body.message_id]) {
                    group[body.message_id] = {timestamp: event.timestamp, ...body, issent: true}
                } else {
                    group[body.message_id].issent = true
                }
                appendToPool(source, body.group_id, group)
                // if(group[body.messageid]) message[body.messageid] = {}
                break;
            case EVENT_TYPE.EVENT_MESSAGE_DELIVERED:
                // if(group[body.messageid]) message[body.messageid] = {}
                break;
            case EVENT_TYPE.EVENT_MESSAGE_READ:
                // if(group[body.messageid]) message[body.messageid] = {}
                break;
            case EVENT_TYPE.EVENT_GROUP:
                break;
            case EVENT_TYPE.EVENT_GROUP_JOINED:
                break;
            case EVENT_TYPE.EVENT_GROUP_LEFT:
                break;
        }
    }

    const openDetailGroup = (event, groupid) => {
        const active = groupList.filter(gl => gl.id === groupid)
        setActiveGroup(active[0])
    }

    const sendMessage = (event) => {
        if(activeGroup.id && chatText) {
            socketConn.send(JSON.stringify({"method":40,"body":{"group_id":activeGroup.id, "data":chatText}}));
        }
    }

    return (
        <div>
            <div className="header-menu w-full grid grid-cols-12 p-2">
                <div className="mr-4 flex">
                    <img
                        className="icon-bg p-1" 
                        alt="profile-menu" 
                        height="36px"
                        width="36px" 
                        src={menuImg} 
                        title="menu"
                        onClick={openMenu}/>
                    <img
                        className="icon-bg ml-1 p-1"
                        alt="profile-menu" 
                        width="34px" 
                        src={addImg}
                        title="buat ruang baru"
                        onClick={openNew}/>
                </div>
                <div className="col-span-3">
                    <Input
                        placeholder="Search" 
                        onChange={handleSearch} 
                        value={roomName}></Input>
                </div>
                <div className="ml-5 pl-5">
                    <img
                        className="room-pict"
                        alt="profile-pict"
                        height="36px"
                        width="36px"
                        src={userImg}/>
                </div>
                <div className="ml-5 mt-1 col-span-3 text-left">
                    <div>{activeGroup.name || (userinfo.username ? userinfo.username: '-')}</div>
                    <div className="txt-desc-meta">
                        {activeGroup.user_ids ? activeGroup.user_ids.length: userinfo.country}
                    </div>
                </div>
                <div className="ml-6 pl-6 col-span-3 room-search">
                    {/* <Input className="transparent night-mode" placeholder="Search in chat"></Input> */}
                </div>
                <div className="ml-6 pl-6">
                    <img
                        className="icon-bg p-1"
                        alt="rooms-menu"
                        height="36px"
                        width="36px"
                        src={settingImg}
                        onClick={openRightMenu}/>
                </div>
            </div>
            <div className="dashboard-container ds-slide-in">
                <div className="left-container drop-shadow-xl">
                    {menu ? 
                        (<LeftMenu roomName={roomName} roomForm={newRoom} openForm={groupList.length} onUpdateGroup={getGroups} />) : 
                        ( groupList.length === 0 ? 
                            <div className="m-2">You don't have any room yet</div> :
                            <RoomList rooms={groupList} onChooseGroup={openDetailGroup}/>
                        )
                    }
                </div>
                <div className="chat-container ds-fade-in">
                    <div className="txt-desc-meta-sm">{activeGroup.name || "no rooms selected"}</div>
                    {
                        activeGroup.id && messagesPool[activeGroup.id] && 
                        Object.values(messagesPool[activeGroup.id]).length ? 
                            Object.values(messagesPool[activeGroup.id]).map((mp,idx) => {
                                return (
                                    <ChatItem key={idx} name={mp.sender_name || 'anonymous'} img={userImg}
                                        isSelf={mp.sender_id === userinfo.id}
                                        message={mp.data} date={mp.timestamp}/>
                                )
                            })
                            :
                            ''
                    }
                    <div className="texting-box flex">
                        <div className="m-3 mt-4">
                            <img 
                                className="icon-bg"
                                alt="friends-pict"
                                src={attachImg}
                                width="26"/>
                        </div>
                        <div className="mt-3">
                            <Input ref={textInput} className="texting-input" placeholder="Write your thought" onChange={handleChange} value={chatText} />
                        </div>
                        <div className="m-4">
                            <img
                                className="icon-bg"
                                alt="friends-pict"
                                src={submitImg}
                                length="26"
                                width="26"
                                onClick={sendMessage}/>
                        </div>
                    </div>
                </div>
                { rightMenu ? <RightMenu />: ''}
            </div>
            {/* <div className="modal">modal</div> */}
        </div>
    )
}