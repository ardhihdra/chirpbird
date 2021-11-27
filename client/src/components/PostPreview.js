import React from "react"

export default class PostPreview extends React.Component {
    render() {
        return (
            <div className="ds-border ds-align-left preview-block">
                <div className="ds-row">
                    <div className="ds-col-12 prev-highlight ds-m-1">
                        <img className="prev-highlight ds-border" src={this.props.data.highlight} alt="highlight" />
                    </div>
                    <div className="ds-col-12 txt-subheader ds-m-1 ds-ml-1 prev-title">{this.props.data.title}</div>
                    <div className="ds-col-12 prev-desc txt-desc ds-ml-1">
                        {this.props.data.description} ...Read More
                    </div>
                    <div className="ds-col-12 prev-info ds-ml-1 ds-mt-1">{this.props.data.date}, {this.props.data.situation}</div>
                </div>
            </div>
        )
    }
}