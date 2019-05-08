# table2tsv

## これは何

標準入力に入ったhtml の`<table>`〜`</table>`の各要素をtsvで標準出力に吐く

## 使用例

```bash
$ go build -o table2tsv table2tsv.go
$ curl --silent -X GET "http://www.htmq.com/html/table.shtml" | ./table2tsv
```

