# MAN Document:
# - 使用 bash migrate.sh 来使用 migration 以将模式迁移到最新模式。
# - 如果带有参数，那么退化到使用指定了参数的 migrate 命令（见 back/README.md）

DIR=$(cd $(dirname $0); pwd)

COMMAND="migrate --database sqlite3://$DIR/database.sqlite --source file://$DIR/migration"

if [[ "$1" != "" ]]; then
    $COMMAND "$@"
else
    $COMMAND up
fi
