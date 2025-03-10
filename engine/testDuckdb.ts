#! d:/dev/engine -x

import {SQLX} from "go/sqlx";

const pgx = new SQLX("duckdb", "./foo.db")
console.log(pgx.query("select 1 as v "));
