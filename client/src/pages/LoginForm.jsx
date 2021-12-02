import React from 'react';
import axios from 'axios';
import { useNavigate } from "react-router-dom";
import Input from '../components/Input'

import tele from '../assets/img/icons/telegram.png';
import mail from '../assets/img/icons/009-message.png';
/** thanks to https://gist.github.com/keeguon/2310008 */
import countries from "./countries"
const MASTER_URL = process.env.REACT_APP_MASTER_URL
const INTERESTS = [
    {id: 1, name: 'just bored', tag: 'social'},
    {id: 2, name: 'programming', tag: 'technology'},
    {id: 3, name: 'science', tag: 'technology'},
    {id: 4, name: 'music', tag: 'art'},
    {id: 5, name: 'art', tag: 'art'},
    {id: 6, name: 'politic', tag: 'social'},
]

function WithNavigate(props) {
    let navigate = useNavigate();
    return <LoginForm {...props} navigate={navigate} />
}

class LoginForm extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            countries: countries, isLoaded: true, username: '', country: '', interests: [], profile: '', page: 0, totalPage: 2,
            rememberme: false, error_message: null};
    }
    

    componentDidMount = () => {
        // axios.get(`${MASTER_URL}/master/products`)
        // .then(response => {
        //     let result = response.data
        //     let newState = {}
        //     result.map(rs => {
        //         this.products_price[rs.name] = rs.price
        //         newState[rs.name] = -1
        //         return rs
        //     })
        //     this.setState({
        //         isLoaded: true,
        //         products: result,
        //         ...newState
        //     });
        // }, (error) => {
        //     this.setState({
        //         isLoaded: true,
        //         error
        //     });
        // }).catch(function (error) {
        //     console.log(error);
        // })
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
        this.setState({[name]: value})
    }
  
    handleSubmit = (event) => {
        event.preventDefault();

        // const phoneIndoRegex = /\+?([ -]?\d+)+|\(\d+\)([ -]\d+)/
        const nameExist = this.state.username && this.state.username.length
        if(!nameExist) this.setState({error_message: 'Username cannot be null!\n'})
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
            })
            .then(response => {
                if(response instanceof Error) throw response
                sessionStorage.setItem('token', response.data.token);
                localStorage.setItem('userinfo', JSON.stringify(response.data.data));
                this.props.navigate(`/dashboard/${response.data.data.id}`)
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
            if(this.state.username && this.state.username.length) {
                this.setState({error_message: ''})
                this.setState({page: val})
            } else {
                this.setState({error_message: 'Username cannot be null!\n'})
            }
        } else {
            this.setState({page: val})
            this.setState({error_message: ''})
        }
    }

    confirmOrder = () => {
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
            <div className="ds-fade-in">
                <h5 className="txt-desc ds-m-4 ds-mt-5" style={style.bluecolor}>
                    Quick profile for chatting :
                </h5>
                <div style={style.flex}>
                    <Input name="username" value={this.state.username} 
                        placeholder="Username, How people will address you?" 
                        onChange={this.handleChange} required={true}/>
                    <select id="country">
                        { countryList.map((ctr, idx) => {
                            return (
                                <option className="ds-option" key={idx} onChange={this.handleChange} value={ctr.name}>{ctr.name}</option>
                                )
                            })}
                    </select>
                </div>
                <div className="ds-m-6" style={{color: 'red'}}>{this.state.error_message}</div>
                <button className="ds-m-2 ds-mb-5" style={style.buttonSolid} type="button" onClick={() => this.changePage(1)}>Next</button>
                <hr/>

                <h5 className="txt-topic ds-m-4 ds-mt-4" style={style.bluecolor}>
                    Or :
                </h5>
                <div className="ds-mt-5 ds-mb-5" style={style.flexRow}>
                    <div className="ds-p-1" target="_blank" rel="noreferrer">
                        <img className="logo" alt="telegram" src={tele} title="click to open"></img>
                        <div type="button" style={style.bluecolor}>Login Via Telegram</div>
                    </div>
                    <div className="ds-p-1 ds-ml-4" target="_blank" rel="noreferrer">
                        <img className="logo" alt="email" src={mail} title="click to open"></img>
                        <div type="button" style={style.bluecolor}>Login Via Email</div>
                    </div>
                </div>
            </div>
        )

        const page2 = (
            <div className="ds-fade-in">
                <div className="txt-desc ds-mt-4 ds-mb-4 " style={style.bluecolor} htmlFor="country">
                    What interest you:
                </div>
                { INTERESTS.map(int => {
                    return (
                        <div key={int.id} className="ds-mt-1" style={style.flexRow}>
                            <input id={int.id+int.name} name="interests" type="checkbox" className="ds-mr-5"
                                value={int.name} onChange={this.handleChange} />
                            <label className="txt-desc" style={style.bluecolor} htmlFor={int.id+int.name}>{int.name}</label>
                        </div>
                    )
                })}
                <Input label="Profile" name="profile" value={this.state.profile} 
                type="textarea"
                placeholder="Add some details so people can know you better *optional though" 
                onChange={this.handleChange}/>
                <button className="ds-m-2 ds-mb-5" style={style.buttonClear} type="button" onClick={() => this.changePage(0)}>Back</button>
                <button className="ds-m-2 ds-mb-5" style={style.buttonSolid} type="submit" value="Submit">Submit</button>
            </div>
        )

        if(error) return <div>Error: {error.message} </div>;
        else if(!isLoaded) return <div>Loading...</div>;
        else 
        return (
            <div>
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