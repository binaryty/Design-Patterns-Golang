### Задача:
Реализовать иерархию классов-животных с абстрактным предком **Animal**, 
содержащим метод **GetInfo**, который возвращает строковое значение, и шестью конкретными потомками: 
_Lion_, _Tiger_, _Leopard_ (кошачьи, cats), 
_Gorilla_, _Orangutan_, _Chimpanzee_ (человекообразные обезьяны, apes). 

Каждый конкретный класс содержит строковое поле name (имя животного), которое определяется в его конструкторе с помощью 
одноименного параметра. 
Метод **GetInfo** возвращает имя класса и имя животного, разделенные пробелом, например, «Lion Tom».

Реализовать иерархию классов-создателей с абстрактным предком **AnimalCreator** и конкретными потомками **CatCreator** и **ApeCreator**. 
Фабричный метод **CreateAnimal**(_ind_, _name_) этих классов принимает параметр целого типа _ind_ и строковый параметр name 
и возвращает объект типа **Animal**. Для конкретных классов **CatCreator** и **ApeCreator** параметр _ind_ метода **CreateAnimal** 
определяет тип создаваемого животного по его индексу (0, 1 или 2) в группе кошачьих или обезьян, а параметр _name_ — имя животного.

В классе **AnimalCreator** также определить метод **GetZoo** с двумя параметрами-массивами _idxes_ и _names_ одинакового размера; 
массив _idxes_ содержит целые числа, массив _names_ — строки (предполагается, что элементы массива _idxes_ всегда имеют значения 
в диапазоне от 0 до 2). Метод **GetZoo** возвращает массив объектов **Animal** того же размера, что и массивы _idxes_ и _names_; 
каждый элемент полученного массива определяется с помощью фабричного метода с параметрами, равными значениям 
соответствующих элементов массивов _idxes_ и _names_.