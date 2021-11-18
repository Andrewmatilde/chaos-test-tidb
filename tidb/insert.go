package tidb

import (
	"bytes"
	"fmt"
	"k8s.io/apimachinery/pkg/util/rand"
	"log"
	"strconv"
	"sync"
	"text/template"
)

type InsertMap struct {
	Tables map[int]Tables
}

type Tables struct {
	Keys map[int]bool
}

func GetInsertMap(tableCount int, KeyCount int) InsertMap {
	var im InsertMap
	im.Tables = make(map[int]Tables)
	for i := 0; i < tableCount; i++ {
		var t Tables
		t.Keys = make(map[int]bool)
		for j := 0; j < KeyCount; j++ {
			key := rand.Int()
			t.Keys[key] = false
		}
		im.Tables[i] = t
	}
	return im
}

func InsertCase(client Client) error {
	im := GetInsertMap(64, 1<<20)
	valueLength := 1 << 5
	var wg sync.WaitGroup
	for tableId, table := range im.Tables {
		wg.Add(1)
		tableId := tableId
		table := table
		go func() {
			var tableSchema TableSchema
			tableSchema.TableName = "t" + strconv.Itoa(tableId)
			for i := 0; i < valueLength; i++ {
				tableSchema.KeyNames = append(tableSchema.KeyNames, "c"+strconv.Itoa(i))
			}
			createTable := template.New("createTable")
			createTable = template.Must(createTable.Parse(CreateTemp))
			var createSql bytes.Buffer
			err := createTable.Execute(&createSql, tableSchema)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("Create Table :", tableSchema.TableName)
			_, err = client.db.Exec(createSql.String())
			if err != nil {
				log.Println(err)
				return
			}
			defer func() {
				dropTable := template.New("dropTable")
				dropTable = template.Must(dropTable.Parse(DropTemp))
				var dropSql bytes.Buffer
				err := dropTable.Execute(&dropSql, tableSchema)
				if err != nil {
					log.Println(err)
					wg.Done()
				}
				fmt.Println("Drop Table :", tableSchema.TableName)
				_, err = client.db.Exec(dropSql.String())
				if err != nil {
					log.Println(err)
					wg.Done()
				}
				wg.Done()
			}()

			insertTable := template.New("insertTable")
			insertTable = template.Must(insertTable.Parse(InsertTemp))

			for key, _ := range table.Keys {
				var insertSql bytes.Buffer
				tableSchema.Id = strconv.Itoa(key)
				tableSchema.Values = make([]string, valueLength)
				for i := 0; i < valueLength; i++ {
					tableSchema.Values[i] = strconv.Itoa(rand.Int())
				}
				err := insertTable.Execute(&insertSql, tableSchema)
				if err != nil {
					log.Println(err)
				}
				_, err = client.db.Exec(insertSql.String())
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}
	wg.Wait()

	return nil
}

type TableSchema struct {
	TableName string
	KeyNames  []string
	Id        string
	Values    []string
}

const CreateTemp = `
CREATE TABLE {{.TableName}} ({{range .KeyNames}}
    {{.}} bigint, 
{{- end }}
id bigint not null primary key
);
`

const DropTemp = `DROP TABLE {{.TableName}};`

const InsertTemp = `
INSERT INTO {{.TableName}} ({{range .KeyNames}}{{.}},{{- end }}id)
VALUES({{range .Values}}{{.}},{{- end }}{{.Id}});
`
