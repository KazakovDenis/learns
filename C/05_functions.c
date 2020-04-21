#include <stdio.h>

/* Синтаксис: тип_возвращаемого_значения имя_функции (тип_параметра1 параметр1, тип_параметраN параметрN)
Совокупность принимаемых и возвращаемых параметров с их типами определяют сигнатуру функции

тип void - функция ничего не возвращает, в аргументах - не принимает
*/
void someFunction(void){}   // ; не требуется


// Изменение переменной внутри функции не влияет на переменную, переданную в аргументах
float a = 1.55;
int get_int(float param){
    return (int)param+5;
}

// Но при работе с массивами - изменение массива вызовет его изменение глобально
// т.к. имя массива - это указатель
const int SIZE = 5;
int array[];

// функция ничего не возвращает, но заполнит массив array[]
int fill_array1(int arr[], int arr_size){
    for (int i = 0; i < arr_size; i++){
        arr[i] = i;
    }
}

int fill_array2(int *arr, int arr_size){
    for (int i = 0; i < arr_size; i++){
        arr[i] = i;
    }
}


// Прототип функции, позволяет вызвать функцию, определённую ниже вызова
// как правило, его помещают в заголовки
void proto(void);


int main(void)
{
    int b = get_int(a);
    printf("a = %.2f, b = %d\n", a, b);

    // заполняем массив передачей через []
    fill_array1(array, SIZE);

    for (int i = 0; i < SIZE; i++){
        printf("%d ", array[i]);
    };

    printf("\n");

    // заполняем массив через указатель
    fill_array2(array, SIZE);

    for (int i = 0; i < SIZE; i++){
        printf("%d ", array[i]);
    };

    printf("\n");
    proto();
	return 0;
}

void proto(){
    printf("Это прототип\n");
}

/* Встроенные функции

// Ввод
char a = getchar();
char b[100] = gets(b);
float c, d;
scanf("%f %f", &c, &d);   // & - адресный оператор, см. pointers.c

// Вывод
putchar(a);
puts(b);

*/