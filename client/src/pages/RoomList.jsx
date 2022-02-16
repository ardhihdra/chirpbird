import React, { useState } from 'react';

import userImg from '../assets/img/icons/010-user.png';

export default function RoomList(props) {
    const rooms = props.rooms
    const onChooseGroup = props.onChooseGroup

    return (
        rooms.map(el => {
            return (
                <div className="room-preview-box flex" onClick={(e) => onChooseGroup(e, el.id)}>
                    <img
                        className="mt-2 mb-3 ml-1 room-pict"
                        alt="room-pict"
                        src={userImg}
                        width="32px"
                        height="32px" />
                    <div className="ml-5 room-title">
                        <div>{el.name}</div>
                        <div className="flex text-sm">
                            <div className="">room desc :</div>
                            <div className="ml-2">room chatting preview</div>
                        </div>
                    </div>
                    <div>
                        <div className="text-sm">12/1/2021</div>
                        <div className="room-unread text-sm mt-1 ml-3">{rooms.length}</div>
                    </div>
                </div>
            )
        })
    )
}
