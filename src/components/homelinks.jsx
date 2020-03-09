import React from "react";

class HomeLinks extends React.Component {
    render() {
        return (
            <div className="container-fluid text-center">
                <h1 className="m-5">Store</h1>
                <ul className="list-unstyled mt-5">
                    <li className="py-1"><a class="h5" href="/webstore">STORE</a></li>
                    <li className="py-1"><a class="h5" href="/media">MEDIA</a></li>
                    <li className="py-1"><a class="h5" href="/info">INFO</a></li>
                </ul>
            </div>
        );
    }
}

export default HomeLinks;
