import React from "react";

import ProductList from "./productlist.jsx"


class WebStore extends React.Component {
    render() {
        return (
            <div className="contaienr-fluid text-center">
                <h1 className="m-5">Store</h1>
                <ProductList></ProductList>
            </div>
        );
    }
}

export default WebStore;
