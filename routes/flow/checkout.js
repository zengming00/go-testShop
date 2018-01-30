var cates = require("./models/cates.js");
var res = require('./lib/resp.js');
var session = require('./lib/session.js');


var user = session.getUser();
if (!user) {
  return res.redirect('/routes/user/login.js');
}

if (request.getMethod() === 'GET') {
  var cart = session.getCart();
  if (cart.items().length == 0) {
    return res.send('购物车为空');
  }
  var allData = cates.find();
  var data = {
    tree: cates.getTree(allData),
    user: user,
    cart: cart,
  };
  res.render('./views/flow/checkout.ejs', data);
}