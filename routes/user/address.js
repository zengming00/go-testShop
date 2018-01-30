var cates = require("./models/cates.js");
var res = require('./lib/resp.js');
var session = require('./lib/session.js');
var users = require("./models/users.js");


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
  return res.render('./views/user/address.ejs', data);
}

if (request.getMethod() === 'POST') {
  var address = request.formValue('address');
  var name = request.formValue('name');
  var phone = request.formValue('phone');
  var addr = (address + '  ' + name + '  ' + phone).trim();

  var r = users.updateByOid(user.oid, { address: addr });
  console.log('%j', r);
  user.address = addr;
  session.set('user', JSON.stringify(user));
  res.redirect('./address.js');
}