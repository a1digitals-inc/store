import React from "react";
import ReactDOM from "react-dom";

import "bootstrap/dist/css/bootstrap.min.css";
import "./main.css";

import DashboardPage from "./components/dashboardpage.jsx";

ReactDOM.render(
    <DashboardPage></DashboardPage>,
    document.getElementById("app")
);
