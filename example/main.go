package main

import (
	"github.com/anxp/table2cli"
)

func main() {

	tableHeader := []string{
		"Заголовок 1",
		"Заголовок 2",
		"Заголовок 3 более длинный чем обычно",
		"Заголовок 4",
		"Заголовок 5",
		"Заголовок 6",
	}

	tableContent := table2cli.TableContent{
		{
			{
				Data: "Есть много вариантов Lorem Ipsum",
				PrintBottomBorder: true,
			},
			{
				Data: "но большинство из них имеет не всегда приемлемые модификации",
				PrintBottomBorder: false,
			},
			{
				Data: "например, юмористические вставки или слова",
				PrintBottomBorder: true,
			},
			{
				Data: "которые даже отдалённо не напоминают латынь.",
				PrintBottomBorder: true,
			},
			{
				Data: "Если вам нужен Lorem Ipsum для серьёзного проекта,",
				PrintBottomBorder: true,
			},
			{
				Data: "вы наверняка не хотите какой-нибудь шутки, скрытой в середине абзаца.",
				PrintBottomBorder: true,
			},
		},
		{
			{
				Data: "Также все другие известные генераторы Lorem Ipsum используют один и тот же текст,",
				PrintBottomBorder: true,
			},
			{
				Data: "который они просто повторяют,",
				PrintBottomBorder: true,
			},
			{
				Data: "пока не достигнут нужный объём.",
				PrintBottomBorder: false,
			},
			{
				Data: "Это делает предлагаемый здесь генератор единственным настоящим Lorem Ipsum генератором.",
				PrintBottomBorder: true,
			},
			{
				Data: "Он использует словарь из более чем 200 латинских слов,",
				PrintBottomBorder: true,
			},
			{
				Data: "а также набор моделей предложений.",
				PrintBottomBorder: true,
			},
		},
		{
			{
				Data: "В результате сгенерированный Lorem Ipsum выглядит правдоподобно,",
				PrintBottomBorder: true,
			},
			{
				Data: "",
				PrintBottomBorder: true,
			},
			{
				Data: "",
				PrintBottomBorder: true,
			},
			{
				Data: "не имеет повторяющихся абзацей или \"невозможных\" слов.",
				PrintBottomBorder: true,
			},
			{
				Data: "",
				PrintBottomBorder: true,
			},
			{
				Data: "",
				PrintBottomBorder: true,
			},
		},
	}

	widths := []int{
		10, 20, 20, 20, 40, 20,
	}

	table, _ := table2cli.NewTable(tableHeader, tableContent, widths)

	table.Print()
}