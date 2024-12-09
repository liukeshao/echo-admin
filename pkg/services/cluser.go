package services

import (
	"context"
	"github.com/liukeshao/echo-admin/ent"
	"github.com/liukeshao/echo-admin/ent/cluster"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ClusterService struct {
	orm *ent.Client
}

func NewClusterService(orm *ent.Client) *ClusterService {
	return &ClusterService{orm: orm}
}

func (s *ClusterService) getK8sClient(ctx context.Context, name string) (*kubernetes.Clientset, error) {
	c, err := s.orm.Cluster.Query().
		Where(cluster.Name(name)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Config))
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}
