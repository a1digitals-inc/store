import React from "react";

class ProductCarousel extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            active: 0
        };
    }
    
    componentDidMount() {
        if (this.props.images) {
            setInterval(() => {
                this.setState({
                    active: this.state.active+1
                });
                if (this.state.active >= this.props.images.length) {
                    this.setState({
                        active: 0
                    });
                }
            }, 3000);
        } 
    }

    render() {
        return <img className="img-fluid fade-in" src={this.props.images && this.props.images.length > 0 ? this.props.images[this.state.active] : undefined} />;
    }
}

export default ProductCarousel;
