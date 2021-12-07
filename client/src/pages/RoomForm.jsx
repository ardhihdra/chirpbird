import React, { useState, useEffect } from 'react';
import axios from 'axios';

import Input from '../components/Input';
import Button from '../components/Button';
import addImg from '../assets/img/icons/011-add.png';

import styles from './RoomForm.css';

export default function RoomForm(props) {
    const [name,setName] = useState('')
    const [member,setMember] = useState('')
    const [listMember,setListMember] = useState([])
    const [checkListMember,setCheckListMember] = useState([])

    const handleChange = (e) => {
        e.preventDefault();
        const input_name = e.target.name
        let value = e.target.value
        if(input_name === "name") {
            setName(value)
        }
        if(input_name === "member") {
            if(value.length > 3) {
                // get to be
                // checkListMember
            }
            setMember(value)
        }
    }

    const handleSubmit = (e) => {
        // if(props.onClick) props.onClick(e)
    }

    const addMember = (e) => {
        e.preventDefault();
        let m_listmember = [...listMember]
        m_listmember.push(member)
        setListMember(m_listmember)
        console.log("list mem", m_listmember)
    }

    const removeMember = (e,idx) => {
        e.preventDefault();
        let m_listmember = [...listMember]
        m_listmember.splice(idx,1)
        setListMember(m_listmember)
    }

    return (
        <form className="menu ds-p-5" onSubmit={handleSubmit}>
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
            { checkListMember.length ? <div className="txt-desc">User available : </div>: ''}
            { !checkListMember.length && member.length > 3 ? <div className="txt-desc">User not found</div>: ''}
            { listMember.length ? <div className="txt-desc ds-mb-2">Member List : </div>: ''}
            {   
                listMember.map((mem,idx) => {
                    return (
                        <div key={idx} className="txt-desc ds-p-1 removable" 
                            title="click to delete"
                            onClick={(e) => removeMember(e,idx)}>
                            {mem} {Object.keys(styles)}
                        </div>
                    )
                })
            }
            <Button className="ds-m-5" type="submit" 
                label="Submit" value="Submit"></Button>
        </form>
    )
}
