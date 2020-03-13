import React from "react";

import StoreHeader from "./storeheader.jsx";

class LoginPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            message: ''
        };
        this.login = this.login.bind(this);
    }
    
    login(event) {
        event.preventDefault();
        const data = new FormData(event.target);
        const options = {
            method: "POST",
            body: data
        };
        fetch("/api/login", options)
            .then(res => res.json())
            .then(
                (result) => {
                    console.log(result);
                    if (result.valid) {
                        window.location.href = "/dashboard"; 
                    } else {
                        this.setState({
                            message: "Invalid password"
                        });
                    }
                },
                (error) => {
                    this.setState({
                        message: "Invalid password"
                    });
                }
            )
    }

    render() {
        const { message } = this.state;
        return (
            <div className="container-fluid text-center">
                <StoreHeader link="/"></StoreHeader>
                <form onSubmit={this.login} className="m-5 p-5">
                    {message && <div className="fade-in alert alert-secondary">{message}</div>}
                    <input type="password" className="form-control" name="password" placeholder="Enter password"/>
                    <input type="submit" className="d-none" />
                </form>
            </div>
        );
    }
}

export default LoginPage;
