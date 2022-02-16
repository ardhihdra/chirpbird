import React from 'react';
import axios from 'axios';
import { useNavigate } from "react-router-dom";
import Input from '../../components/Input'
import Button from '../../components/Button'

import tele from '../../assets/img/icons/telegram.png';
import mail from '../../assets/img/icons/009-message.png';
/** thanks to https://gist.github.com/keeguon/2310008 */
import countries from "../../assets/js/countries"
import interests from "../../assets/js/interests"
import { getRandomUsername } from '../../assets/js/data';
const MASTER_URL = `http://${process.env.REACT_APP_MASTER_URL}`
const INTERESTS = interests

function WithNavigate(props) {
    let navigate = useNavigate();
    return <LoginForm {...props} navigate={navigate} />
}

class LoginForm extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            countries: countries, isLoaded: true, username: '', country: '', interests: [], profile: '', page: 0, totalPage: 2,
            rememberme: false, error_message: null, username_timeout: null};
    }

    calcPrice = (name, value) => {
        let total = 0
        const count = this.state[name] > 0 ? this.state[name]: 0
        const diff = (value - count) * this.products_price[name]
        total = this.state.total + diff
        
        return total
    }

    handleChange = (event) => {
        // console.log("diff price", this.state[event.target.name], event.target.value)
        const name = event.target.name
        let value = event.target.value
        if(event.target.type === 'checkbox') {
            let listCheckbox = [...this.state[name]]
            listCheckbox.push(value)
            value = listCheckbox
        }
        if(event.target.name === 'username') {
            if(this.state.username_timeout) clearTimeout(this.state.username_timeout)
            const timeout = setTimeout(async () => {
                console.log("Checking uname...")
                const resp = await axios.post(`${MASTER_URL}/username?username=${value}`, {
                    headers: { "Content-Type": 'multipart/form-data' }
                })
                const isValid = resp.data.valid
                if(!isValid) {
                    this.setState({error_message: 'Username is taken'})
                } else {
                    this.setState({error_message: null})
                }
                this.setState({username_timeout: null})
            }, 1000)
            this.setState({username_timeout: timeout})
        }
        this.setState({[name]: value})
    }

    handleSubmit = (event) => {
        event.preventDefault();
        // const phoneIndoRegex = /\+?([ -]?\d+)+|\(\d+\)([ -]\d+)/
        const nameExist = this.state.username && this.state.username.length
        if(!nameExist) this.setState({error_message: 'Username cannot be empty!\n'})
        if(nameExist) {
            this.setState({'isConfirmationShow': true})
            this.setState({'error_message': ''})
        }
        const data = new FormData()
        data.append('username', this.state['username'])
        data.append('country', this.state['country'])
        data.append('interests', this.state['interests'])
        data.append('profile', this.state['profile'])

        axios.post(`${MASTER_URL}/register`, data, {
                headers: { "Content-Type": 'multipart/form-data' }
        }).then(async response => {
            if(response instanceof Error) throw response
            localStorage.setItem('userinfo', JSON.stringify(response.data.data));
            sessionStorage.setItem('token', response.data.token);

            const config = {
                headers: { 'Authorization': 'Bearer ' + response.data.token }
            }
            const resp = await axios.post(`${MASTER_URL}/sessions`, {}, config)
            sessionStorage.setItem('access_token', resp.data.access_token);
            this.props.navigate(`/`)
        }).catch((error) => {
            alert('Login failed', error)
            console.log(error);
        })
    }

    showCounter(name, e) {
        e.preventDefault()
        const value = 1
        const total = this.calcPrice(name, value)
        this.setState({[name]: value, total})
    }


    changePage = (val) => {
        if(val === 1) {
            if(!this.state.username || !this.state.username.length) {
                this.setState({error_message: 'Username cannot be null!\n'})
            } else if(this.state.error_message) {
                this.setState({error_message: this.state.error_message + '!!'})
            } else {
                this.setState({error_message: null})
                this.setState({page: val})
            }
        } else {
            this.setState({page: val})
            this.setState({error_message: ''})
        }
    }

    quickLogin = (e) => {
        const uname = getRandomUsername()
        this.setState({
            username: uname,
            country: countries[0],
            interests: [],
            profile: ''
        });
        this.handleSubmit(e)
    }

    styling() {
        const button = {
            cursor: 'pointer',
            width: '6.5rem',
            // width: '100%',
            height: '2.5rem',
            marginLeft: 'auto',
            marginRight: 'auto',
            border: '2px solid',
            borderRadius: 'var(--border-radius-norm)',
        }
        const largeButton = {
            width: '100%',
            // width: '8rem',
            fontSize: 'larger'
        }
        return {
            bluecolor: {
                color: 'var(--primary-font)'
            },
            listproducts: {
                // display: 'flex',
                width: '100%',
                // overflowX: 'scroll'
            },
            productsimg: {
                'width': '11rem',
                'borderRadius': '8px',
                'justifyContent': 'center'
            },
            ordercounter: {
                'display': 'flex',
                'flexDirection': 'column',
                borderRadius: 'var(--border-radius-norm)',
                boxShadow: '5px 2px 8px #c3c3c3'
            },
            button: button,
            buttonClear: {
                ...button,
                backgroundColor: 'white',
                borderColor: 'var(--blue-soft)',
                color: 'var(--blue-soft)',
                cursor: 'pointer'
            },
            largeButton: largeButton,
            buttonSolid: {
                ...button,
                ...largeButton,
                backgroundColor: 'var(--blue-soft)',
                borderColor: 'var(--blue-soft)',
                color: 'white',
            },
            cancelButtonSolid: {
                ...button,
                ...largeButton,
                backgroundColor: 'white',
                borderColor: 'var(--red)',
                color: 'var(--red)'
            },
            flex: {
                display: 'flex',
                flexDirection: 'column',
                marginLeft: 'auto',
                marginRight: 'auto'
            },
            flexRow: {
                display: 'flex',
                flexDirection: 'row',
                justifyContent: 'center'
            },
            summaryBox: {
                backgroundColor: Number(this.state.total) ? 'white': 'transparent',
                maxWidth: '400px',
                marginLeft: 'auto',
                marginRight: 'auto'
            }
        }
    }

    render() {
        const { error, isLoaded, countries } = this.state;
        const style = this.styling()
        const countryList = countries

        const page1 = (
            <div className="ds-fade-in py-6">
                <h5 className="mb-3" style={style.bluecolor}>
                    Fill your profile :
                </h5>
                <div style={style.flex}>
                    <Input
                        className="mb-5"
                        name="username" 
                        value={this.state.username} 
                        placeholder="Username" 
                        onChange={this.handleChange} 
                        required={true}
                    />
                    <select 
                        id="country"
                        name="country"
                        value={this.state.country}
                        onChange={this.handleChange}
                        className="ds-select"
                    >
                        { countryList.map((ctr, idx) => {
                            return (
                                <option className="ds-option" key={idx} value={ctr.name}>{ctr.name}</option>
                                )
                            })
                        }
                    </select>
                </div>
                <div className="m-6 text-red-400">{this.state.error_message}</div>
                <Button 
                    className="m-2 mb-5"
                    type="button" 
                    onClick={() => this.changePage(1)}
                    label="Next"
                />
                <hr class="ds-hr"/>

                <h5 className="m-4 mt-4" style={style.bluecolor}>
                    Or :
                </h5>
                <Button 
                    type="button"
                    style="clear"
                    onClick={this.quickLogin}
                    label="Go Annonymous"
                />
                {/* <div className="mt-5 mb-5" style={style.flexRow}>
                    <div className="ds-p-1" target="_blank" rel="noreferrer">
                        <img className="logo" alt="telegram" src={tele} title="click to open"></img>
                        <div type="button" style={style.bluecolor}>Login Via Telegram</div>
                    </div>
                    <div className="ds-p-1 ml-4" target="_blank" rel="noreferrer">
                        <img className="logo" alt="email" src={mail} title="click to open"></img>
                        <div type="button" style={style.bluecolor}>Login Via Email</div>
                    </div>
                </div> */}
            </div>
        )

        const page2 = (
            <div className="ds-fade-in">
                <div className="txt-desc mt-4 mb-4 " style={style.bluecolor} htmlFor="country">
                    What interest you:
                </div>
                <div className="grid grid-cols-2 mb-5">
                    { INTERESTS.map(int => {
                        return (
                            <div key={int.id} className="grid grid-cols-2 mt-1 mx-2 text-left" style={style.flexRow}>
                                <input id={int.id+int.name} name="interests" type="checkbox" className="mt-1 mr-5"
                                    value={int.name} onChange={this.handleChange} />
                                <div className="txt-desc w-full" style={style.bluecolor} htmlFor={int.id+int.name}>{int.name}</div>
                            </div>
                        )
                    })}
                </div>
                <Input
                    className="mt-3"
                    label="Profile" 
                    name="profile" 
                    value={this.state.profile} 
                    type="textarea"
                    placeholder="Add some details so people can know you better" 
                    onChange={this.handleChange}
                />
                <Button 
                    className="m-2 mb-5"
                    type="button"
                    style="clear"
                    color="danger"
                    label="Back"
                    onClick={() => this.changePage(0)}
                />
                <Button 
                    className="m-2 mb-5"
                    type="submit"
                    value="Submit"
                    label="Submit"
                />
            </div>
        )

        if(error) return <div>Error: {error.message} </div>;
        else if(!isLoaded) return <div>Loading...</div>;
        else 
        return (
            <div class="login-form">
                <form onSubmit={this.handleSubmit}>
                { 
                this.state.page === 0 ?
                    page1
                    :
                    page2
                }
                </form>
            </div>
        );
    }
}

export default WithNavigate;