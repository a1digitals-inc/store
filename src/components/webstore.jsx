import React from "react";

import StoreHeader from "./storeheader.jsx";
import ProductList from "./productlist.jsx";


class WebStore extends React.Component {
    render() {
        return (
            <div className="container text-center">
                <StoreHeader link={"/"}></StoreHeader>
                <ProductList></ProductList>
            </div>
        );
    }
}

export default WebStore;
