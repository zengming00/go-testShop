var png = require('image/png');
var capcha = require('./lib/image/capcha.js');


var r = capcha.makeCapcha()
var img = r.img;
var str = r.str;

session.set('__verify', str);

response.header().set('content-type', 'image/png')
var err = png.encode(response, img)
if (err) {
  throw err;
}