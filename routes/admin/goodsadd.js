var res = require('./lib/resp.js');
var cates = require("./models/cates.js");
var goods = require("./models/goods.js");
var filepath = require('path/filepath');
var dbUtils = require('./models/dbUtils.js');
var io = require('io');
var os = require('os');
var strings = require('strings');
var Dirs = require("./lib/Dirs.js");
var utils = require('./lib/utils.js');


if (!session.get("isAdmin")) {
  return res.redirect('/routes/admin/login.js');
}

if (request.getMethod() === 'GET') {
  var allData = cates.find();
  var data = {
    tree: cates.getTree(allData),
  };
  return res.render('./views/admin/goodsadd.ejs', data);
}

if (request.getMethod() === 'POST') {
  var err = request.parseMultipartForm(1024 * 1024)
  if (err) {
    throw err
  }

  var data = {
    oid: dbUtils.makeOid(),
    goods_name: request.formValue('goods_name'),
    cat_id: request.formValue('cat_id'),
    shop_price: request.formValue('shop_price'),
    goods_desc: request.formValue('goods_desc'),
    goods_number: utils.toInt(request.formValue('goods_number'), 0),
    is_best: utils.toInt(request.formValue('is_best'), 0),
    is_new: utils.toInt(request.formValue('is_new'), 0),
    is_hot: utils.toInt(request.formValue('is_hot'), 0),
    is_on_sale: utils.toInt(request.formValue('is_on_sale'), 0),
  };

  var o = request.formFile('goods_img');
  if (o.err) {
    if (request.isMissingFile(o.err)) {
      res.send("未上传图片");
      return;
    }
    throw o.err;
  }
  var contentType = o.header['Content-Type'][0];
  if (!strings.hasPrefix(contentType, "image")) {
    res.send("上传的不是图片");
    return;
  }
  var ext = filepath.ext(o.name);
  var oidFilename = dbUtils.makeOid() + ext;
  var dateDir = Dirs.getDateDir();
  // 文件的本地磁盘路径
  var filename = filepath.join(dateDir.fullpath, oidFilename);
  // web访问路径
  var publicPath = dateDir.dir + oidFilename;

  var r = os.openFile(filename, os.O_CREATE | os.O_WRONLY, 0666)
  if (r.err) {
    throw r.err;
  }
  var file = r.value
  r = io.copy(file, o.file)
  if (r.err) {
    throw r.err;
  }
  o.file.close()
  file.close()

  data.goods_img = publicPath;
  // res.json(data)
  console.log("%j", data)
  goods.add(data);

  res.redirect('./goodslist.js');
}
