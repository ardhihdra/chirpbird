import React from "react"

export default class PostPreview extends React.Component {
    render() {
        return (
            <div className="ds-border ds-align-left preview-block">
                <div className="ds-row">
                    <div className="ds-col-12 prev-highlight m-1">
                        <img className="prev-highlight rounded-lg" src={this.props.data.highlight} alt="highlight" />
                    </div>
                    <div className="ds-col-12 txt-subheader m-1 ml-1 prev-title">{this.props.data.title}</div>
                    <div className="ds-col-12 prev-desc txt-desc ml-1">
                        {this.props.data.description} ...Read More
                    </div>
                    <div className="ds-col-12 prev-info ml-1 mt-1">{this.props.data.date}, {this.props.data.situation}</div>
                </div>
            </div>
        )
    }
}