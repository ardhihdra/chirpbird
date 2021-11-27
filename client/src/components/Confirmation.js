import React from 'react';

export default class Confirmation extends React.Component {
    constructor(props) {
        super(props);
        this.state = {isShow: props.isShow}
    }
       
    handleChange = (e) => {
        this.setState({isShow: e.target.isShow})
    }

    styling() {
        return {
            backdrop: {
                overflow: 'none',
                visibility: 'hidden',
                display: 'none',
                justifyContent: 'center',
                alignItems: 'center',
                position: 'fixed',
                zIndex: '1400',
                top: 0,
                right: 0,
                bottom: 0,
                left: 0,
                // left: '50vw',
                // top: '50vh',
            },
            modal: {
                width: '40vw',
                borderRadius: 'var(--border-radius-xlg)',
                backgroundColor: 'grey',
                margin: '2rem auto 0',
                minWidth: '320px',
                maxWidth: '600px',
                maxHeight: 'calc(100vh - 150px)',
                position: 'relative',
                display: 'flex',
                flexDirection: 'column'
            },
        }
    }

    render() {
        const value = this.state.isShow
        const text = this.props.text || 'text'
        const style = this.styling()
        return value ? (
            <div style={style.backdrop}>
                <div style={style.modal}>
                    {text}
                </div>
            </div>
        ): ''
    }
}