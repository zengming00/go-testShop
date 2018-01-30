var sql = require('sql');

function getConn() {
  var reUse = true;
  var r = sql.open("sqlite3", "./models/test.db", reUse);
  if (r.err) {
    throw r.err;
  }
  var db = r.value;
  if (!r.isReUse) {
    db.setMaxOpenConns(100);
  }
  return db;
}

exports.getConn = getConn;


