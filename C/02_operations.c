#include <stdio.h>

/* Операторы в большинстве своём аналогичны Python:
Сравнение: =, ==, !=, >, >=, <, <=, 
Арифметика: +, -, /, *, %, +=, -=, *=, /=, %= [нет оператора "//" ;)], ++, --
Логические: && = and, || = or, ! = not
Побитовые: & = and, | = or, ~ = not (унарная), ^ - исключающее ИЛИ, а также &=, |=, ^=
Сдвиговые: >>, <<  (Арифметический сдвиг целого числа вправо >> на 1 разряд соответствует делению числа на 2 и наоборот)
*/

int a = 5, b = 1, c = 0;
float d = 1.7;


int main(void){

    // Примеры
    // унарные операции
    a++; ++a; a--; --a;

    // арифметические
    c += a*++b;         // допускается такая запись
    c = a*b--;          // вычитание будет выполнено в конце

    // если тип различается, применяется явное приведение типов
    c = b + (int)d;     // у d будет отброшена дробная часть

    printf("%d \n", c);
    return 0;
}