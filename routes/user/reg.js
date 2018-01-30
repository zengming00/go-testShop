var res = require('./lib/resp.js');
var Verify = require('./lib/Verify.js');
var session = require('./lib/session.js');
var utils = require('./lib/utils.js');
var Common = require('./lib/Common.js');
var users = require("./models/users.js");
var dbUtils = require("./models/dbUtils.js");


if (request.getMethod() === 'GET') {
  var data = {
    tree: [],
    user: session.getUser(),
    cart: session.getCart(),
  }
  return res.render('./views/user/reg.ejs', data);
}

if (request.getMethod() === 'POST') {
  var yzm = request.formValue('yzm');
  if (!Verify.verify(yzm)) {
    console.log('验证码错误');
    return res.json({ error: '验证码错误' });
  }

  var $username = request.formValue('username');
  var $email = request.formValue('email');
  var $password = request.formValue('password');
  var $repassword = request.formValue('repassword');
  var $phone = request.formValue('phone');

  if (!/^(13[0-9]|14[5|7]|15\d|18\d)\d{8}$/.test($phone)) {
    return res.json({ error: "手机号无法接受" });
  }
  if (!/^[a-zA-Z0-9]{6,16}$/.test($username) || $username === 'admin') {//直接拒绝admin账号
    return res.json({ error: "用户名只能是6-16位英文字母+数字" });
  }
  if (!/^[a-zA-z0-9]+@[a-zA-z0-9]+(\.[a-zA-z0-9]+)+$/.test($email)) {
    return res.json({ error: "邮箱无法接受" });
  }
  if (!/^[a-zA-Z0-9]{6,16}$/.test($password)) {
    return res.json({ error: "密码只能是6-16位英文字母+数字" });
  } else if ($password !== $repassword) {
    return res.json({ error: "两次密码不一致" });
  }

  // //看用户名、手机、邮箱是否已存在
  var cond = { $or: [{ userName: $username }, { phone: $phone }, { email: $email }] };
  var r = users.find(cond, {});
  console.log('%j', r);
  if (r.length == 0) {//未找到返回空数组，此时才可以入库
    var $salt = Common.makeSalt(6);
    var data = {
      oid: dbUtils.makeOid(),
      userName: $username,
      phone: $phone,
      email: $email,
      password: utils.encodePassword($password, $salt),
      salt: $salt,
    };
    data = users.add(data);
    console.log('%j', data);
    res.json({ success: '注册成功！' });
  } else {
    res.json({ error: '注册失败！请更换用户名、手机或邮箱' });
  }

}