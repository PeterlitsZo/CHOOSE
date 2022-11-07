DIR=$(cd $(dirname $0); pwd)

CREATE_COMMAND="migrate create -ext sql -seq -dir $DIR/migration"
OTHER_COMMAND="migrate -database sqlite3://$DIR/database.sqlite -source file://$DIR/migration"

if [[ "$1" == "create" ]]; then
    shift # 移除 create - 因为 CREATE_COMMDN 已经定义子命令了。
    $CREATE_COMMAND "$@"
elif [[ "$1" != "" ]]; then
    $OTHER_COMMAND "$@"
else
    echo "migrate.sh COMMAND [arg...]"
    echo
    echo "migrate.sh 是对 migrate 的简单封装，包括项目相关的命令行配置，有利于保持项目级别的一致。"
    echo "使用 migrate -help 来了解 migrate 本身的帮助。"
    echo "如果不带参数运行本脚本，则显示本文档。"
    echo
    echo "COMMAND:"
    echo "  create          调用 migrate create 命令，但是限定了 create 迁移 SQL 路径和命名规则相关参数。"
    echo "                  建议在后方仅仅添加迁移名，运行之后会在默认的迁移文件夹生成新的空迁移 SQL 文件。"
    echo "  <others>        其他命令（比如 up，down），会根据默认的迁移文件夹来迁移默认的数据库。"
    echo
    echo "默认值:"
    echo "  默认数据库      <migrate.sh 所在文件夹>/database.sqlite3"
    echo "  默认迁移文件夹  <migrate.sh 所在文件夹>/migration/"
fi
