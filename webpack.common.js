const path = require("path");

module.exports = {
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /\.css$/,
                use: ["style-loader", "css-loader"]
            }
        ]
    },
    entry: {
        home: "./src/home.jsx",
        store: "./src/cart.jsx",
        product: "./src/product.jsx",
        info: "./src/info.jsx",
        store: "./src/store.jsx",
        login: "./src/login.jsx",
        dashboard: "./src/dashboard.jsx",
        adminproducts: "./src/adminproducts.jsx"
    },
    output: {
        path: path.resolve(__dirname + "/static"),
        filename: "[name].bundle.js",
        chunkFilename: "common.bundle.js"
    },
    optimization: {
        splitChunks: {
            chunks: "all",
        },
    },
};
