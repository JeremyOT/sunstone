vagrant_api_version = "2"
instance_count = 3

if ENV['INSTANCES']
  instance_count = ENV['INSTANCES'].to_i
end

Vagrant.configure(vagrant_api_version) do |config|
  (1..instance_count).each do |srv|
    config.vm.define "sunstone#{srv}" do |server|
      server.vm.box = "phusion/ubuntu-14.04-amd64"
      ip = "172.16.19.12#{srv}"
      server.vm.network "private_network", ip: ip

      server.vm.provider :virtualbox do |v, override|
        v.customize ["modifyvm", :id, "--memory", "2048"]
      end

      server.vm.provider :vmware_fusion do |v, override|
        v.vmx['memsize'] = 2048
        v.vmx['displayName'] = "sunstone Test"
        v.vmx['numvcpus'] = "2"
      end

      server.vm.provider :vmware_workstation do |v, override|
        v.vmx['memsize'] = 2048
        v.vmx['displayName'] = "sunstone Test"
        v.vmx['numvcpus'] = "2"
      end

      server.vm.hostname = "sunstone#{srv}"
      server.vm.provision :shell, :inline => <<-SCRIPT
        echo 'DOCKER_OPTS="--icc=true --iptables=true ${DOCKER_OPTS}"\n' >> /etc/default/docker
      SCRIPT

      server.vm.provision :docker do |d|
        d.pull_images 'jeremyot/etcd'
      end
      script = <<-SCRIPT
        apt-get install -y bridge-utils screen
        wget https://storage.googleapis.com/golang/go1.3.3.linux-amd64.tar.gz
        tar -C /usr/local -xzf go1.3.3.linux-amd64.tar.gz
        mkdir -p /var/go/src/github.com/JeremyOT
        ln -s /vagrant /var/go/src/github.com/JeremyOT/sunstone
        echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/bash.bashrc
        echo 'export GOPATH=/var/go' >> /etc/bash.bashrc
        export PATH=$PATH:/usr/local/go/bin
        export GOPATH=/var/go
        go get github.com/JeremyOT/sunstone
        go build -o /home/vagrant/sunstone github.com/JeremyOT/sunstone
        echo "description \\"run sunstone\\"

        start on startup
        stop on runlevel [!2345]

        respawn

        script
        /var/go/bin/sunstone -b docker0 -etcd sunstone 2>&1 >> /var/log/sunstone
        end script
        " > /etc/init/sunstone.conf
        start sunstone
        restart docker
        sleep 1
SCRIPT
      if srv > 1
        script += "docker run -d --net host jeremyot/etcd -peer-addr #{ip}:7001 -addr #{ip}:4001 -peers 172.16.19.121:7001\n"
      else
        script += "docker run -d --net host jeremyot/etcd -peer-addr #{ip}:7001 -addr #{ip}:4001\n"
      end
      script += "sleep 1\n/var/go/bin/sunstone -command -join 127.0.0.1:4001\n"
      server.vm.provision :shell, :inline => script
    end
  end
end
