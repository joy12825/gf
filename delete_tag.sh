#!/bin/bash

VERSION=$1
if [ ! "$VERSION" ] ;then
    echo "Version cannot empty."
    exit
fi
echo "Git delete tag version $VERSION."

git tag -d contrib/trace/otlphttp/$VERSION
git tag -d contrib/trace/otlpgrpc/$VERSION
git tag -d contrib/trace/jaeger/$VERSION
git tag -d contrib/sdk/httpclient/$VERSION
git tag -d contrib/rpc/grpcx/$VERSION
git tag -d contrib/registry/zookeeper/$VERSION
git tag -d contrib/registry/polaris/$VERSION
git tag -d contrib/registry/nacos/$VERSION
git tag -d contrib/registry/file/$VERSION
git tag -d contrib/registry/etcd/$VERSION
git tag -d contrib/nosql/redis/$VERSION
git tag -d contrib/metric/otelmetric/$VERSION
git tag -d contrib/drivers/sqlitecgo/$VERSION
git tag -d contrib/drivers/sqlite/$VERSION
git tag -d contrib/drivers/pgsql/$VERSION
git tag -d contrib/drivers/oracle/$VERSION
git tag -d contrib/drivers/mysql/$VERSION
git tag -d contrib/drivers/mssql/$VERSION
git tag -d contrib/drivers/dm/$VERSION
git tag -d contrib/drivers/clickhouse/$VERSION
git tag -d contrib/config/polaris/$VERSION
git tag -d contrib/config/nacos/$VERSION
git tag -d contrib/config/kubecm/$VERSION
git tag -d contrib/config/consul/$VERSION
git tag -d contrib/config/apollo/$VERSION
git tag -d cmd/gf/$VERSION
git tag -d cmd/gf/internal/cmd/testdata/build/varmap/$VERSION

git push origin :refs/tags/contrib/trace/otlphttp/$VERSION
git push origin :refs/tags/contrib/trace/otlpgrpc/$VERSION
git push origin :refs/tags/contrib/trace/jaeger/$VERSION
git push origin :refs/tags/contrib/sdk/httpclient/$VERSION
git push origin :refs/tags/contrib/rpc/grpcx/$VERSION
git push origin :refs/tags/contrib/registry/zookeeper/$VERSION
git push origin :refs/tags/contrib/registry/polaris/$VERSION
git push origin :refs/tags/contrib/registry/nacos/$VERSION
git push origin :refs/tags/contrib/registry/file/$VERSION
git push origin :refs/tags/contrib/registry/etcd/$VERSION
git push origin :refs/tags/contrib/nosql/redis/$VERSION
git push origin :refs/tags/contrib/metric/otelmetric/$VERSION
git push origin :refs/tags/contrib/drivers/sqlitecgo/$VERSION
git push origin :refs/tags/contrib/drivers/sqlite/$VERSION
git push origin :refs/tags/contrib/drivers/pgsql/$VERSION
git push origin :refs/tags/contrib/drivers/oracle/$VERSION
git push origin :refs/tags/contrib/drivers/mysql/$VERSION
git push origin :refs/tags/contrib/drivers/mssql/$VERSION
git push origin :refs/tags/contrib/drivers/dm/$VERSION
git push origin :refs/tags/contrib/drivers/clickhouse/$VERSION
git push origin :refs/tags/contrib/config/polaris/$VERSION
git push origin :refs/tags/contrib/config/nacos/$VERSION
git push origin :refs/tags/contrib/config/kubecm/$VERSION
git push origin :refs/tags/contrib/config/consul/$VERSION
git push origin :refs/tags/contrib/config/apollo/$VERSION
git push origin :refs/tags/cmd/gf/$VERSION
git push origin :refs/tags/cmd/gf/internal/cmd/testdata/build/varmap/$VERSION