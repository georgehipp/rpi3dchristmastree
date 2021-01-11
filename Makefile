rpi-ip = 192.168.0.73
rpi-user = pi

build:
	@echo "Compile binary for Raspberry Pi 3 and 4"
	GOARM=7 GOARCH=arm GOOS=linux go build cmd/christmastree.go

push: 
	@echo "Push file to Raspberry Pi"
	scp christmastree $(rpi-user)@$(rpi-ip):/home/pi/

run:
	@echo "Execute on Raspberry Pi"
	ssh -t $(rpi-user)@$(rpi-ip) "./christmastree"

deploy: build push run
	@echo "Build and Deploy"
