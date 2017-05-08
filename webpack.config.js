const webpack = require('webpack')
const path = require('path')

module.exports = {
  entry: [
    './src/index'
  ],
  output: {
    filename: 'bundle.js',
    path: path.join(__dirname, 'dist'),
    publicPath: 'static'
  },
  module: {
    loaders: [{
      test: /\.jsx?$/,
      loaders: ['babel-loader'],
      include: path.join(__dirname, 'src')
    }]
  }
}
