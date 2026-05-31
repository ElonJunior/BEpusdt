## MySQL Deprecation Note (Main Branch, Last ~6 Months)

### Conclusion

MySQL is not deprecated because it is unusable. It is being gradually deprecated because the project's main path has shifted to PostgreSQL, while MySQL introduces higher cross-database compatibility and maintenance cost.

### Evidence

1. `docs/docker/docker.md` explicitly states:
   - MySQL support will be gradually deprecated
   - PostgreSQL is recommended for new deployments
   - MySQL remains available for existing deployments

2. Database initialization priority in `app/model/model.go` is:
   - `postgres -> mysql -> sqlite`
   This indicates PostgreSQL is now the preferred path.

3. In `app/model/model.go`, MySQL initialization still carries compatibility switches:
   - `DisableDatetimePrecision`
   - `DontSupportRenameIndex`
   - `DontSupportRenameColumn`
   These are strong signals of extra compatibility burden around MySQL/MariaDB variants.

4. `6a2f3b6` (2026-02-26, `feat: add PostgreSQL database support`) introduces `gorm.io/driver/postgres` and refactors multiple model/query definitions away from MySQL-specific style, including:
   - reducing MySQL-coupled type assumptions
   - cleaning SQL style toward dialect-neutral expressions

5. MySQL operational tuning was still needed in January:
   - `c4644e5`
   - `5ee6895`
   Both add DSN timeout parameters, implying real maintenance overhead in connection/runtime behavior.

6. `2992a1f` (2026-02-27, `perf: remove redundant quotes in SQL queries`) continues the same direction: reduce SQL dialect coupling.

### What "dialect compatibility cost" means

"Dialect" means SQL and behavior differences among database engines (MySQL, PostgreSQL, SQLite), such as:

- data type syntax differences (`tinyint(1)`, `datetime`, boolean mapping)
- identifier quoting differences (backticks vs double quotes)
- DDL capability gaps (rename index/column support varies by engine/version)
- default value, precision, transaction, and lock behavior differences

"Compatibility cost" is the extra engineering effort required to keep one codebase working across those differences:

- writing/maintaining engine-specific branches and flags
- repeatedly adjusting queries/types for portability
- carrying more test matrix and regression risk
- spending extra ops effort on engine-specific tuning (timeouts, connection behavior)

### Practical Interpretation

The project keeps MySQL for backward compatibility, but new architecture and maintenance focus are moving to PostgreSQL. This is a strategic maintenance decision, not an emergency removal.
