## Паттерны проектирования.

**Шаблон проектирования** (паттерн, от англ. design pattern) — повторяемая архитектурная конструкция в сфере проектирования 
программного обеспечения, предлагающая решение проблемы проектирования в рамках некоторого часто возникающего контекста.

Обычно шаблон не является законченным образцом, который может быть прямо преобразован в код; 
это лишь пример решения задачи, который можно использовать в различных ситуациях. 
Объектно-ориентированные шаблоны показывают отношения и взаимодействия между классами или объектами, без определения того, 
какие конечные классы или объекты приложения будут использоваться

«Низкоуровневые» шаблоны, учитывающие специфику конкретного языка программирования, называются идиомами. 
Это хорошие решения проектирования, характерные для конкретного языка или программной платформы, и потому не универсальные.

На наивысшем уровне существуют архитектурные шаблоны, они охватывают собой архитектуру всей программной системы.

Алгоритмы по своей сути также являются шаблонами, но не проектирования, а вычисления, так как решают вычислительные задачи.

### Плюсы
В сравнении с полностью самостоятельным проектированием, шаблоны обладают рядом преимуществ. 
Основная польза от использования шаблонов состоит в снижении сложности разработки за счёт готовых абстракций для решения 
целого класса проблем. Шаблон даёт решению своё имя, что облегчает коммуникацию между разработчиками, 
позволяя ссылаться на известные шаблоны. Таким образом, за счёт шаблонов производится унификация деталей решений: 
модулей, элементов проекта, — снижается количество ошибок. 
Применение шаблонов концептуально сродни использованию готовых библиотек кода. 
Правильно сформулированный шаблон проектирования позволяет, отыскав удачное решение, пользоваться им снова и снова. 
Набор шаблонов помогает разработчику выбрать возможный, наиболее подходящий вариант проектирования.

### Минусы
Хотя легкое изменение кода под известный шаблон может упростить понимание кода, по мнению Стива Макконнелла, 
с применением шаблонов могут быть связаны две сложности. 

Во-первых, слепое следование некоторому выбранному шаблону может привести к усложнению программы. 
Во-вторых, у разработчика может возникнуть желание попробовать некоторый шаблон в деле без особых оснований.

Многие шаблоны проектирования в объектно-ориентированном проектировании можно рассматривать как идиоматическое 
воспроизведение элементов функциональных языков. 
Питер Норвиг утверждает, что 16 из 23 шаблонов, описанных в книге «Банды четырёх», в динамически-типизируемых языках 
реализуются существенно проще, чем в C++, либо оказываются незаметны. 
Пол Грэхэм считает саму идею шаблонов проектирования — антипаттерном, сигналом о том, что система не обладает 
достаточным уровнем абстракции, и необходима её тщательная переработка. 
Нетрудно видеть, что само определение шаблона как «готового решения, но не прямого обращения к библиотеке» 
по сути означает отказ от повторного использования в пользу дублирования. 
Это, очевидно, может быть неизбежным для сложных систем при использовании языков, не поддерживающих комбинаторы и 
полиморфизм типов, и это в принципе может быть исключено в языках, обладающих свойством гомоиконичности 
(хотя и не обязательно эффективно), так как любой шаблон может быть реализован в виде исполнимого кода.