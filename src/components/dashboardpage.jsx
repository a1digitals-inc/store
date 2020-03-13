import React from "react";

class DashboardPage extends React.Component {
    render() {
        return (
            <div className="container-fluid text-center">
                <h1 class="m-5">Dashboard</h1>
                <ul className="list-unstyled mt-5"> 
                    <li className="py-1"><a className="h5" href="/dashboard/products">PRODUCTS</a></li>
                    <li className="py-1"><a className="h5" href="/dashboard/orders">ORDERS</a></li>
                    <li className="py-1"><a className="h5" href="/dashboard/settings">SETTINGS</a></li>
                </ul>
            </div>
        );
    }
}

export default DashboardPage;
