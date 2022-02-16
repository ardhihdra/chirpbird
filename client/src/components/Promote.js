import React from "react";

export default class Promote extends React.Component {
    render() {
        return (
           <div className="ds-m-1">
                <div className="ds-border p-4 b-orange">
                    <div>I'll tell you my new post immadiately!</div>
                    <input className="ds-mt-3 mb-3 p-3" type="text" placeholder="email address"></input>
                    <button className="ds-mt-1 p-3 ds-btn-yes">SEND ME NEW POST</button>
                </div>
                <div className="ds-container">
                    <div className="ds-col-4"></div>
                    <div className="ds-col-4"></div>
                    <div className="ds-col-4"></div>
                </div>
           </div>
        )
    }
}