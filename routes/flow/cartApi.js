var goods = require("./models/goods.js");
var res = require('./lib/resp.js');
var session = require('./lib/session.js');
var utils = require('./lib/utils.js');


if (request.getMethod() === 'GET') {
  var cart = session.getCart();
  var query = utils.getQuery();
  switch (query.a) {
    case 'add':
      var gs = goods.getByOid(query.oid);
      cart.add(gs.oid, gs.goods_name, gs.goods_img, gs.shop_price, 1);
      break;
    case 'del':
      cart.del(query.oid);
      break;
    case 'incr':
      cart.incr(query.oid);
      break;
    case 'decr':
      cart.decr(query.oid);
      break;
    default:
  }
  session.setCart(cart);
  res.redirect('./cart.js');
}

