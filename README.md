# extened with [engine](https://github.com/ZenLiuCN/engine)

## example
### install
```shell
CGO_LDFLAGS="-L$(pwd)/.libs/win/" \
go install -tags=duckdb_use_lib,all github.com/ZenLiuCN/engineX/engine@latest
```
### init an project
```shell
cd example
engine -dp
```
### write script
helper.ts
```ts
// @ts-ignore
import {SQLX} from "go/sqlx"


export class Duck {
    db: SQLX

    constructor(relativeExtensions: boolean = true) {
        this.db = new SQLX("duckdb", "?access_mode=READ_WRITE&allow_unsigned_extensions=true")
        if (relativeExtensions)
            this.db.exec("SET extension_directory= './extensions'")
    }


    use(fn: (d: Duck) => void) {
        fn(this)
        return this
    }

    apply<T>(fn: (d: Duck) => T): T {
        return fn(this)
    }
    read(fn: (array: Array<any>) =>void, query: string, args?: Record<string, any>){
        fn(this.query(query,args))
        return this
    }
    modify(fn: (array: { lastInsertedId: number; rowsAffected: number }) =>void, query: string, args?: Record<string, any>){
        fn(this.exec(query,args))
        return this
    }
    get<T>(fn: (array: Array<any>) => T, query: string, args?: Record<string, any>) {
        return fn(this.query(query, args))
    }

    mut<T>(fn: (result: {
        lastInsertedId: number
        rowsAffected: number
    }) => T, query: string, args?: Record<string, any>) {
        return fn(this.exec(query, args))
    }

    query(query: string, args?: Record<string, any>) {
        return this.db.query(query, args)
    }

    exec(query: string, args?: Record<string, any>): {
        lastInsertedId: number
        rowsAffected: number
    } {
        return this.db.exec(query, args)
    }

    close() {
        this.db?.close()
        this.db = null
    }

    install(ext: string) {
        this.db.exec(`INSTALL ${ext}`)
        return this
    }

    xPostgres() {
        this.install("postgres")
        return this
    }

    xMySQL() {
        return this.install("mysql")

    }

    xHttp() {
        return this.install("http")
    }

    xExcel() {
        return this.install("excel")

    }

    xFts() {
        return this.install("fts")

    }

    xArrow() {
        return this.install("arrow")

    }

    xInet() {
        return this.install("inet")

    }

    xVss() {
        return this.install("vss")

    }

    extensions(): {
        aliases: string[],
        description: string,
        extension_name: string,
        extension_version: string,
        install_mode?: string,
        install_path: string,
        installed: boolean,
        installed_from: string,
        loaded: boolean
    }[] {
        return this.query("FROM duckdb_extensions()")
    }
}
```
testing.ts
```ts
#! engine

import {Duck} from "./helper";
const db=new Duck()
const {log,table}=console
    db
        .read(log,"SELECT current_setting('access_mode')")
        .use(d=>table(d.extensions()))

/*log(db.exec("CREATE TABLE users(name VARCHAR, age INTEGER, height FLOAT, awesome BOOLEAN, bday DATE)"))
log(db.exec("INSERT INTO users VALUES('marc', 99, 1.91, true, '1970-01-01')"))
log(db.exec("INSERT INTO users VALUES('macgyver', 70, 1.85, true, '1951-01-23')"))
log(db.query('	SELECT name, age, height, awesome, bday	FROM users	WHERE (name = :name1 OR name = :name2) AND age > :age AND awesome = :awesome',	{
    name1:  "macgyver", name2:"marc", age:30, awesome:true,
}))
log(db.query("SELECT ('{\"duck\": 42}'::JSON->'$.duck')::INTEGER Val"))
log(db.query("SELECT excel_text(1234567.897, 'h:mm AM/PM') AS timestamp"))

log(db.exec("CREATE TABLE documents (document_identifier VARCHAR, text_content VARCHAR, author VARCHAR, doc_version INTEGER)"))
log(db.exec("INSERT INTO documents VALUES ('doc1', 'The mallard is a dabbling duck that breeds throughout the temperate.', 'Hannes Mühleisen', 3), ('doc2', 'The cat is a domestic species of small carnivorous mammal.', 'Laurens Kuiper', 2);"))
log(db.exec("PRAGMA create_fts_index('documents', 'document_identifier', 'text_content', 'author')"))
log(db.query("SELECT document_identifier, text_content, score FROM ( SELECT *, fts_main_documents.match_bm25(  document_identifier, 'Muhleisen', fields := 'author') AS score FROM documents) sq WHERE score IS NOT NULL AND doc_version > 2 ORDER BY score DESC"))
log(db.query("SELECT document_identifier, text_content, score FROM ( SELECT *, fts_main_documents.match_bm25( document_identifier,'small cats') AS score  FROM documents ) sq WHERE score IS NOT NULL ORDER BY score DESC"))
table(db.query("FROM duckdb_extensions()"))
log(db.query("SELECT '127.0.0.1'::INET AS addr"))*/
db.close()
```
### execute
```shell
engine testing.ts
```
