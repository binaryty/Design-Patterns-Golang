## Singleton

**Одиночка** (англ. Singleton) — порождающий шаблон проектирования, гарантирующий, что в однопоточном приложении будет единственный экземпляр некоторого класса, и предоставляющий глобальную точку доступа к этому экземпляру.

Проблемы которые решает паттерн - нарушая принцип единственной ответственности, может гарантировать наличие единственного экземпляра обекта, 
а также предоставляет глобальную точку доступа к одному экземпляру оъекта 

Преимущества:
- гарантирование единсвтенного экземаляра объекта
- глобальная точка доступа к одному обхекту
- реализует отложенную инициаллизацию одного объекта

Недостатки:
- нарушение принципа единственной ответственности
- маскирует плохой дизайн проекта
- проблемы многопоточности
- требует постоянного создания MOC-объектов при тестировании