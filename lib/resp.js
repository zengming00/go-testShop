
// 加载node兼容，给ejs用
require('./lib/mfs.js');

var ejs = require('./lib/ejs.js');
var utils = require("./lib/utils.js");
var fmt = require('fmt');


exports = module.exports = response;

ejs.fileLoader = utils.fileLoader;


exports.redirect = function (path) {
  response.header().set('location', path)
  response.writeHeader(302)
}

exports.json = function (data) {
  response.header().set('Content-Type', 'application/json; charset=utf-8')
  response.write(JSON.stringify(data, null, 2));
}

function html(data) {
  response.header().set('Content-Type', 'text/html; charset=utf-8')
  response.write(data);
}

exports.htmlFile = function (filename) {
  var c = cache.get(filename);
  // if (!c) {
    c = utils.fileLoader(filename)
    cache.set(filename, c);
  // }
  html(c);
}

exports.render = function (view, data) {
  response.header().set('content-type', 'text/html; charset=utf-8')
  var content = utils.fileLoader(view);
  var func = ejs.compile(content, { filename: view });
  var ret = func(data);
  response.write(ret)
}

exports.html = html;
exports.send = response.write;