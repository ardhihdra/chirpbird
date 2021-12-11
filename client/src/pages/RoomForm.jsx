import React, { useState, useEffect } from 'react';
import axios from 'axios';

import Input from '../components/Input';
import Button from '../components/Button';
import addImg from '../assets/img/icons/011-add.png';

import styles from './RoomForm.css';

export default function RoomForm(props) {
    const onSuccess = props.onSuccess
    const [name,setName] = useState('')
    const [member,setMember] = useState('')
    const [listMember,setListMember] = useState([])
    const [checkListMember,setCheckListMember] = useState([])
    const userinfo = JSON.parse(localStorage.getItem('userinfo'));
    const token = sessionStorage.getItem('token')
    const config = {
        headers: { 'Authorization': 'Bearer ' + token }
    }

    const handleChange = (e) => {
        e.preventDefault();
        const input_name = e.target.name
        let value = e.target.value

        if(input_name === "name") {
            setName(value)
        }
        if(input_name === "member") {
            setMember(value)
        }
    }

    useEffect(() => {
        // TO DO : add debounce
        const search = async () => {
            const users = await axios.post(`${global.MASTER_URL}/users?username=${member}`, {}, config)
            if(users.error || users instanceof Error) {
                console.log(users)
            }
            const result = users.data.userinfo || []
            setCheckListMember(result.filter(res => res.username !== userinfo.username))
            return users
        }
        search()
    }, [member, userinfo.username])

    const handleSubmit = async (e) => {
        e.preventDefault()
        // if(props.onClick) props.onClick(e)
        const data = new FormData()
        data.append('name', name)
        data.append('user_ids', listMember.map(lm => lm.id))
        const created = await axios.post(`${global.MASTER_URL}/groups`, data, config)
        if(created.error || created instanceof Error) {
            return
        }
        console.log(created)
        onSuccess()
    }

    const addMember = (e) => {
        e.preventDefault();
        if(!member) {
            return
        }
        if(listMember.includes(member)) {
            console.log("user already add")
            return
        }
        let m_listmember = [...listMember]
        // TO DO add user by click
        const toBeAdd = checkListMember.filter(clm => clm.username === member)
        m_listmember.push(toBeAdd[0])
        console.log("tobe add", toBeAdd)
        setListMember(m_listmember)
    }

    const removeMember = (e,idx) => {
        e.preventDefault();
        let m_listmember = [...listMember]
        m_listmember.splice(idx,1)
        setListMember(m_listmember)
    }

    return (
        <form className="menu ds-ml-5 ds-mr-5" onSubmit={handleSubmit}>
            <div className="ds-m-3">Create New Room</div>
            <Input name="name" placeholder="Room name" onChange={handleChange} value={name}/>
            <div className="ds-row">
                <div className="ds-col-10">
                    <Input name="member" placeholder="Add room member (username)" onChange={handleChange} value={member}/>
                </div>
                <div className="ds-col-2 ds-ml-4 icon-bg">
                    <img alt="add-member" className="icon" src={addImg} height="28" onClick={addMember}/>
                </div>
            </div>
            { checkListMember.length ? <div className="txt-desc">User available : {checkListMember.map(cl => cl.username + ", ")}</div>: ''}
            { !checkListMember.length && member.length > 3 ? <div className="txt-desc">User not found</div>: ''}
            { listMember.length ? <div className="txt-desc ds-mb-2">Member List : </div>: ''}
            {   
                listMember.map((mem,idx) => {
                    return (
                        mem && mem.id ? 
                        <div key={idx} className="txt-desc ds-p-1 removable" 
                            title="click to delete"
                            onClick={(e) => removeMember(e,idx)}>
                            {mem.username} {Object.keys(styles)}
                        </div> : ''
                    )
                })
            }
            <Button className="ds-m-5" type="submit" 
                label="Submit" value="Submit"></Button>
        </form>
    )
}
