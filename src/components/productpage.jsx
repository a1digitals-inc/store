import React from "react";

import StoreHeader from "./storeheader.jsx";
import ProductCarousel from "./productcarousel.jsx";


class ProductPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            id: window.location.pathname.split("/").pop(),
            data: []
        };
    }

    componentDidMount() {
        fetch("/api/product/" + this.state.id).then(res => res.json()).then(
            (result) => {
                console.log(result);
                document.title = result.name;
                this.setState({
                    isLoaded: true,
                    data: result
                });
            },
            (error) => {
                this.setState({
                    isLoaded: true,
                    error
                });
            }
        )
    }

    render() {
        const { error, isLoaded, data} = this.state;
        if (error) {
            return <div>Error: {error.message}</div>;
        } else if (!isLoaded || data == null) {
            return (
                <div className="container text-center">
                <StoreHeader link="/webstore"></StoreHeader>
                    <div className="spinner-border"></div>
                </div>
            );
        } else {
            return (
                <div className="container text-center">
                    <StoreHeader link="/webstore"></StoreHeader>
                    <div className="row justify-content-center fade-in">
                        <div className="col-sm">
                            <ProductCarousel images={data.images}></ProductCarousel> 
                        </div>
                        <div className="col-sm">
                            <h2>{data.name}</h2>
                            <div className="my-3">
                            {data.description.split('\n').map(text => (<p className="my-1" key={text}>{text}</p>))}
                            </div>
                            {data.discount * data.price < data.price ? (
                                <div>
                                    <p><del>{data.price} $</del></p>
                                    <p><strong>{+(Math.round(data.discount * data.price + "e+2") + "e-2")} $</strong></p>
                                </div>
                            ) : (<p><strong>{data.price} $</strong></p>)}
                            {data.options && data.options.length > 0 &&
                            <select className="form-control">
                                {data.options.map(option => (
                                    <option key={option}>{option}</option>
                                ))}
                            </select>}
                            {data.options ? <button type="button" className="btn btn-dark my-3">ADD TO CART</button> : <button type="button" className="btn btn-dark my-3" disabled>SOLDOUT</button>}
                        </div>
                    </div>
                </div>
            );            
        }
    }
}

export default ProductPage;
