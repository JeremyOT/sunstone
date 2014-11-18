Sunstone
========
Named after the Viking navigational tool ([Sunstone][]), Sunstone will help your packets find their way between your [Docker](http://docker.com) containers.

[Sunstone]: http://en.wikipedia.org/wiki/Sunstone_(medieval)

How it works
------------
Sunstone uses GRE ([Generic Routing Encapsulation][]) to create tunnels between your Docker hosts. Sunstone runs on each host and tracks the status of all other hosts in the cluster. The `docker0` bridge on each host is assigned a unique subnet. Tunnels and routing rules linking each bridge to the rest of the cluster are created and destroyed as nodes join and leave the cluster. This allows any container or host to communicate with any other container in the cluster without having to expose ports publicly on the host.

Example: in a cluster with web services, a database, a message queue and workers, all containers may communicate with each other and only the web service would require publicly exposed ports. E.g. `docker run -p 443:443 webservice`.

[Generic Routing Encapsulation]: http://en.wikipedia.org/wiki/Generic_Routing_Encapsulation

Usage
-----
Currently, Sunstone requires that every host in a cluster be on the same /24 subnet. There are two ways to run a Sunstone cluster, either internally managed or externally, with [Etcd][]. In either case, it is best to start Sunstone before Docker and either join a cluster at startup or at some point later.

First, start sunstone and create (or replace) a `docker0` bridge with a unique /24 subnet based on the host's `eth0` IPv4. In the Etcd case, configure Sunstone to read cluster state from `<etcd_host>/v2/keys/sunstone` once connected to a cluster.

```
#Internal Cluster:
sunstone -b docker0

#Etcd Cluster:
sunstone -b docker0 -etcd sunstone
```

Now join a cluster. If using Etcd, this is mandatory. For Sunstone managed clusters, skip this for the first node.

```
#Internal Cluster:
sunstone -command -join <comma delimited list of seed addresses>

#Etcd Cluster:
sunstone -command -join <comma delimited list of Etcd addresses>
```

Congratulations, you've built a cluster. Wait a moment and run `ip address` to see your newly created GRE tunnels. As you add and remove nodes you can watch as tunnels between each host are created and destroyed. Now, start some containers and you'll be able to make connections across hosts.

[Etcd]: https://github.com/coreos/etcd

Todo
----
- Where are the tests?
- Allow clusters to span larger subnets and WAN.
- Add support for alternative tunnelling protocols.
