import React, { useState } from 'react';

import userImg from '../assets/img/icons/010-user.png';

export default function RoomList(props) {
    const rooms = props.rooms

    return (
        rooms.map(el => {
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
}
