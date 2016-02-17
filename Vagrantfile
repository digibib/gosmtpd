# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|

  config.vm.box = "dduportal/boot2docker"
  config.vm.network :forwarded_port, guest: 8080, host: 8080
  config.vm.provision :docker do |d|

    # BUILD APP
    d.build_image "/vagrant",
      args: "-t digibib/build -f /vagrant/Dockerfile.build"

    # COPY COMPILED APP
    d.run "builder",
      image: "digibib/build",
      daemonize: false,
      restart: false,
      args: "--rm -v /vagrant/build:/app",
      cmd: "cp app /app"
  end

  config.vm.provision :docker do |d|
    # BUILD MINIMAL DOCKER IMAGE WITH COMPILED APP
    d.build_image "/vagrant",
      args: "-t digibib/gosmtpd -f /vagrant/Dockerfile"

  end

end
