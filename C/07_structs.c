#include <stdio.h>

const float PI = 3.14;
/* Структуры - переменная с полями

- может в качестве поля иметь ссылку на себя же
- может быть определена как тип
- имеет своё пространство имён (структуры видны с ключевым словом struct)

Создаём структуру > создаём экземпляр > инициализируем экземпляр.
Тип является экземпляром структуры, создаётся ключевым словом typedef

Синтаксис:

struct ИмяСтруктуры {
    тип поле1;
    тип поле2, поле3;
    тип *указатель;
    struct ИмяСтруктуры *next;   // указатель на себя
} экземплярСтруктуры;            // может не объявляться

Если экземпляр не создаётся сразу:
struct ИмяСтруктуры экземплярСтруктуры;

*/

// Объявление структуры
struct Object {
    char *string;
    int val2, val3;
    float *pPI;
    struct Object *pointer;
};

// Создание экземпляра + инициализация
struct Object obj1 = {"Поле 1", 2, 3, &PI};   // но так не рекомендуется

struct Object obj2 = {
    .string = "Поле 1",
    .val2 = 2,
    .val3 = 3,
    .pPI = &PI
};


// Передаём структуру в фукнцию
typedef struct User {
    char *login;
    char *password;
    int id;
} User;

int getId(User *user) {
    int id = user->id;
    return id;
}


int main(){

    // Простейший пример
    struct Point {
        int x;
        int y;
    } point;

    point.x = 10, point.y = 20;
    printf("Point coordinates: %d x %d\n", point.x, point.y);

    //Определяем новый тип
    typedef struct newType {
        int x, y;
    } type;

    type var = {5, 10};
    printf("New type var values: %d x %d\n", var.x, var.y);

    // Получаем указатели на структуру, в т.ч. через поле самой структуры
    struct node {
        struct node *next;
    } nod;
    printf("Node: %x x %x\n", &nod, &nod.next);

    // Работаем со структурой в функции
    struct User user = {"Superuser", "Superpassword", 55};
    printf("User id is %d\n", getId(&user));
    return 0;
}