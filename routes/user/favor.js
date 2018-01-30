var cates = require("./models/cates.js");
var res = require('./lib/resp.js');
var session = require('./lib/session.js');

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
  };
  res.render('./views/user/favor.ejs', data);
}