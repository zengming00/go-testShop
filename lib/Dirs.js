var os = require('os');
var filepath = require('path/filepath');
var Comm = require('./lib/Common.js');


/**
 * 获取日期文件夹（不存在则自动创建）
 */
function getDateDir() {
    var r = os.getwd()
    if (r.err) {
        throw r.err;
    }
    var baseDir = r.value;
    var date = Comm.getCnDate();
    var dir = '/public/uploads/' + date.getFullYear() + '/' + Comm.getPadStr(date.getMonth() + 1) + '/' + Comm.getPadStr(date.getDate()) + '/';
    var path = filepath.join(baseDir, dir);

    var err = os.mkdirAll(path, 0666);
    if (err) {
        throw err;
    }
    return { dir: dir, fullpath: path };
}


exports.getDateDir = getDateDir;