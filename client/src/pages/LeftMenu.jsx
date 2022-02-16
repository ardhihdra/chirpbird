import React, { useState, useEffect } from 'react';
import { useNavigate } from "react-router-dom";
import axios from 'axios';

import Icon from '../components/Icon';
import Button from '../components/Button';
import RoomForm from './RoomForm.jsx';
import closeImg from '../assets/img/icons/close.png';

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
    const isOpenNew = props.roomForm || false

    const [openRoomForm, setOpenRoomForm] = useState(isOpenNew)
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

    useEffect(() => {
        setOpenRoomForm(isOpenNew)
    }, [isOpenNew])

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
                    {/* <Icon className="ds-mt-3 mr-4 ml-auto" img={closeImg} onClick={roomFormControl} title="cancel"/> */}
                    <RoomForm onSuccess={updateGroup}/>
                </div>
                :
                <div className="flex flex-col ds-fade-in p-4">
                    <div className="mb-5">
                        <select
                            className="ds-select"
                            id="country"
                            name="country"
                            value={filterCountry}
                            onChange={handleChange}>
                            { countryList.map((ctr, idx) => {
                                return (
                                    <option className="ds-option" key={idx} value={ctr.name}>{ctr.name}</option>
                                    )
                                })}
                        </select>
                    </div>
                    <div className="mb-5">
                        <div>Filter interests</div>
                        <div className="grid grid-cols-3">
                            { interests.map(int => {
                                return (
                                    <div key={int.id} className="grid grid-cols-4 text-left mr-5">
                                        <input 
                                            id={int.id+int.name}
                                            name="interests"
                                            type="checkbox" 
                                            className="mr-2 mt-1"
                                            value={int.name} onChange={handleChange} />
                                        <div className="col-span-3 text-sm w-full" htmlFor={int.id+int.name}>{int.name}</div>
                                    </div>
                                )
                            })}
                        </div>
                    </div>
                    <Button 
                        className="m-auto" 
                        onClick={searchRooms}
                        label="Filter">
                    </Button>
                    <div className="ds-col-12 m-5">
                        <hr/>
                        <div className="m-6">Search Result</div>
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
