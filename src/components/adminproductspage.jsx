import React from "react";

class AdminProductsPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            items: []
        }
    }
    
    componentDidMount() {
        fetch("/api/products")
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        items: result
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
        const { error, isLoaded, items } = this.state;
        if (error) {
            return <div>Error: {error.message}</div>
        } else if (!isLoaded || items == null) {
            return (
                <div className="container-fluid text-center"> 
                    <h1 className="m-5"><a href="/dashboard">Dashboard</a></h1>
                    <div className="spinner-border"></div>
                </div>
                );
        } else {
            return (
                <div className="container-fluid text-center">
                    <h1 className="m-5"><a href="/dashboard">Dashboard</a></h1>
                    <ul className="list-group">
                    {items.map(product => (
                        <li className="fade-in list-group-item">
                            <img className="float-left mini-image" src={product.thumbnail} />
                            {product.name}
                            <a className="float-right btn btn-primary" href={"/dashboard/product/" + product.identifier}>Edit</a>
                            {product.soldout && <p><strong>Soldout</strong></p>}
                        </li>
                    ))}
                    </ul>
                    <a className="float-right btn btn-primary m-3" href="/dashboard/product">Add</a>
                </div>
            );
        }
    }
}

export default AdminProductsPage;
