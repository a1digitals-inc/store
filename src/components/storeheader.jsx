import React from "react";

class StoreHeader extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return <h1 className="m-5"><a href={this.props.link}>Store</a></h1>;
    }
}

export default StoreHeader;
