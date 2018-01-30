var cates = require("./models/cates.js");
var ordinfos = require("./models/ordinfos.js");
var res = require('./lib/resp.js');
var session = require('./lib/session.js');
var Common = require('./lib/Common.js');


var user = session.getUser();
if (!user) {
  return res.redirect('/routes/user/login.js');
}

if (request.getMethod() === 'GET') {
  var allData = cates.find();
  var data = {
    tree: cates.getTree(allData),
    user: user,
    cart: session.getCart(),
    ords: ordinfos.find({ userId: user.oid }, { sort: { id: 'desc' } }),
    toDateStr: Common.toDateStr,
  };
  res.render('./views/user/ordlist.ejs', data);
}