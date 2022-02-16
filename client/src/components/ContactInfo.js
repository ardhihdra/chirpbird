import React from 'react';
import ig from '../assets/img/icons/028-instagram.png';
import wa from '../assets/img/icons/041-whatsapp.png';

export default class Confirmation extends React.Component {
    constructor(props) {
        super(props);
        this.state = {showinfo: false}
        this.wacopy = React.createRef()
        this.copyCodeToClipboard = this.copyCodeToClipboard.bind(this)
    }

    copyCodeToClipboard() {
        const el = this.wacopy
        if(el.current) {
            el.current.select()
            document.execCommand("copy")
            this.setState({showinfo: true})
            setTimeout(() => {
                this.setState({showinfo: false})
            }, 3000)
        }
    }

    render() {
        return (
            <div className="txt-subheader contact-info pt-6">
                <div className="contact-flex">
                    <a href="https://www.instagram.com/kanemufreshdrink.id/"><img className="logo" alt="instagram" src={ig} title="click to open"></img></a>
                    <div className="grow2 ml-3">Follow us :<br/> kanemufreshdrink.id <br/></div>
                    <img onClick={() => this.copyCodeToClipboard()} className="logo ml-2 copyable" alt="whatsapp" src={wa} title="click to copy"></img>
                    <div className="grow2 ml-3">
                        Order & Kemitraan :<br/>
                        <div>+62 8966 2362510 <br/>{this.state.showinfo ? 'Copied!': ''}</div>
                    </div>
                    <input ref={this.wacopy} value="+6289662362510" style={{position: 'fixed', top:'10000px', color: 'transparent', backgroundColor: 'transparent'}} disabled/>
                </div>
                <div className="ds-m-6 hashtag">
                    #menemuKANEMU #KANEMUan #menyegarKANEMU
                </div>
            </div>
        )
    }
}