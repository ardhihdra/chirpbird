import React, { useState } from 'react';

import userImg from '../assets/img/icons/010-user.png';

export default function ChatItem(props) {
    const name = props.name
    const message = props.message
    const date = props.date
    const img = props.img

    return (
        <div className="chatbox ds-flex">
            <div className="ds-ml-5 ds-mt-1 ds-mb-1 room-pict">
                <img alt="friends-pict" src={img || userImg} height="40"/>
            </div>
            <div className="ds-ml-5 ds-mt-1">
                <div className="chat-username">{name}</div>
                <div className="ds-flex chat-content">
                    <div className="txt-desc-sm">{message}</div>
                </div>
            </div>
            <div>
                <div className="txt-desc-meta-sm ds-mt-2">{date}</div>
            </div>
        </div>
    )
}
