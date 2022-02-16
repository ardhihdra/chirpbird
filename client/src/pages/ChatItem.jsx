import React from 'react';

import userImg from '../assets/img/icons/010-user.png';

export default function ChatItem(props) {
    const name = props.name
    const message = props.message
    const date = props.date
    const img = props.img
    const isSelf = props.isSelf

    return (
        <div className="chatbox flex">
            <div className="ml-5 mt-3 mb-1">
                <img className="room-pict" alt="friends-pict" src={img || userImg} width="40"/>
            </div>
            <div className="ml-5 mt-1">
                <div className="chat-username text-sm" style={isSelf ? {color: 'orange'}:{}}>{name}</div>
                <div className="flex chat-content">
                    <div className="text-sm">{message}</div>
                </div>
            </div>
            <div>
                <div className="text-xs mt-2">{new Date(date).toDateString()}</div>
            </div>
        </div>
    )
}
