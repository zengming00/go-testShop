
var cates = require("./models/cates.js");
var Cart = require('./lib/Cart.class.js');
var res = require('./lib/resp.js');
var session = require('./lib/session.js');

if (request.getMethod() === 'GET') {
  var allData = cates.find();
  var data = {
    tree: cates.getTree(allData),
    user: session.getUser(),
    cart: session.getCart(),
  };
  res.render('./views/flow/cart.ejs', data);
}