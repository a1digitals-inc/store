import React from "react";
import ReactDOM from "react-dom";

import "bootstrap/dist/css/bootstrap.min.css";
import "./main.css";

import WebStore from "./components/webstore.jsx";

ReactDOM.render(
    <WebStore></WebStore>,
    document.getElementById("app")
);
