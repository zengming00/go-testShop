var cates = require('./models/bak/cates.js');
var goods = require('./models/bak/goods.js');
var db = require('./models/db.js');
var dbUtils = require('./models/dbUtils.js');

var db = db.getConn();

function createCates(db) {
  var ddl = 'CREATE TABLE IF NOT EXISTS [cates] ('
    + '[id] INTEGER PRIMARY KEY AUTOINCREMENT,'
    + '[oid] TEXT,'
    + '[cat_name] TEXT,'
    + '[intro] TEXT,'
    + '[parent_id] TEXT,'
    + '[created_at] DATETIME DEFAULT CURRENT_TIMESTAMP'
    + ');';

  r = db.exec(ddl);
  if (r.err) {
    throw r.err;
  }
  console.log('create cates success!');
}

function insertIntoCates(db, oid, cat_name, intro, parent_id) {
  var sql = "insert into cates(oid,cat_name,intro,parent_id) values(?,?,?,?)";
  return dbUtils.insert(db, sql, [oid, cat_name, intro, parent_id]);
}


function initCatesData(db) {
  cates.forEach(function (item) {
    var r = insertIntoCates(db, item._id.$oid, item.cat_name, item.intro, item.parent_id)
    console.log("%j", r)
  })
}


function createGoods(db) {
  var ddl = 'CREATE TABLE IF NOT EXISTS [goods] ('
    + '[id] INTEGER PRIMARY KEY AUTOINCREMENT,'
    + '[oid] TEXT,'
    + '[goods_name] TEXT,'
    + '[cat_id] TEXT,'
    + '[shop_price] INTEGER,'
    + '[goods_img] TEXT,'
    + '[goods_desc] TEXT,'
    + '[goods_number] INTEGER,'
    + '[is_best] INTEGER,'
    + '[is_new] INTEGER,'
    + '[is_hot] INTEGER,'
    + '[is_on_sale] INTEGER,'
    + '[created_at] DATETIME DEFAULT CURRENT_TIMESTAMP,'
    + '[updated_at] DATETIME DEFAULT CURRENT_TIMESTAMP'
    + ');';


  r = db.exec(ddl);
  if (r.err) {
    throw r.err;
  }
  console.log('create goods success!');
}

function insertIntoGoods(db, oid, goods_img, goods_name, cat_id, shop_price, goods_desc, goods_number, is_on_sale, is_hot, is_new, is_best) {
  var sql = "insert into goods(oid, goods_img, goods_name, cat_id, shop_price, goods_desc, goods_number, is_on_sale, is_hot, is_new, is_best) values(?,?,?,?,?,?,?,?,?,?,?)";
  return dbUtils.insert(db, sql, [oid, goods_img, goods_name, cat_id, shop_price, goods_desc, goods_number, is_on_sale, is_hot, is_new, is_best]);
}

function initGoodsData(db) {
  goods.forEach(function (item) {
    var r = insertIntoGoods(db, item._id.$oid, item.goods_img, item.goods_name, item.cat_id, item.shop_price, item.goods_desc, item.goods_number, item.is_on_sale, item.is_hot, item.is_new, item.is_best);
    console.log("%j", r)
  })
}

function createOrdgoods(db) {
  var ddl = 'CREATE TABLE IF NOT EXISTS [ordgoods] ('
    + '[id] INTEGER PRIMARY KEY AUTOINCREMENT,'
    + '[oid] TEXT,'
    + '[ordId] TEXT,'
    + '[goodsId] TEXT,'
    + '[goodsName] TEXT,'
    + '[price] TEXT,'
    + '[num] INTEGER,'
    + '[created_at] DATETIME DEFAULT CURRENT_TIMESTAMP'
    + ');';

  r = db.exec(ddl);
  if (r.err) {
    throw r.err;
  }
  console.log('create ordgoods success!');
}

function createUsers(db) {
  var ddl = 'CREATE TABLE IF NOT EXISTS [users] ('
    + '[id] INTEGER PRIMARY KEY AUTOINCREMENT,'
    + '[oid] TEXT,'
    + '[userName] TEXT,'
    + '[phone] TEXT,'
    + '[email] TEXT,'
    + '[password] TEXT,'
    + '[salt] TEXT,'
    + '[address] TEXT,'
    + '[created_at] DATETIME DEFAULT CURRENT_TIMESTAMP'
    + ');';

  r = db.exec(ddl);
  if (r.err) {
    throw r.err;
  }
  console.log('create users success!');
}

function createOrdinfos(db) {
  var ddl = 'CREATE TABLE IF NOT EXISTS [ordinfos] ('
    + '[id] INTEGER PRIMARY KEY AUTOINCREMENT,'
    + '[oid] TEXT,'
    + '[ordId] TEXT,'
    + '[userId] TEXT,'
    + '[userName] TEXT,'
    + '[address] TEXT,'
    + '[payType] TEXT,'
    + '[payState] INTEGER,'
    + '[money] TEXT,'
    + '[fuyan] TEXT,'
    + '[created_at] DATETIME DEFAULT CURRENT_TIMESTAMP'
    + ');';

  r = db.exec(ddl);
  if (r.err) {
    throw r.err;
  }
  console.log('create ordinfos success!');
}


var stat = db.stats();
console.log("%j", stat)

createCates(db);
initCatesData(db);

createGoods(db);
initGoodsData(db);

createUsers(db)
createOrdinfos(db)
createOrdgoods(db)

console.log('done.');