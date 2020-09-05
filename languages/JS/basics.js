// alert. propmt, confirm
/* Комментарий 2
*/
document.write('<h1>Учим Javascript</h1>');

var a = '20', b = 20, c = 30;
var dict = {key1: 111, 'key 2': [1, true, {}, 'string'],};

if (a === b) {    
    document.write('<p>Сработал if</p>');
} else if (a == b) {
    document.write('<p>Сработал else if</p>');
} else {
    document.write('<p>Сработал else</p>');
};

document.write('keys: ', Object.keys(dict), '<br>');
document.write('values: ', Object.values(dict), '<br>');
document.write('value 1: ', dict.key1, '<br>');
document.write('value 2: ', dict['key 2'], '<br>');

var b = '1';
var c = 1;

while (b == c) {
    prompt('Цикл while сработал');
    c++;
};

for (var i = 0; i < 5; i += 2) {
    confirm('Цикл for сработал');
};

function name (arg) {
    // тело функции
};

var name = function (arg) {
    // тело функции
};