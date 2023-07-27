package main

import (
	"github.com/anxp/table2cli"
)

func main() {

	tableHeader := []string{
		"TITLE 1",
		"TITLE 2",
		"TITLE 3 a bit longer",
		"TITLE 4",
	}

	tableContent := table2cli.TableContent{
		{
			{
				Data:              "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. ",
				PrintBottomBorder: true,
			},
			{
				Data:              "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				PrintBottomBorder: false,
			},
			{
				Data:              "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. ",
				PrintBottomBorder: true,
			},
			{
				Data:              "Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.",
				PrintBottomBorder: true,
			},
		},
		{
			{
				Data:              "Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.",
				PrintBottomBorder: true,
			},
			{
				Data:              "",
				PrintBottomBorder: true,
			},
			{
				Data:              "Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?",
				PrintBottomBorder: false,
			},
			{
				Data:              "At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga.",
				PrintBottomBorder: true,
			},
		},
		{
			{
				Data:              "Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus.",
				PrintBottomBorder: true,
			},
			{
				Data:              "",
				PrintBottomBorder: true,
			},
			{
				Data:              "",
				PrintBottomBorder: true,
			},
			{
				Data:              "Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae.",
				PrintBottomBorder: true,
			},
		},
	}

	widths := []int{
		10, 20, 20, 20,
	}

	table, _ := table2cli.NewTable(tableHeader, tableContent, widths)

	table.Print()
}
