import React, { useState } from 'react';
import { useParams, useNavigate } from "react-router-dom";
import axios from 'axios';

import logoutImg from '../assets/img/icons/logout.png';
import Icon from '../components/Icon';

import styles from './RightMenu.css';
const MASTER_URL = `http://${process.env.REACT_APP_MASTER_URL}`

export default function RightMenu(props) {
    let navigate = useNavigate();
    const userinfo = JSON.parse(localStorage.getItem('userinfo'));

    const logout = async () => {
        const data = new FormData()
        data.append('id', userinfo.id)
        const resp = await axios.post(`${MASTER_URL}/logout`, data, {})
        if(resp instanceof Error) throw resp
        // must not comeback, remove session, refersh indices
        localStorage.removeItem('userinfo')
        localStorage.removeItem('messaging_url')
        sessionStorage.removeItem('token')
        sessionStorage.removeItem('access_token')
        navigate(`/login`)
    }

    return (
        <div className="right-menu ds-border ds-flex-column">
            <div className="right-menu-list">Hello, {userinfo.username}!</div>
            <div className="right-menu-footer ds-m-3">
                <Icon img={logoutImg} onClick={logout}/>
            </div>
        </div>
    )
}
