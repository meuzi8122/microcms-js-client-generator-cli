package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"microcms-js-client-generator-cli/model"

	"github.com/gertd/go-pluralize"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateBaseClientFile(ip string, op string) {
	t := template.Must(template.ParseGlob("./template/base.tmpl"))

	fp, err := os.Create(op + "/base.ts")

	if err != nil {
		fmt.Println(err)
		panic("出力先ディレクトリの取得に失敗しました")
	}

	defer fp.Close()

	if t.Execute(fp, nil) != nil {
		fmt.Println(err)
		panic("ファイルの作成に失敗しました")
	}

}

func GenerateClientFiles(ip string, op string) {

	plu := pluralize.NewClient()

	funcMap := template.FuncMap{
		"convert": func(word string, is_title bool, is_plural bool) string {
			if is_title {
				word = cases.Title(language.Und).String(word)
			}
			if is_plural {
				word = plu.Plural(word)
			}
			return word
		},
	}

	/* 自作関数をテンプレートで使用するためNew()を挟む */
	/* 引数はParseGlob()のファイル名の部分だけ渡す */
	t := template.Must(template.New("client.tmpl").Funcs(funcMap).ParseGlob("./template/client.tmpl"))

	definitions := make([]*model.Definition, 0)

	bytes, err := ioutil.ReadFile(ip)

	if err != nil {
		fmt.Println(err)
		panic("定義ファイルの形式が不正です")
	}

	json.Unmarshal(bytes, &definitions)

	for _, d := range definitions {
		fp, err := os.Create(op + "/" + d.Name + ".ts")

		if err != nil {
			fmt.Println(err)
			panic("出力先ディレクトリの取得に失敗しました")
		}

		defer fp.Close()

		err = t.Execute(fp, map[string]interface{}{
			"name": d.Name,
		})

		if err != nil {
			fmt.Println(err)
			panic("ファイルの作成に失敗しました")
		}
	}

}
