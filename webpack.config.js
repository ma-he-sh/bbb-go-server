var path = require("path");
var webpack = require("webpack");

module.exports = {
    entry: {
        'frontend': './public_src/script/frontend.ts',
        'backend': './public_src/script/backend.ts',
    },
    stats: { warnings:false },
    devtool: 'none',
    module: {
        rules: [
            {
                test: /\.ts(x)?$/,
                use: 'ts-loader',
                exclude: /node_modules/
            },
            {
                test: /\.scss$/,
                use: [
                    { loader: "style-loader" },
                    { loader: "css-loader" },
                    { loader: "sass-loader" }
                ]
            },
            {
                test: /\.svg$/,
                loader: 'svg-inline-loader'
            },
            {
                enforce: "pre",
                test: /\.js$/,
                loader: "source-map-loader"
            }
        ]
    },
    resolve: {
        extensions: ['.tsx', '.ts', '.js', '.json'],
    },
    node: {
        globals:
        fs: 'empty',
    },
    output: {
        path: path.resolve(__dirname, 'public/script'),
        filename: "[name].js"
    },
    target: 'web',
    optimization: {
        //minify settings
        minimize: true,
    },
    // change the mod to production : when export : to dev : development
    mode: 'production',
    //live compile
    watch: true
};
