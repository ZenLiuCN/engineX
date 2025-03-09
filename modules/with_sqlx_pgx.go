//go:build ((pgx && sqlx) || all) && !no_pgx

package modules

import _ "github.com/ZenLiuCN/engineX/modules/go/sqlx/pgx"
