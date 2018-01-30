var font = require('./lib/image/font.js')
var image = require('./lib/image/image.js')


// 仿PHP的rand函数
function rand(min, max) {
  return Math.random() * (max - min + 1) + min | 0;
}

// 制造验证码图片
function makeCapcha() {
  var img = new image.Image(100, 40);
  img.drawCircle(rand(0, 100), rand(0, 40), rand(10, 40), rand(0, 0xffffff));
  // 边框
  img.drawRect(0, 0, img.w - 1, img.h - 1, rand(0, 0xffffff));
  img.fillRect(rand(0, 100), rand(0, 40), rand(10, 35), rand(10, 35), rand(0, 0xffffff));
  img.drawLine(rand(0, 100), rand(0, 40), rand(0, 100), rand(0, 40), rand(0, 0xffffff));
  // return img;
  // 画曲线
  var w = img.w / 2;
  var h = img.h;
  var color = rand(0, 0xffffff);
  var y1 = rand(-5, 5);
  var w2 = rand(10, 15);
  var h3 = rand(4, 6);
  var bl = rand(1, 5);
  for (var i_1 = -w; i_1 < w; i_1 += 0.1) {
    var y_1 = Math.floor(h / h3 * Math.sin(i_1 / w2) + h / 2 + y1);
    var x_1 = Math.floor(i_1 + w);
    for (var j = 0; j < bl; j++) {
      img.drawPoint(x_1, y_1 + j, color);
    }
  }
  var p = 'ABCDEFGHKMNPQRSTUVWXYZ3456789';
  var str = '';
  for (var i_2 = 0; i_2 < 5; i_2++) {
    str += p.charAt(Math.random() * p.length | 0);
  }
  var fonts = [font.font8x16, font.font12x24, font.font16x32];
  var x = 15, y = 8; // tslint:disable-line
  for (var _i = 0, str_1 = str; _i < str_1.length; _i++) {
    var ch = str_1[_i];
    var f = fonts[Math.random() * fonts.length | 0];
    y = 8 + rand(-10, 10);
    img.drawChar(ch, x, y, f, rand(0, 0xffffff));
    x += f.w + rand(2, 8);
  }
  return {
    img: img.img,
    str: str,
  };
}

exports.makeCapcha = makeCapcha;
