import React, { useState, useEffect, createRef } from 'react';
import axios from 'axios';
import { useNavigate } from "react-router-dom";

import Input from '../components/Input';
import { EVENT_TYPE, fetchGroups } from '../assets/js/data';
import menuImg from '../assets/img/icons/menu.png';
import menuDotImg from '../assets/img/icons/menu-dot.png';
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
    const userinfo = JSON.parse(localStorage.getItem('userinfo'));
    const payload = sessionStorage.getItem('access_token')
    const token = sessionStorage.getItem('token')
    const config = {
        headers: { 'Authorization': 'Bearer ' + token }
    }
    
    const openMenu = (e) => {
        setMenu(!menu)
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
        if(!userinfo) {
            navigate(`/login`)
        } else {
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
        console.log("keaskdfyy", event, chatText, activeGroup)
        if(activeGroup.id && chatText) {
            socketConn.send(JSON.stringify({"method":40,"body":{"group_id":activeGroup.id, "data":chatText}}));
        }
    }

    return (
        <div>
            <div className="header-menu">
                <div className="ds-mr-5 icon-bg">
                    <img alt="profile-menu" src={menuImg} height="28" onClick={openMenu}/>
                </div>
                <div className="ds-mr-5 search-input">
                    <Input placeholder="Search" onChange={handleSearch} value={roomName}></Input>
                </div>
                <div className="ds-ml-1 ds-mt-1 ds-mb-1 room-pict">
                    <img alt="profile-pict" src={userImg} height="30"/>
                </div>
                <div className="ds-ml-5 ds-mt-1 room-title">
                    <div>{activeGroup.name || (userinfo.username ? userinfo.username: '-')}</div>
                    <div className="txt-desc-meta">
                        {activeGroup.user_ids ? activeGroup.user_ids.length: userinfo.country}
                    </div>
                </div>
                <div className="ds-mr-5 room-search">
                    <Input className="transparent night-mode" placeholder="Search in chat"></Input>
                </div>
                <div className="ds-mt-1 ds-mr-5 icon-bg" onClick={openRightMenu}>
                    <img alt="rooms-menu" src={menuDotImg} height="22px"/>
                </div>
            </div>
            <div className="dashboard-container ds-slide-in">
                <div className="left-container ds-fade-in">
                    {menu || groupList.length === 0 ? 
                        (<LeftMenu roomName={roomName} openForm={groupList.length} onUpdateGroup={getGroups} />) : 
                        (<RoomList rooms={groupList} onChooseGroup={openDetailGroup}/>)}
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
                    <div className="texting-box ds-flex">
                        <div className="ds-m-3 ds-mt-4 icon-bg">
                            <img className="" alt="friends-pict" src={attachImg} height="26"/>
                        </div>
                        <div className="ds-mt-3">
                            <Input ref={textInput} className="texting-input" placeholder="Write your thought" onChange={handleChange} value={chatText} />
                        </div>
                        <div className="ds-m-3 icon-bg">
                            <img alt="friends-pict" src={submitImg} height="28" onClick={sendMessage}/>
                        </div>
                    </div>
                </div>
                { rightMenu ? <RightMenu />: ''}
            </div>
            {/* <div className="modal">modal</div> */}
        </div>
    )
}