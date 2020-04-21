// Хороший портал https://prog-cpp.ru/category/c-posts/
// импорт библиотеки (т.н. заголовка) ввода/вывода
#include <stdio.h>
// для импорта своих модулей указывать в кавычках
// #include "path/file.c"


// основная функция main - точка входа, ищется компилятором
int main(){
	printf("Hello world!\n");
	// getchar();   // Задержка окна консоли
	return 0;    // Сообщение системе, что программа закончилась успешно
}


// inline comment
/* multiline comment
*/
/*
Чтобы запустить программу, необходимо скомпилировать код в исполняемый вид
hello_world.c > hello_world.bin, или hello_world.exe, или другой в зависимости от платформы
Пример для linux (можно и без расширения выходного файла):
gcc hello_world.c -o hello_world.bin && ./hello_world.bin
*/