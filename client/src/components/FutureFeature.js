import React from "react"
const dsFeature = {
  'background-color': 'rgb(190, 190, 4)',
};

const title = {
    'color': 'black', 
    'font-weight': 'bold'
}
// .ds-future-feature 

export default class FutureFeature extends React.Component {
    constructor() {
        super();
        this.message = ''
    }

    render() {
        const description = this.props.description || 'Future updates!';
        return (
            <div className="ds-m-2 mr-5 ml-5 p-4 rounded-lg" style={dsFeature}>
                {/* Side Menu */}
                <p style={title}>Future Feature {this.message}</p>
                <p>{description}</p>
            </div>
        )
    }
}