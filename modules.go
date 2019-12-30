package main

import (
	_ "github.com/arj22/registrator/consul"
	_ "github.com/arj22/registrator/consulkv"
	_ "github.com/arj22/registrator/etcd"
	_ "github.com/arj22/registrator/skydns2"
	_ "github.com/arj22/registrator/zookeeper"
)
