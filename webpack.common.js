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
        store: "./src/store.jsx"
    },
    output: {
        path: path.resolve(__dirname + "/static"),
        filename: "[name].bundle.js",
    }
};
