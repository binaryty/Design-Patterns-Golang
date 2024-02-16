## Фасад (Facade)
структурный шаблон проектирования, позволяющий скрыть сложность системы путём сведения всех возможных внешних вызовов 
к одному объекту, делегирующему их соответствующим объектам системы.

### Проблема
Как обеспечить унифицированный интерфейс с набором разрозненных реализаций или интерфейсов, например, с подсистемой, 
если нежелательно сильное связывание с этой подсистемой или реализация подсистемы может измениться?

### Решение
Определить одну точку взаимодействия с подсистемой — фасадный объект, обеспечивающий общий интерфейс с подсистемой, 
и возложить на него обязанность по взаимодействию с её компонентами. Фасад — это внешний объект, обеспечивающий 
единственную точку входа для служб подсистемы. Реализация других компонентов подсистемы закрыта и не видна внешним компонентам. 
Фасадный объект обеспечивает реализацию GRASP паттерна Устойчивый к изменениям (Protected Variations) с точки зрения защиты 
от изменений в реализации подсистемы.

### Особенности применения
Шаблон применяется для установки некоторого рода политики по отношению к другой группе объектов. 
Если политика должна быть яркой и заметной, следует воспользоваться услугами шаблона Фасад. Если же необходимо обеспечить 
скрытность и аккуратность (прозрачность), более подходящим выбором является шаблон Заместитель (Proxy).