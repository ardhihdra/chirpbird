import React from 'react';
import './Input.css';
export default class Input extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: props.value}
        this.handleChangeParent = this.props.onChange
    }
       
    handleChange = (e) => {
        this.setState({value: e.target.value})
        this.handleChangeParent(e)
    }

    styling = () => {
        return {
            label: {
                fontSize: 'smaller',
                position: 'relative',
                top: this.props.type !== 'textarea' ?'-3rem': '-8.5rem',
                left: '-30%',
                width: 'fit-content',
                backgroundColor: this.props.label ? 'white': 'transparent',
                borderRadius: 'var(--border-radius-sm)',
                padding: '1px 2px',
                color: 'var(--primary-font)',
            },
            smallinput: {
                width: '80px',
                marginLeft: 'auto',
                marginRight: 'auto'
            }
        }
    }

    render() {
        const value = this.state.value
        const name = this.props.name
        const label = this.props.label
        const required = this.props.required
        const type = this.props.type || 'text'
        const min = type === 'number' ? this.props.min: null
        const max = type === 'number' ? this.props.max: null 
        const placeholder = this.props.placeholder
        const style = this.styling()
        const divClass = `${this.props.className} ds-ml-3 ds-p-1`
        return (
            <div className="input">
                { type !== 'textarea' ? 
                    <input className={divClass} name={name} type={type} value={value} 
                        min={min} max={max} style={type === 'number' ? style.smallinput: {}}
                        placeholder={placeholder} onChange={this.handleChange} required={required}/>
                :
                    <textarea className="ds-p-2 ds-mt-4" name={name} type={type} value={value} 
                    min={min} max={max} style={type === 'number' ? style.smallinput: {}}
                    placeholder={placeholder} onChange={this.handleChange} required={required}/>
                }
                <span className={this.props.className} style={style.label}>{label}</span>
            </div>
        )
    }
}