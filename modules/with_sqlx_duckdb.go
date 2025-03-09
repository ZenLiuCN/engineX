//go:build (all || (duckdb && sqlx)) && !no_duckdb

package modules

import _ "github.com/ZenLiuCN/engineX/modules/go/sqlx/duckdb"
