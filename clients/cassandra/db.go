package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	//connect to the cluster
	cluster = gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
