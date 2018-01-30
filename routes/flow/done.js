var cates = require("./models/cates.js");
var goods = require("./models/goods.js");
var ordgoods = require("./models/ordgoods.js");
var ordinfos = require("./models/ordinfos.js");
var dbUtils = require("./models/dbUtils.js");
var res = require('./lib/resp.js');
var session = require('./lib/session.js');
var comm = require('./lib/Common.js');


var user = session.getUser();
if (!user) {
  return res.redirect('/routes/user/login.js');
}

if (request.getMethod() === 'POST') {
  var cart = session.getCart();

  var date = comm.getCnDate();
  var ordId = '' + date.getFullYear() + (date.getMonth() + 1) + date.getDate()
    + date.getHours() + date.getMinutes() + date.getSeconds() + Math.random().toString().substr(5);

  var carts = cart.items();
  if (carts.length == 0) {
    res.send("没有购买商品");
    return;
  }
  for (var i = 0; i < carts.length; i++) {
    var gs = carts[i];
    if (gs.num > 0) {
      var ordgood = {
        oid: dbUtils.makeOid(),
        ordId: ordId,
        goodsId: gs.id,
        goodsName: gs.name,
        price: gs.price,
        num: gs.num
      };
      var r = goods.decrGoodsNum(ordgood.goodsId, ordgood.num);
      if (r.rowsAffected === 0) {
        return res.send('下单失败！    ' + err.goodsName + "   商品不存在或库存不足！");
      }
      ordgoods.add(ordgood);
    }
  }
  var ordinfo = {
    oid: dbUtils.makeOid(),
    ordId: ordId,
    userId: user.oid,
    userName: user.userName,
    address: user.address,
    payType: 'RMB',
    payState: false,
    money: cart.getTotalMoney(),
    fuyan: request.formValue('fuyan'),
  };
  ordinfos.add(ordinfo);

  cart.clear();
  session.setCart(cart);
  var allData = cates.find();
  res.render('./views/flow/done.ejs', {
    ordId: ordId,
    user: user,
    cart: cart,
    tree: cates.getTree(allData),
    money: ordinfo.money,
    payForm: comm.getPayForm(ordId, ordinfo.money)
  });
}
