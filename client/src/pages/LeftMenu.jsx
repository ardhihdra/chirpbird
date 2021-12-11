import React, { useState, useEffect } from 'react';
import { useNavigate } from "react-router-dom";
import axios from 'axios';

import Icon from '../components/Icon';
import Button from '../components/Button';
import RoomForm from './RoomForm.jsx';
import closeImg from '../assets/img/icons/close.png';
import addImg from '../assets/img/icons/011-add.png';

import './LeftMenu.css';
import countries from "../assets/js/countries"
import interests from "../assets/js/interests"
const MASTER_URL = `http://${process.env.REACT_APP_MASTER_URL}`

export default function RightMenu(props) {
    let navigate = useNavigate();

    const userinfo = JSON.parse(localStorage.getItem('userinfo'));
    const countryList = countries
    const roomName = props.roomName
    const onUpdateGroup = props.onUpdateGroup

    const [openRoomForm, setOpenRoomForm] = useState(false)
    const [filterInterests, setFilterInterests] = useState([])
    const [filterCountry, setFilterCountry] = useState()
    const [searchResult, setSearchResult] = useState({ groups: undefined, chatting: undefined})

    const handleChange = (event) => {
        const name = event.target.name
        let value = event.target.value
        if(event.target.type === 'checkbox') {
            let listCheckbox = [...filterInterests]
            listCheckbox.push(value)
            setFilterInterests(listCheckbox)
        }
        if(name === 'country') {
            setFilterCountry(value)
        }
    }

    useEffect(() => {
        searchRooms()
    }, [roomName])

    const searchRooms = async (e) => {
        const data = new FormData()
        data.append('search', roomName)
        if(filterCountry) data.append('country', filterCountry)
        if(filterInterests) data.append('interests', filterInterests)
        axios.post(`${MASTER_URL}/search`, data, {}).then(resp => {
            if(resp instanceof Error) throw resp
            setSearchResult({
                groups: resp.data.groups,
                chatting: resp.data.chatting,
            })
        }).catch(err => {
            console.log(err)
        })
    }

    const logout = async () => {
        const data = new FormData()
        data.append('id', userinfo.id)
        const resp = await axios.post(`${MASTER_URL}/logout`, data, {})
        if(resp instanceof Error) throw resp
        // must not comeback, remove session, refersh indices
        localStorage.removeItem('userinfo')
        localStorage.removeItem('token')
        localStorage.removeItem('access_token')
        navigate(`/login`)
    }

    const roomFormControl = () => {
        setOpenRoomForm(!openRoomForm)
    }

    const updateGroup = () => {
        onUpdateGroup()
    }

    return (
        <div>
            { openRoomForm ? 
                <div>
                    <Icon className="ds-mt-3 ds-mr-4 ds-ml-auto" img={closeImg} onClick={roomFormControl} title="cancel"/>
                    <RoomForm onSuccess={updateGroup}/>
                </div>
                :
                <div className="ds-row ds-fade-in">
                    <div className="left-menu-toolbox ds-col-8">
                        <select id="country" name="country" value={filterCountry} onChange={handleChange}>
                            { countryList.map((ctr, idx) => {
                                return (
                                    <option className="ds-option" key={idx} value={ctr.name}>{ctr.name}</option>
                                    )
                                })}
                        </select>
                    </div>
                    <div className="left-menu-toolbox ds-col-4">
                        <div className="ds-flex ds-pt-2">
                            <div className="txt-desc-meta ds-pt-3 ds-pl-2">New room</div>
                            <Icon img={addImg} onClick={roomFormControl} title="Create new room"/>
                        </div>
                    </div>
                    <div className="left-menu-toolbox ds-col-12">
                        <div>Filter interests</div>
                        <div className="ds-flex">
                            { interests.map(int => {
                                return (
                                    <div key={int.id} className="ds-m-3 ds-flex-col">
                                        <input id={int.id+int.name} name="interests" type="checkbox" className="ds-mr-5"
                                            value={int.name} onChange={handleChange} />
                                        <label className="txt-desc-meta" htmlFor={int.id+int.name}>{int.name}</label>
                                    </div>
                                )
                            })}
                        </div>
                    </div>
                    <div className="ds-col-12">
                        <Button className="ds-m-auto" onClick={searchRooms}
                            label="Filter"></Button>
                    </div>
                    <div className="ds-col-12 ds-m-5">
                        <hr/>
                        <div className="ds-m-6">Search Result</div>
                        <div className="ds-flex-col">
                            {   searchResult.groups  ?
                                searchResult.groups.map((sr,idx) => {
                                    <div className="search-room-list">{sr.name}</div>
                                }) : '-'
                            }
                        </div>
                    </div>
                </div>
            }
        </div>
    )
}
