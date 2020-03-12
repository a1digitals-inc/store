import React from "react";

class ProductCarousel extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            active: 0
        };
    }
    
    componentDidMount() {
        setInterval(() => {
            this.setState({
                active: this.state.active+1
            });
            if (this.state.active >= this.props.images.length) {
                this.setState({
                    active: 0
                });
            }
        }, 3000)
    }

    render() {
        return <img className="img-fluid fadet" src={this.props.images[this.state.active]} />;
    }
}

export default ProductCarousel;
