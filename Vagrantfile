Vagrant.configure("2") do |config|
  config.vm.define "master" do |m|
    m.vm.box = "ubuntu/focal64"
    m.vm.provider "virtualbox" do |v|
      v.memory = 2048
      v.cpus = 2
    end

    m.vm.network :private_network, ip: "192.168.56.10"
    m.vm.synced_folder ".", "/app"
  end
end
