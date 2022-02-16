import axios from 'axios';

const config = (token) => {
    return { headers: { 'Authorization': 'Bearer ' + token } }
}
const MASTER_URL = `http://${process.env.REACT_APP_MASTER_URL}`

const getRandomUsername = () => {
    const listNames = ['Elon', 'Musk', 'Jeff', 'Bezos', 'Jungkook', 'John', 'Doe', 'Ujang', 'Hitlar', 'Puton']
    const rand = Math.floor(Math.random() * (listNames.length - 1))
    const randNum = Math.floor(Math.random() * 1000)
    return `${listNames[rand]}${randNum}`
}

const fetchPosts = () => {
    return [
        {
            id: 1,
            title: "Ini Judul",
            // highlight: img2,
            description: "Every human environment is embedded with its own intellectual",
            date: "Fri 12 March 2021",
            situation: "lagi cape",
            text: `<p>The export statement is used when creating JavaScript modules to export live bindings to functions, objects, or primitive values from the module so they can be used by other programs with the import statement. Bindings that are exported can still be modified locally; when imported, although they can only be read by the importing module the value updates whenever it is updated by the exporting module.</p>
            <p>The export statement is used when creating JavaScript modules to export live bindings to functions, objects, or primitive values from the module so they can be used by other programs with the import statement. Bindings that are exported can still be modified locally; when imported, although they can only be read by the importing module the value updates whenever it is updated by the exporting module.</p>
            <p>The export statement is used when creating JavaScript modules to export live bindings to functions, objects, or primitive values from the module so they can be used by other programs with the import statement. Bindings that are exported can still be modified locally; when imported, although they can only be read by the importing module the value updates whenever it is updated by the exporting module.</p>
            <p>The export statement is used when creating JavaScript modules to export live bindings to functions, objects, or primitive values from the module so they can be used by other programs with the import statement. Bindings that are exported can still be modified locally; when imported, although they can only be read by the importing module the value updates whenever it is updated by the exporting module.</p>
            `
        },
        {
            id: 2,
            title: "Ini Lainya",
            // highlight: img2,
            description: "Every human is crazy with its own intellectual",
            date: "Fri 14 March 2021",
            situation: "segar"
        },
        {
            id: 3,
            title: "Ini Gatau",
            // highlight: img2,
            description: "Every human is unique with its own intellectual",
            date: "Fri 06 April 2021",
            situation: "lagi ngelamun"
        }
    ];
}

const fetchGroups = async (userinfo) => {
    const token = sessionStorage.getItem('token')
    const rooms = await axios.get(`${MASTER_URL}/rooms?user_id=${userinfo.id}`, {}, config(token))
    if(rooms.error || rooms instanceof Error) throw rooms
    const result = rooms.data.groups || []
    return result
}

const EVENT_TYPE = {
	EVENT_MESSAGE           : 20,
	EVENT_MESSAGE_SENT      : 21,
	EVENT_MESSAGE_DELIVERED : 22,
	EVENT_MESSAGE_READ      : 23,
	//EVENT_MESSAGE_UPDATED   : 24,
	//EVENT_MESSAGE_DELETED   : 25,
	EVENT_TYPING_START : 40,
	EVENT_TYPING_END   : 41,
	EVENT_GROUP        : 70,
	//EVENT_GROUP_UPDATED     : 71,
	EVENT_GROUP_JOINED : 72,
	EVENT_GROUP_LEFT   : 73,
}

export {
    getRandomUsername,
    fetchPosts,
    fetchGroups,
    EVENT_TYPE
}