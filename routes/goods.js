var res = require('./lib/resp.js');
var cates = require("./models/cates.js");
var goods = require("./models/goods.js");
var dbUtils = require('./models/dbUtils.js');
var utils = require('./lib/utils.js');
var Page = require('./lib/Page.class.js');
var Common = require('./lib/Common.js');
var session = require('./lib/session.js');


if (request.getMethod() === 'GET') {
  var query = utils.getQuery();
  var gs = goods.getByOid(query.oid);
  var cats = cates.find();

  //添加到历史记录
  var gsData = {
    oid: gs.oid,
    goods_img: gs.goods_img,
    goods_name: gs.goods_name,
    shop_price: gs.shop_price
  };
  var history = session.get('history') ? JSON.parse(session.get('history')) : [];
  var temp = [];
  for (var i = 0; i < history.length; i++) { //过滤重复数据
    if (history[i]._id != gsData._id) {
      temp.push(history[i]);
    }
  }
  temp.unshift(gsData);//插入到前面
  if (temp.length > 5) {
    temp.pop();
  }
  session.set('history', JSON.stringify(history));

  //渲染
  var data = {
    gs: gs,
    tree: cates.getTree(cats),
    family: cates.getFamily(cats, gs.cat_id),
    user: session.getUser(),
    cart: session.getCart(),
  };
  if (query.json) {
    res.json(data);
    return;
  }
  res.render('./views/goods.ejs', data);
}

