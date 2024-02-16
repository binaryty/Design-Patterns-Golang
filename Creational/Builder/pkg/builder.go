package pkg

// Builder содержит методы для конструирования первого символа идентификатора (BuildStart),
// первого символа каждого последующего слова (BuildFirstChar), последующих символов слов (BuildNextChar)
// и символов-разделителей между словами (BuildDelim). Методы BuildStart, BuildFirstChar и BuildNextChar
// имеют символьный параметр, используемый при конструировании (возможно, после изменения его регистра);
// метод BuildDelim не имеет параметров. Все эти методы не возвращают значений;
// в классе Builder они не выполняют никаких действий. Кроме того,
// класс Builder содержит абстрактный метод GetResult без параметров, возвращающий строковое значение.
type Builder interface {
	BuildStart(char rune)
	BuildFirstChar(char rune)
	BuildNextChar(char rune)
	BuildDelim()
	GetResult() string
}
