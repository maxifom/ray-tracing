## Трассировка лучей
В папке example_images находятся примеры изображений сцен в формате {название сцены}\_{ширина}\_{высота}\_{количество пикселей для усреднения}.png

Также можно примеры изображений можно посмотреть в [галерее](https://media.maxifom.com/ray-tracing/) 

Сущности трассировщика: 
* Луч
* Интерфейс Hittable - то, во что может попасть лучь (сфера, прямоугольник, треугольник, октаэдр)
* Интерфейс Material - определяет поведение луча при попадании в объект
* Интерфейс Texture - определяет текстуру объекта
* Интерфейс PDF (Функция распределения вероятностей)


### Тестирование скорости
Тестировалось на 4 ядерном Intel Core I7-4770 в 4 потока
##### Cornell Box
Cornell Box 555x555x10: 10.81 секунд

Cornell Box 555x555x25: 22.93 секунд

Cornell Box 555x555x100: 80.37 секунд

Cornell Box 555x555x1000: 768.13 секунд

##### Cornell Box with octahedrons
Cornell Box with octahedrons 600x600x10: 14.54 секунд 

Cornell Box with octahedrons 600x600x25: 30.56 секунд 

Cornell Box with octahedrons 600x600x100: 111.88 секунд

Cornell Box with octahedrons 600x600x1000: 1059.41 секунд

##### Simple Random scene
Simple Random Scene 1920x1080x10: 717.84 секунд 

Simple Random Scene 1920x1080x25: 1299.10 секунд 

Simple Random Scene 1920x1080x100: 5348.86 секунд

##### Simple Scene with octahedrons
Simple Scene (with octahedrons) 1920x1080x10: 1203.50 секунд

Simple Scene (with octahedrons) 1920x1080x25: 3170.42 секунд

Simple Scene (with octahedrons) 1920x1080x100: 10524.04 секунд

##### Simple Scene with spheres
Simple Scene (spheres) 1920x1080x10: 300.26 секунд

Simple Scene (spheres) 1920x1080x25: 699.64 секунд

Simple Scene (spheres) 1920x1080x100: 2705.62 секунд 

Simple Scene (spheres) 1920x1080x100: 22138.72 секунд

#### Вывод по результатам тестирования скорости
На основании результатов 2 последних сцен можно сделать вывод, что алгоритмом трассировки лучей сложнее рисовать фигуры, для которых сложнее считать попал ли луч в фигуру (для октаэдра используется алгоритм Моллера — Трумбора).

Разница в скорости сцены с октаэдрами и с сферами примерно в 4 раза.



##### Использованная литература
[Peter Shirley - Ray Tracing in One Weekend](https://raytracing.github.io/books/RayTracingInOneWeekend.html)

[Peter Shirley - Ray Tracing: The Next Week](https://raytracing.github.io/books/RayTracingTheNextWeek.html)

[Peter Shirley - Ray Tracing: The Rest of Your Life](https://raytracing.github.io/books/RayTracingTheRestOfYourLife.html)
