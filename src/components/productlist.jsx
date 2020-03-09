import React from "react";

class ProductList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            items: []
        };
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
            return <div>Error: {error.message}</div>;
        } else if (!isLoaded) {
            return <div className="spinner-border"></div>
        } else {
            return (
                <ul className="list-inline">
                {items.map(product => (
                    <li className="list-inline-item" style={{width: "178px", height: "178px"}}>
                        <a href={"/product/" + product.id}>
                            <img className="w-100"src={product.thumbnail} />
                            <p>{product.name}</p>
                            {product.soldout && <p><strong>Soldout</strong></p>}
                        </a>
                    </li>
                ))}
                </ul>
            );
        }
    }
}

export default ProductList;
