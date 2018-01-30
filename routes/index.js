var res = require('./lib/resp.js');
var cates = require("./models/cates.js");
var goods = require("./models/goods.js");
var session = require('./lib/session.js');

if (request.getMethod() === 'GET') {
  var data = {
    history: session.get('history') ? JSON.parse(session.get('history')) : [],

    tree: cates.getTree(cates.find()),
    bestGs: goods.find({ is_best: 1, is_on_sale: 1 }, { sort: { id: 'desc' }, skip: 0, limit: 3 }),
    newGs: goods.find({ is_new: 1, is_on_sale: 1 }, { sort: { id: 'desc' }, skip: 0, limit: 3 }),
    hotGs: goods.find({ is_hot: 1, is_on_sale: 1 }, { sort: { id: 'desc' }, skip: 0, limit: 3 }),

    user: session.getUser(),
    cart: session.getCart(),
  };
  res.render('./views/index.ejs', data);
}

