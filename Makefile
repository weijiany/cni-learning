build:
	env GOOS=linux GOARCH=arm go build -o my-cni

test:
	CNI_COMMAND=ADD CNI_CONTAINERID=lab-ns CNI_NETNS=/var/run/netns/lab-ns CNI_IFNAME=eth0 CNI_PATH=/app go run main.go < 100-my-cni.conf

fetch-keys:
	etcdctl get --keys-only --prefix=true "/my-cni/ipam/"
